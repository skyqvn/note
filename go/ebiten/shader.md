# Shader

## Example

### 3D光影算法

```go
//go:build ignore

// Specify the 'pixel' mode.
//kage:unit pixels

package main

//传入的参数（统一变量）
// Uniform variables.
var Time float
var Cursor vec2

// Fragment is the entry point of the fragment shader.
// Fragment returns the color value for the current position.
func Fragment(dstPos vec4, srcPos vec2, color vec4) vec4 {
	// You can define variables with a short variable declaration like Go.
	//目标点相对于目标原点的位置
	pos := dstPos.xy - imageDstOrigin()
	//灯光在鼠标xy,50位置
	lightpos := vec3(Cursor, 50)
	//计算灯光与目标点的相对三维位置（向量方向）
	lightdir := normalize(lightpos - vec3(pos, 0))
	//获取法线图（Src1）上目标点的法线方向（向量方向）。法线图（Src1）上每一点对应着二维图像图（Src1）上点的法线方向。
	normal := normalize(imageSrc1At(srcPos) - 0.5)
	//环境光，25%的权。
	const ambient = 0.25
	//照明光强度。
	//dot(normal.xyz, lightdir)计算法线方向与光照方向的点积，点积越大，说明夹角越小，则物体法线离光照的轴线越近，光照越强。乘0.75表示占75%的权。
	//当dot(normal.xyz, lightdir)小于0时，法线与光线夹角大于90°，即面背对光源，所以不发光。
	diffuse := 0.75 * max(0.0, dot(normal.xyz, lightdir))

	// You can treat multiple source images by
	// imageSrc[N]At or imageSrc[N]UnsafeAt.
	//原图颜色乘光照强度（照明光+环境光）
	return imageSrc0At(srcPos) * (ambient + diffuse)
}

```

### 渐消失

```go
//go:build ignore

//kage:unit pixels

package main

var Time float
var Cursor vec2

func Fragment(dstPos vec4, srcPos vec2, color vec4) vec4 {
	// Triangle wave to go 0-->1-->0...
	//三角波0->1->0...
	//Time除数越大就越慢
	limit := abs(2*fract(Time/3) - 1)
	//用红色通道判断
	level := imageSrc3UnsafeAt(srcPos).x

	// Add a white border
	//添加白色边缘
	if limit-0.1 < level && level < limit {
		//原图片点透明通道值
		alpha := imageSrc0UnsafeAt(srcPos).w
		//相当于vec4(alpha,alpha,alpha,alpha)
		return vec4(alpha)
	}

	//是否有颜色
	//step(limit, level)如果limit大于level就返回0，否则为1
	return step(limit, level) * imageSrc0UnsafeAt(srcPos)
}

```

