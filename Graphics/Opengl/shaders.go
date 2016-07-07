package Opengl

import "strconv"

const (
	NO_TEXTURE float32 = 0.1
	TEXTURED   float32 = 1.1
)

func samplerIdx(idx int) string {
	strIdx := strconv.Itoa(idx)
	return `    
    if (int(samplerIdxOut) == ` + strIdx + `) {
        tex = texture2D( myTextureSampler[` + strIdx + `], UV );
    }`
}

func VertexShader() string {
	return `

#version ` + OGL_VERSION + ` 

// Input vertex data, different for all executions of this shader.
in vec3 vertexPosition_modelspace;

in vec4 vertexColor;

in vec2 vertexUV;

// mode set
in float mode;
out float modeOut;

//sampleridx set
in float samplerIdx;
out float samplerIdxOut;

// Output data ; will be interpolated for each fragment.
out vec2 UV;

// Output data ; will be interpolated for each fragment.
out vec4 fragmentColor;

// Values that stay constant for the whole mesh.
uniform mat4 MVP;

void main(){



    gl_Position = MVP * vec4 (vertexPosition_modelspace,1); //* MVP;

    // The color of each vertex will be interpolated
    // to produce the color of each fragment
    fragmentColor = vertexColor;

    // display uv textures?
    modeOut = mode;
    
    // which sampler to use
    samplerIdxOut = samplerIdx;

    // UV of the vertex. No special space for this one.
    UV = vertexUV;

}
` + "\x00"

}

func FragmentShader() string {
	strTexUnits := strconv.Itoa(texUnits)

	samplerIdxStr := ""

	// texNum := int(Probe().MaxTextureImageUnits)

	for i := 0; i < texUnits; i++ {
		samplerIdxStr += samplerIdx(i)
	}

	return `
#version ` + OGL_VERSION + ` 


// Interpolated values from the vertex shaders
in mediump vec4 fragmentColor;

// Display uv?
in mediump float modeOut;

// which sampler to use
in mediump float samplerIdxOut;

// Interpolated values from the vertex shaders
in mediump vec2 UV;

// Ouput data
out mediump vec4 color;

// Values that stay constant for the whole mesh.
uniform sampler2D myTextureSampler[` + strTexUnits + `];

void main()
{
    // // Output color = red
    // color = vec4(fragmentColor,1.0);

    // Output color = color of the texture at the specified UV
    //color = texture2D( myTextureSampler, UV ).rgba;
    
    mediump vec4 tex ;
    
   ` + samplerIdxStr + `

    // Do we display textures or not?
    if (int(modeOut) == 1) {
        // mediump vec4 tex = texture2D( myTextureSampler[0], UV );
    
        tex.a *= fragmentColor[3];
        // color = tex.rgba;
        color =  tex + vec4(fragmentColor[0], fragmentColor[1], fragmentColor[2], 1)*tex.a;
    } else {

        color =  vec4(fragmentColor[0], fragmentColor[1], fragmentColor[2], 1);
     }
}
` + "\x00"
}

var (
	OGL_VERSION string = "130"
)
