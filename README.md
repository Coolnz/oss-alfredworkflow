
# oss-alfredworkflow

统一用一个图床上传工具，是最优选择，因为不同软件，上传图片到阿里云OSS时，设置的目录结构是不同的（比如有的是按照时间区分文件夹，有的按照分类名称区分文件夹），并且很多软件的图床工具并不好用。用一个全局的图床工具，就能避免这些问题，维护成本也很低，也避免了在不同软件重复配置阿里云OSS，以及上述所说目录结构不同的问题。


## Installation

在 [Releases](https://github.com/Coo6ee/oss-alfredworkflow/releases) 页面下载即可

## Usage/Examples

~~复制图片后使用自定义的`oss-alfredowrkflow关键字`，即可上传图片到OSS~~

使用Universal Actions上传文件


## Features

- 加个"应用"关键字，用来分类，以区分不同应用
- 自定义文件名
- 自定义后缀
- ~~可配置默认返回格式~~


## Environment Variables

OSS相关配置

- Bucket
- Access Key
- Secret Key
- Endpoint
- Custom Domain
- Fileame
- Suffix

### Filename

- `${MDFile}`: Current markdown file name, such as "example"
- `${Suffix}`: Current image file suffix, such as "png" ${yyyyMMdd}: Current date, such as "20200401"
- `${Timestamp}`: Current time stamp, such as "1585819668627"
- `${UUID}`: Random 32-bit string, such as "67b52ab3e50643e08b8cb980c2ecdaed"

### Suffix

比如`?x-oss-process= image/auto-orient, 1/interlace, 1/quality, q_50/format jpg`，更多后缀查看 [图片处理操作方式](https://help.aliyun.com/document_detail/44686.html)



