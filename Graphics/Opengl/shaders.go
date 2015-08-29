package Opengl

const (
	vertexShaderSource = `
#version 330 core

// Input vertex data, different for all executions of this shader.
in vec3 vertexPosition_modelspace;

in vec4 vertexColor;

in vec2 vertexUV;


in vec3 translation;
in vec4 rotation;
in vec3 scale;

// Output data ; will be interpolated for each fragment.
out vec2 UV;

// Output data ; will be interpolated for each fragment.
out vec4 fragmentColor;

// Values that stay constant for the whole mesh.
uniform mat4 MVP;


mat4 rotationMatrix(vec3 axis, float angle)
{
    axis = normalize(axis);
    float s = sin(angle);
    float c = cos(angle);
    float oc = 1.0 - c;
    
    return mat4(oc * axis.x * axis.x + c,           oc * axis.x * axis.y - axis.z * s,  oc * axis.z * axis.x + axis.y * s,  0.0,
                oc * axis.x * axis.y + axis.z * s,  oc * axis.y * axis.y + c,           oc * axis.y * axis.z - axis.x * s,  0.0,
                oc * axis.z * axis.x - axis.y * s,  oc * axis.y * axis.z + axis.x * s,  oc * axis.z * axis.z + c,           0.0,
                0.0,                                0.0,                                0.0,                                1.0);
}

mat4 translationMatrix(vec3 pos)
{
    // create identity
    mat4 translation = mat4(1.0);
    translation[3] = vec4(pos,1.0);
    return translation;
}

mat4 scaleMatrix(vec3 s)
{
    // create identity
    mat4 scale = mat4(1.0);
    scale[0][0] = s[0];
    scale[1][1] = s[1];
    scale[2][2] = s[2];
    return scale;
}


void main(){
    

    mat4 Model = mat4(1.0);

    mat4 t = translationMatrix(translation);
    mat4 s = scaleMatrix(scale);
    mat4 r = rotationMatrix(vec3(rotation), rotation[3]);

    Model *= t * r * s;


    gl_Position = MVP * Model * vec4 (vertexPosition_modelspace,1); //* MVP;

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
