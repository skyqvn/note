# Mutool

> See [MuPDF on the command line - MuPDF 1.24.0 documentation](https://mupdf.readthedocs.io/en/latest/mupdf-command-line.html)

### 转图片
```bash
//以600分辨率将pdf转为图片
//See https://mupdf.readthedocs.io/en/latest/mutool-draw.html
mutool draw -o output-%d.png -r 600 input.pdf

```



### 提取资源

```bash
mutool extract -N xxx.pdf
```



### 合并/提取页面

```bash
mutool merge -o output.pdf input1.pdf 1,2,5-N input2.pdf
```

