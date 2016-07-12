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
out vec3 fragPos;

// model view normal
in vec3 mNorm;

// world normal
in vec3 wNorm;
out vec3 wNormFrag;

in vec4 diffuse;

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
out vec4 diffuseFragment;

// Values that stay constant for the whole mesh.
uniform mat4 MVP;




void main(){



    gl_Position = MVP * vec4 (vertexPosition_modelspace,1); //* MVP;

    // The color of each vertex will be interpolated
    // to produce the color of each fragment
    diffuseFragment = diffuse;

    // display uv textures?
    modeOut = mode;

    // which sampler to use
    samplerIdxOut = samplerIdx;

		// pass the pos to fragment shader
		fragPos = vertexPosition_modelspace;

		// pass the world normal to fragment shader
		wNormFrag = wNorm;

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


struct Light {
   vec3 position;
   vec3 intensities; //a.k.a the color of the light
   float attenuation;
   float ambientCoefficient;
} ;

uniform Light light;
uniform vec3 cameraPosition;

// Interpolated values from the vertex shaders
in mediump vec4 diffuseFragment;

// fragment position passed vertex position.. transformed already
in mediump vec3 fragPos;

// fragment world normal
in mediump vec3 wNormFrag;

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

vec3 linearLight(in vec4 surfaceColor) {
	vec3 normal = wNormFrag;
	vec3 surfacePos = fragPos;


	vec3 surfaceToLight = normalize(light.position - surfacePos);
	vec3 surfaceToCamera = normalize(cameraPosition - surfacePos);

	//ambient
	vec3 ambient = .2 * surfaceColor.rgb * light.intensities;

	//diffuse
	float diffuseCoefficient = max(0.0, dot(normal, surfaceToLight));
	vec3 diffuse = diffuseCoefficient * surfaceColor.rgb * light.intensities;

	//specular
	float specularCoefficient = 0.0;
	if(diffuseCoefficient > 0.0)
			specularCoefficient = pow(max(0.0, dot(surfaceToCamera, reflect(-surfaceToLight, normal))), .5);
	vec3 specular = specularCoefficient * vec3(.2,1,.2) * light.intensities;

	//attenuation
  float distanceToLight = length(light.position - surfacePos);
  float attenuation = 1.0 / (1.0 + light.attenuation * pow(distanceToLight, 2));

	return ambient + attenuation*(diffuse + specular);
}

void main()
{
    // // Output color = red
    // color = vec4(fragmentColor,1.0);

    // Output color = color of the texture at the specified UV
    //color = texture2D( myTextureSampler, UV ).rgba;

    mediump vec4 tex ;

   ` + samplerIdxStr + `

    // Do we display textures or not?
    if (int(modeOut) == 0) {
			tex = vec4(0.0);
			tex.a = 1.0;
		}
    // mediump vec4 tex = texture2D( myTextureSampler[0], UV );

    tex.a *= diffuseFragment[3];
    // color = tex.rgba;
		vec4 surfaceColor = vec4(diffuseFragment[0], diffuseFragment[1], diffuseFragment[2], 1);
    color =  tex + surfaceColor*tex.a;

		vec3 gamma = vec3(1.0/2.2);
		// color = vec4(pow(linearLight(tex),gamma),tex.a);
}


` + "\x00"
}

var (
	OGL_VERSION string = "130"
)
