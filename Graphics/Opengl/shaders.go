package Opengl

const (
	vertexShaderSource = `
#version 330 core

// Input vertex data, different for all executions of this shader.
in vec3 vertexPosition_modelspace;

in vec4 vertexColor;

in vec2 vertexUV;

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

    // UV of the vertex. No special space for this one.
    UV = vertexUV;

}
` + "\x00"
	fragmentShaderSource = `
#version 330 core

// Interpolated values from the vertex shaders
in vec4 fragmentColor;

// Interpolated values from the vertex shaders
in vec2 UV;

// Ouput data
out vec4 color;

// Values that stay constant for the whole mesh.
uniform sampler2D myTextureSampler;

void main()
{

    // // Output color = red 
    // color = vec4(fragmentColor,1.0);

    // Output color = color of the texture at the specified UV
    //color = texture2D( myTextureSampler, UV ).rgba;
    vec4 tex = texture2D( myTextureSampler, UV );
    tex.a *= fragmentColor[3];
    // color = tex.rgba;
    color =  tex + vec4(fragmentColor[0], fragmentColor[1], fragmentColor[2], 1)*tex.a;


}
` + "\x00"
)
