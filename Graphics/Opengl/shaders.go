package Opengl

const (
	NO_TEXTURE float32 = 0.1
	TEXTURED   float32 = 1.1
)

const (
	vertexShaderSource = `
#version 300 es

// Input vertex data, different for all executions of this shader.
in vec3 vertexPosition_modelspace;

in vec4 vertexColor;

in vec2 vertexUV;

in float mode;
out float modeOut;

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

    // UV of the vertex. No special space for this one.
    UV = vertexUV;

}
` + "\x00"
	fragmentShaderSource = `
#version 300 es


// Interpolated values from the vertex shaders
in mediump vec4 fragmentColor;

// Display uv?
in mediump float modeOut;

// Interpolated values from the vertex shaders
in mediump vec2 UV;

// Ouput data
out mediump vec4 color;

// Values that stay constant for the whole mesh.
uniform sampler2D myTextureSampler;

void main()
{

    // // Output color = red
    // color = vec4(fragmentColor,1.0);

    // Output color = color of the texture at the specified UV
    //color = texture2D( myTextureSampler, UV ).rgba;

    

    // Do we display textures or not?
    if (int(modeOut) == 1) {
        mediump vec4 tex = texture2D( myTextureSampler, UV );
    
        tex.a *= fragmentColor[3];
        // color = tex.rgba;
        color =  tex + vec4(fragmentColor[0], fragmentColor[1], fragmentColor[2], 1)*tex.a;
    } else {

        color =  vec4(fragmentColor[0], fragmentColor[1], fragmentColor[2], 1);
     }
}
` + "\x00"
)
