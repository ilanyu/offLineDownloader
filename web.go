package main

import (
	"github.com/valyala/fasthttp"
	"log"
	"os"
	"strings"
	"bytes"
	"io/ioutil"
	"strconv"
	"time"
	"fmt"
	"os/exec"
	"runtime"
	"flag"
)

//从url中获取文件名
func getFileName(url string) string {
	fileNames := strings.Split(url, "/")
	fileName := fileNames[len(fileNames) - 1]
	if fileName == "" {
		fileName = "noname"
	}
	fileName = strings.Replace(fileName, "?", "_", -1)
	fileName = strings.Replace(fileName, "#", "_", -1)
	fileName = strings.Replace(fileName, "&", "_", -1)
	return fileName
}

//使用fasthttp下载文件
func download(url string) bool {
	fileName := getFileName(url)
	out, _ := os.Create("./download/" + fileName)
	defer out.Close()
	statusCode, body, err := fasthttp.Get(nil, url)
	if statusCode == 200 && err == nil {
		out.Write(body)
	}
	return true
}

//调用wget下载文件
func downloadByWget(url string) bool {
	fileName := getFileName(url)
	cmd := exec.Command("wget", "-c", "-O", "./download/" + fileName, url)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
	return true
}

func main() {
	host := flag.String("host", "0.0.0.0", "host")
	port := flag.String("port", "8081", "port")
	flag.Parse()
	m := func(ctx *fasthttp.RequestCtx) {
		switch {
		case string(ctx.Path()) == "/":
			ctx.SetContentType("text/html; charset=utf-8")
			ctx.WriteString(`<!doctype html>
<html class="no-js">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="description" content="">
    <meta name="keywords" content="">
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no">
    <title>离线下载</title>
    <meta name="renderer" content="webkit">
    <meta http-equiv="Cache-Control" content="no-siteapp"/>
    <meta name="apple-mobile-web-app-capable" content="yes">
    <meta name="apple-mobile-web-app-status-bar-style" content="black">
    <meta name="apple-mobile-web-app-title" content="Amaze UI"/>
    <meta name="msapplication-TileColor" content="#0e90d2">
    <link rel="stylesheet" href="//cdn.amazeui.org/amazeui/2.4.2/css/amazeui.min.css">
</head>
<body>
<header class="am-topbar am-topbar-inverse">
    <div class="am-container am-cf">
        <div class="am-g">
            <h1 class="am-topbar-brand">
                <a href="#">离线下载</a>
            </h1>
            <!-- topbar-right -->
            <div class="am-topbar-right">
                <button class="am-btn am-btn-primary am-topbar-btn am-btn-sm">基于golang的下载神器</button>
            </div>
        </div>
    </div>

</header>

<div class="am-container">
    <div class="am-g">
        <div class="am-u-md-10 am-u-sm-centered">
            <legend>离线下载</legend>
            <div class="am-form-group">
                <p>欢迎使用！采用golang后台下载的方法实现简单的离线下载功能</p>
                <p>文件永久保留，下载空间为512M，若空间不足，请手工删除，且用且珍惜！</p>
                <div class="am-input-group">
						<span class="am-input-group-label">
							<i class="am-icon-cloud-download"></i>
						</span>
                    <input type="text" id="download" class="am-form-field" placeholder="请把文件的下载地址粘贴到这里,然后点击Enter即可。" />
                </div>
            </div>
            <button type="submit" class="am-btn am-btn-primary am-btn-block" onclick="download();">Enter</button>
        </div>
    </div>
    <div class="am-g">
        <div class="am-u-md-14 am-u-sm-centered">
            <table class="am-table am-table-hover">
                <thead>
                <tr>
                    <th>文件列表</th>
                </tr>
                </thead>
                <tbody id="downlist">

                </tbody>
            </table>
        </div>
    </div>
    <hr>

    <!-- 多说评论框 start -->
    <div class="ds-thread" data-thread-key="c7ef117e8fb5d1de76d99897dbbdad64" data-title="php离线下载" data-url="http://lyxz.gq/"></div>
    <!-- 多说评论框 end -->
    <!-- 多说公共JS代码 start (一个网页只需插入一次) -->
    <script type="text/javascript">
        var duoshuoQuery = {short_name:"lyxzgq"};
        (function() {
            var ds = document.createElement('script');
            ds.type = 'text/javascript';ds.async = true;
            ds.src = (document.location.protocol == 'https:' ? 'https:' : 'http:') + '//static.duoshuo.com/embed.js';
            ds.charset = 'UTF-8';
            (document.getElementsByTagName('head')[0]
            || document.getElementsByTagName('body')[0]).appendChild(ds);
        })();
    </script>
    <!-- 多说公共JS代码 end -->

</div>
<footer ata-am-widget="footer" class="am-footer am-footer-default" style="
    background:#555;
	font-color:#000;
">
    <style>
        .am-footer-default a{
            color:#999;
        }
    </style>
    <div class="am-footer-miscs">
        <p>友情链接 &nbsp;<a href="http://www.lanyus.com/" rel="friend" target="_blank">无心问世</a></p> &nbsp;
        <p><a href="https://github.com/ilanyu/offLineDownloader" target="_blank">关于</a></p>
        <p class="am-article-meta">
        </p><p><a href="#" target="_blank"><span class="am-icon-qq am-success" target="_blank"></span> 程序开发 无心问世</a></p>
        <br>
        <p>© 2016 无心问世. copyright</p>
    </div>
</footer>

<div class="am-modal am-modal-no-btn" tabindex="-1" id="downerror">
    <div class="am-modal-dialog">
        <div class="am-modal-hd">错误
            <a href="javascript: void(0)" class="am-close am-close-spin" data-am-modal-close>&times;</a>
        </div>
        <div class="am-modal-bd">
            下载出现错误，可能是不允许的文件后缀或下载空间已满。
        </div>
    </div>
</div>

<!--[if lt IE 9]>
<script src="http://libs.baidu.com/jquery/1.11.1/jquery.min.js"></script>
<script src="//cdn.staticfile.org/modernizr/2.8.3/modernizr.js"></script>
<script src="/js/polyfill/rem.min.js"></script>
<script src="http://libs.baidu.com/respond.js/1.4.2/respond.min.js"></script>
<script src="//cdn.amazeui.org/amazeui/2.4.2/js/amazeui.legacy.min.js"></script>
<![endif]-->

<!--[if (gte IE 9)|!(IE)]><!-->
<script src="http://libs.baidu.com/jquery/2.1.1/jquery.min.js"></script>
<script src="//cdn.amazeui.org/amazeui/2.4.2/js/amazeui.min.js"></script>
<!--<![endif]-->
<script>
    function writeDownList() {
        $.getJSON('./downlist',null,function (data) {
            var downlist = document.getElementById('downlist');
            downlist.innerHTML = '<tr><td>#</td><td>名称</td><td>修改时间</td><td>大小</td><td>操作</td></tr>';
            for (var i = 0 ; i < data.length ; i++) {
                downlist.innerHTML = downlist.innerHTML + '<tr><td>' + i + '</td><td><a href="./download/' + data[i]['name'] + '">' + data[i]['name'] + '</a></td><td>' + data[i]['mtime'] + '</td><td>' + data[i]['size'] + '</td><td><button type="button" class="am-btn" onclick="del(\'' + data[i]['name'] + '\');">删除</button></td></tr>';
            }
        });
    }
    function download() {
        var url = document.getElementById("download").value;
        $.get('./down',{'url':url},function (data) {
            if (data == 'success') {
                writeDownList();
            } else {
                $('#downerror').modal();
            }
        });
    }
    function del(file) {
        $.get('./delete',{'file':file},function (data) {
            writeDownList();
        });
    }
    writeDownList();
</script>
</body>
</html>
`)
		case string(ctx.Path()) == "/down":
			//点击“下载”触发，下载远程文件
			args := ctx.QueryArgs()
			url := string(args.Peek("url"))
			protocol := strings.Split(url, "://")
			if protocol[0] != "http" && protocol[0] != "https" && protocol[0] != "ftp" {
				ctx.WriteString("failed")
				log.Println(protocol)
				return
			}
			//
			if runtime.GOOS != "linux" {
				go download(url)
			} else {
				go downloadByWget(url)
			}
			ctx.WriteString("success")
			log.Println("visiter " + ctx.RemoteIP().String() + " use " + string(ctx.UserAgent()) + " downloaded " + url + " in " + time.Now().Format("2006-01-02 15:04:05"))
		case string(ctx.Path()) == "/delete":
			//删除文件
			args := ctx.QueryArgs()
			file := string(args.Peek("file"))
			file = strings.Replace(file, "/", "", -1)
			os.Remove("./download/" + file)
			ctx.WriteString("success")
		case bytes.HasPrefix(ctx.Path(), []byte("/download/")):
			fasthttp.FSHandler("./download", 1)(ctx)
		case string(ctx.Path()) == "/downlist":
			//生成文件列表
			dirList, _ := ioutil.ReadDir("./download")
			length := len(dirList)
			ctx.WriteString("[")
			for i := 0; i < length; i++ {
				ctx.WriteString(`{"name":"` + dirList[i].Name() + `","mtime":"` + dirList[i].ModTime().Format("2006-01-02 15:04:05") + `","size":"` + strconv.FormatInt(dirList[i].Size() / 1024 / 1024, 10) + ` MB"}`)
				if i != length - 1 {
					ctx.WriteString(",")
				}
			}
			ctx.WriteString("]")
		default:
			ctx.Error("not found", fasthttp.StatusNotFound)
		}
	}
	if fileInfo, err := os.Stat("./download"); err != nil || !fileInfo.IsDir() {
		os.Mkdir("./download", 0755)
		log.Println("mkDir dirctionary")
	}
	//bind := fmt.Sprintf("%s:%s", os.Getenv("OPENSHIFT_GO_IP"), os.Getenv("OPENSHIFT_GO_PORT"))
	bind := fmt.Sprintf("%s:%s", *host, *port)
	log.Printf("run on %s...\n", runtime.GOOS)
	log.Printf("listening on %s...\n", bind)
	if err := fasthttp.ListenAndServe(bind, m); err != nil {
		log.Fatal(err)
	}
}
