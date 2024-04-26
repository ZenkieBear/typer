[![Go Reference](https://img.shields.io/badge/Reference-white?logo=go&labelColor=white&color=91DFFB)](https://pkg.go.dev/github.com/ZenkieBear/typer)

<br>
<br>

<p align='center'>
  <img src="doc/typer-logo.png" width='120' height='120' style='border-radius: 15px'>
</p>

# Typer
Typer 是一款打字机输出模拟器，让你的 log 输出更加生动！

![Shown](doc/show.gif)


# 立即体验！
## 先决条件
在您的个人电脑上已经安装 [Go](https://go.dev/doc/install) 和 [Git](https://git-scm.com/downloads)。

## 将 Typer 仓库克隆到您本机
```shell
git clone https://github.com/ZenkieBear/typer
```

## 运行 `main.go`
```shell
go run main.go
```



# 在您的项目中使用 Typer
## 安装 Typer
```shell
go get github.com/ZenkieBear/typer
```

## 将 Typer 导入您的 go 文件
```go
import "github.com/ZenkieBear/typer/typer"
```

## 使用 Typer
```go
func some() {
	typer.Print("Hey Judy\n")
	typer.Println("Don't be afraid")
}
```


# 配置
Typer 支持自定义配置。

你可以创建一个 `typer.Typer`, 并根据您的喜好来定义属性。

```go
func demo() {
  // ...
	myTyper := typer.Typer{
		Base:       300,
		FloatRange: 50,
		Printer:    fmt.Print,
	}

	if err := myTyper.Print("你好啊!"); err != nil {
		fmt.Println(err.Error())
	}
  // ...
}
```


# 许可证
Typer 使用 [MIT](LICENSE) 开源证书授权

