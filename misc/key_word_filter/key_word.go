package key_word_filter

func keyWords() []string {
	key_words := []string{"已售", "零售", "拦截器", "有效率", "赠送", "精仿", "dvd9", "逻辑搞死人", "大家大讲堂", "影印", "王小波", "世界图典", "电子书", "中国民意与公民社会", "卫星系统接收机", "邓小平时代", "我们要活得有尊严", "我们的", "青燈北岛", "红色记忆", "武器", "中国山川地图", "谎言与真相", "怎样学剪纸", "新闻界", "女人闺房音乐", "十四世", "漂流欲室", "脊梁", "绿色版", "保密知识", "绝对陷阱", "明报", "在历史的漩涡中", "工具软件", "软件安装", "电脑应用技巧", "弹道导弹", "性价比最高", "顶级", "最高级", "绝无仅有", "最具", "最高", "全国首家", "极端", "首选", "空前绝后", "绝对", "最大", "世界领先", "唯一", "巅峰", "顶峰", "中国+名牌", "最新", "最先进", "举世瞩目", "当当", "新东方", "新东方大愚", "托福考位", "自考专科", "火热促销", "西藏生死书", "附", "光盘", "21世纪", "教材", "学习卡", "教师课件", "普通高等教育", "十一五", "十二五", "普通高等教育", "附光盘", "国家级规划教材", "含CD光盘1张", "新世纪土木工程系列教材", "社会工作专业主干课教材", "自考教材", "燎原教育", "高等学校教材", "?", "附MP3光盘", "大愚英语学习丛书", "附光盘", "国家规划重点教材电视学系列教程", "《", "附MP3", "文都教育", "+带", "普通高等教育十一五国家级规划教材", "高等学校机械基础课程系列教材", "普通高等教育", "“", "”", "21世纪高等学校教材", "21世纪公共行政系列教材", "中国高等院校艺术设计学系列教材", "含光盘", "附光盘", "统计数据分析与应用丛书", "公共管理核心课程系列教材", "21世纪高等院校自动化专业", "发展学专业系列教材", "教材", "近机类专业适用", "》", "附赠MP3", "普通高等教育经济管理类专业", "附光盘1张", "普通高等教育", "普通高等教育十一五国家级规划教材", "普通高等教育", "“", "十二五", "”", "规划教材", "全国中医行业高等教育", "星火英语", "1CD", "（配光盘）", "普通高等学校体育教育专业主干课", "国家精品课程", "（21世纪经济学系列）", "21世纪信息通信系列", "（普通高校 管理学系列）", "（第九艺术学院游戏开发系列）", "高等医学院校有机化学立体化写组", "高等学校工程管理系列经典", "（工商管理经典译丛）", "全国高等院校财经类专业统编", "全国会计证从业资格考试", "高等院校安全工程专业：", "面向21世纪课程", "（普通高校经济及管理学科）", "（配光盘）", "（软件开发视频大讲堂）", "（高等学校经济与工商管理系列）", "新编服装院校系列", "国家重点图书出版", "（教育部国家精品课程", "（附CD-ROM光盘一张）", "（1cd）", "含DVD", "（21世纪房地产系列）", "（高等院校国际经济与贸易专业系列）", "经典高等学校英语专业系列", "（全国高等学校自动化专业系列）", "国家级", "（中国政法大学精品系列）", "（全国高等学校自动化专业系列）", "配MP3光盘", "含DVD光盘2张", "含DVD光盘1张", "（附CD光盘1张）", "（高等院校信", "（21世纪大学本科计算机专业系", "高等学校教学用书", "（21世纪工商管理特色）", "马克思主义理论研究和建设工程重点：", "高等学校会计学专业主要课程", "国家精品课程配套", "（21世纪高职高专艺术设计系列）", "（21世纪高等学校计算机）", "（配光盘）（21世纪大学本科计算", "（全国高等农业院校）", "高等院校园林专业系列", "面向21世纪课程普通高等学校社会工作专业主干课系列：", "（全国高等农业院校）", "（高等院校信息管理与信息", "（21世纪高等学校计算机应用技术）", "（本科临床配教）国家级", "高等学校工商管理类核心课程", "（B&E经济学系列）", "（附教师课件）", "（21世纪经济学）", "十五国家级", "（高校环境类）", "新世纪高校经济学管理学系列", "高等院校精品课系列", "（21世纪经管核心课程）", "教育技术学专业系列", "教育技术学专业系列", "国家级精品课程配套 ", "全国外经贸院校21世纪高职高专统编", "21世纪高等教育系列", "北京大学医学远程教育系列", "()", "郭敬明", "小时代", "（送教师课件）", "我用一生去寻找", "（教育部经济管理类主干课程）", "息与通信工程系列", "息管理与信息系统专业", "21世纪高等院校法学系列精品", "21世纪工商管理系列", "含1DVD", "附综合练习题及答案", "全国普通高等专科教育学类", "含MP3光盘", "附辅学光盘", "1张", "经济类核心课", "附CDROM一张", "附DVD", "附140元大礼包", "国家司法考试专题讲座", "九章丛书", "星火燎原", "高校经典同步辅导丛书", "含MP3一张", "MP3", "大学本科计算机专业系列", "普通高校", "高等学校计算机基础教育", "清华大学计算机", "工商管理优秀译丛", "营销学系列", "经济管理类课程", "税收系列", "金融系列", "法学系列", "教育部", "经济与管理系列", "公共行政系列", "公共管理硕士", "1DVD", "高职高专规划", "含盘一张", "含CD一张", "国内外经典辅导系列", "普通高等学校", "高等院校", "精品课程", "诺贝尔经济学奖获得者丛书", "工商管理经典译丛", "工商管理系列", "高等学校", "新世纪", "二十一世纪", "北京高等教育精品", "高等学校规划", "9787510443206", "伟大", "团购电话", "010-51571169", "扣扣", "V信", "v信", "电话", "QQ", "qq", "微信号", "微信", "博客", "糖溪帮探险记", "增高中药", "习主席", "性价比", "巨星", "奢侈", "习见平", "标杆人生", "主席酷爱", "规划", "“”", "系列", "面向", "课程", "面向", "课程", "附CD一张", "主席酷爱", "微博", "邮箱", "旺旺", "新浪", "新浪微博", "新浪旺信", "旺信", "5.3天天练", "如焉", "学历&证书", "成绩&单", "特惠品", "第一（NO.1", "Top1）", "万能等", "提优训练", "站台", "本杰明", "六A的力量", "七彩课堂", "亲笔", "签名", "签名版", "签名本", "珍藏", "珍藏版", "珍藏本", "附赠", "附盘", "附送", "ROM", "VCD", "DVD", "CD", "MP3", "光盘", "学习卡", "赠", "教学课件", "光碟", "限量", "限量版", "特供", "七彩课堂", "mp3", "七彩课堂", "QQ", "微博", "微信", "降脂", "万物", "人力资源", "育婴师", "点拨训练", "好卷", "星级口算", "点拨中考", "王凤仪嘉言录", "永不褪色", "考试必过", "立竿见影", "七彩课堂", "全国", "高职", "高专", "高校", "院校", "高等职业", "推荐", "专业", "中华民国", "七彩课堂", "华章教育", "考拉", "王长喜", "星火", "经典", "精品", "教学指导", "委员会", "类", "用书", "教学用书", "九五", "十五", "中国科学院", "核心", "经济类", "新航道", "圣才教育", "排毒", "人体写真", "经济管理类", "长喜英语", "课", "课程", "北京市", "主干", "经济与管理", "工商管理类", "双语教学", "工商管理", "经济管理", "国际化", "国际化管理", "中央财经大学", "学科", "重点", "本科", "：", ":", "·", "普通", "河北省", "国内外", "经典", "自学考试", "自考", "中学教师", "进修", "中学", "资格考试", "卫生职业", "中等职业", "国外", "人才培养", "精选", "译丛", "通用", "配套", "辅导", "教辅", "学前教育", "上海", "资格考试", "三级", "年", "你最近查看的商品和相关推荐", "原价", "最给力", "最便捷", "010-", "QQ群", "cmpedu", "央视", "CCTV", "中央电视台", "自然瘦", "女巫", "自然瘦", "第一", " ", "，", "（）"}
	return key_words
}