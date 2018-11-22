package main

import "fmt"

// 本地化资源

var locales map[string]map[string]string

func main() {
	// 本地化文本消息：
	// 建立需要的语言相应的map来维护一个key-value的关系，在输出之前按需从适合的map中去获取相应的文本
	locales = make(map[string]map[string]string, 2)
	// 创建英文本地map，存储英文流文本消息
	en := make(map[string]string, 5)
	en["pea"] = "pea"
	en["bean"] = "bean"
	// 创建中文本地map，存储中文流文本消息
	cn := make(map[string]string, 5)
	cn["pea"] = "豌豆"
	cn["bean"] = "黄豆"
	// 将英文区和中文区加入本地化资源中
	locales["zh_CN"] = cn
	locales["en_US"] = en
	// 测试加入的两个本地化资源是否可用
	lang1, lang2 := "zh_CN", "en_US"
	fmt.Println(msg(lang1, "pea"))	// 豌豆
	fmt.Println(msg(lang2, "pea"))	// pea
}

func msg(locale, key string) string {
	// 判断该区域的资源是否添加到本地化资源中
	if v, ok := locales[locale]; ok {
		// 判断该本地化资源中是否存在该key的资源
		if vv, ok := v[key]; ok {
			return vv
		}
	}
	// 如果不存在，则返回""
	return ""
}

/* 本地化时间和货币等：
 * 首先获取本地地区，然后获取当地时区的时间，然后在通过格式转换成当地时间格式
 *
 * 本地化视图和资源，目录结构如下：
 * views
 * |--en	//英文模板
 *	  |--images     //存储图片信息
 *	  |--js         //存储JS文件
 *	  |--css        //存储css文件
 *	  index.tpl     //用户首页
 *	  login.tpl     //登陆首页
 * |--zh-CN	//中文模板
 *	  |--images
 *	  |--js
 *	  |--css
 *	  index.tpl
 *	  login.tpl
 *
 * 加载时：
 * s1, _ := template.ParseFiles("views/"+lang+"/index.tpl")
 * VV.Lang=lang
 * s1.Execute(os.Stdout, VV)
 * 
 * index.html：
 * // js文件
 * <script type="text/javascript" src="views/{{.Lang}}/js/jquery/jquery-1.8.0.min.js"></script>
 * // css文件
 * <link href="views/{{.Lang}}/css/bootstrap-responsive.min.css" rel="stylesheet">
 * // 图片文件
 * <img src="views/{{.Lang}}/images/btn.png">
 *
 */
