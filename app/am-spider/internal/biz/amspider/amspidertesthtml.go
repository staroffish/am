package amspider

const miobtHtml = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
        <meta name="robots" content="index, follow">
    
        <title>＊MioBT＊</title>
        <meta name="keywords" content="动画 漫画 Galgame 游戏 动漫音乐 片源 RAW BT下载站 BT分享页 BT发布页 字幕组 下载 Torrent"/>
            <meta name="description" content="＊MioBT＊★就是＊MioBT＊喵～"/>
        
        <meta name="google-site-verification" content="ZxR2uuSG/HHhXN3nDGNPiea1PeboQ998w+XJWwg/hHLPb53BdDKHbBU3c8mcdZCo8l3WGtTWZvY4P3aF6TMA6LxZrqDFZyqmAvj8iBh8XuYAfEKbupciNlU1bgNkr5X1nj5Msxl7Li5YqsGf9k7N6OmxYkKLEYy2rKaP5gko0HIR9PjfYDgKlwwXA0siTi0DMXlJEwhpNptwjYH2r4bQyw8UIKUIqrfqjrexHhlBGTVEX0n0g8fSfmcAFrSmM/4y1OGQCkt4Qj9TfNyrTRaNfeBSQne/POVrbkcwvsDsvunVCWar4OTe0SA2NVkUeqUA6cL/faQ7yLCdD2cjmL3JpPqh/ErIRw3+c7nscfqfmbl2H3Qld7D1DjF2iMv37hKWBDLpOdCbNWn4zWiWDmIedw6U2UlD0bWlJbgQAREenUOeY68jNoyxF6HdAWCMm0DQ"/>
    
    <link rel="search" title="＊MioBT＊" type="application/opensearchdescription+xml" href="addon.php?r=opensearch.xml"/>

        <link rel="alternate" type="application/rss+xml" title="＊MioBT＊[RSS订阅]" href="rss.xml"/>
            <link rel="shortcut icon" href="images/favicon/miobt.ico" type="image/x-icon"/>
        <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/font-awesome@4.7.0/css/font-awesome.min.css">
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/balloon-css@0.5.0/balloon.min.css">
    <link type="text/css" rel="stylesheet" href="css/btmaster.css?v20081216"/>
    <link type="text/css" rel="stylesheet" href="css/common.css?20210905.1"/>
    <link type="text/css" rel="stylesheet" href="css/style_base.css?v20170813"/>
        <link type="text/css" rel="stylesheet" href="css/gg_base.css?v20081216"/>
                <style type="text/css">@import url("http://www.miobt.com/css/style_blue1.css");</style>
            <!-- <link type="text/css" rel="stylesheet" href="css/random/r1.css?v20081216" /> -->
            <script type="text/javascript" src="addon/assets/bower_components/jquery/jquery.min.js?2021"></script>
    <script type="text/javascript">jQuery.noConflict();</script>
    <script type="text/javascript" src="javascripts/mt-core.js?v20081130"></script>
    <script type="text/javascript" src="javascripts/btmaster.js?v20081130"></script>
    <script src="https://cdn.jsdelivr.net/npm/jquery.cookie@1.4.1/jquery.cookie.min.js"></script>
    <script src="addon/assets/common/ads.js?20190526.2"></script>
    <script src="addon/assets/common/gg.js?20210909"></script>
    <script type="text/javascript">
        Config['url_rewrite'] = true;
        Config['in_script'] = "index";
        Config['link_board'] = {"datasets":{"gg_link_board":{"data_src":"text_vars","sections":{"gg_header_link_board":{"enabled":false,"columns":6,"text_align":"left"}}},"gg_side_board":{"data_src":"api_data","sections":{"gg_side_board_games":{"enabled":true,"limit":6,"more_link":["https:\/\/galgame.erolabs.net"]}}}}};
        Config['mika_mode'] = {
            'enabled': true        };
        Config['user_script'] = {
            'enabled': true,
            'revision': '20181120.2',
            'platform': 'desktop'
        };
        if (typeof Config['gg_blocked'] === 'undefined') {
            Config['gg_blocked'] = true;
        }
    </script>
    <script type="text/javascript">
        // No-Opener-No-Phishers
        jQuery('[target="_blank"]').attr("rel", "noopener noreferrer");
    </script>
    
</head>

<body>

<script type="text/javascript">
    Config['search_tip'] = "_(:з」∠)_";

    function doSearch() {
        var a = $F("topsearch");
        if (!a || Config['search_tip'] == a) {
            Window.alert(Config['search_tip']);
            return false;
        }
        window.location.href = "search.php?keyword=" + encodeURI(a);
    }

    function searchTip(a) {
        if (a == 1) {
            if (!$F("topsearch")) {
                $("topsearch").value = Config['search_tip'];
            }
        } else {
            if ($F("topsearch") == Config['search_tip']) {
                $("topsearch").value = "";
            }
            $("topsearch").onblur = function () {
                searchTip(1);
            };
        }
    }
</script>

<div id="btm" style="min-width:1230px; max-width: 80%;">
<div class="header">
    <div class="top_nav">
        <p class="left">
            <a href="./" style="display: none;">＊MioBT＊</a>
                    </p>

        <p class="right">
            <!--<a href="./">茵蒂克丝</a> |-->
            <a href="addon.php?r=tagcloud"            target="_blank">
            标签云</a> | <a href="addon.php?r=document/view&amp;page=ui-preference"            target="_blank">
            页面选项</a> | <a href="addon.php?r=document/view&page=group-application"            target="_blank">
            开通联萌</a> | <a href="addon.php?r=document/view&page=abuse"            target="_blank">
            滥用举报</a> | 
            <script type="text/javascript" src="javascripts/locale.js?v20150715"></script> |
                        <a href="user.php">会员面板</a> |            <!--<a href="search.php">高级搜索</a> | --><a href="user.php?o=upload" style="font-weight: bold;" target="_blank">我要发布</a>                                    </p>
    </div>
    <div class="func">
        <div class="topsearch"><input type="text" id="topsearch" onkeypress="doSearchEvent(event);" value="_(:з」∠)_" onfocus="searchTip(0);" /><a href="javascript:void(0);" onclick="doSearch();return false;">搜索</a></div>
        <script type="text/javascript">var kcount=0, kheight=18, kst=0;</script>
        <div class="hotsearch" style="display: none;">
            <span class="left">人气搜索:</span>
            <div id="searchkey">
                <div id="searchkeybody">
                                </div>
            </div>
        </div>
    </div>
</div>

<div class="nav" id="smenu">
    <ul>
        <li><a href="./">全部</a></li>
                <li><a href="sort-1-1.html">动画</a></li>
                <li><a href="sort-2-1.html">漫画</a></li>
                <li><a href="sort-3-1.html">音乐</a></li>
                <li><a href="sort-4-1.html">周边</a></li>
                <li><a href="sort-5-1.html">其它</a></li>
                <li><a href="sort-6-1.html">Raw</a></li>
        
        <li><a href="animovie-1.html">Ova</a></li>
        <li><a href="complete-1.html">合集</a></li>
        <li><a href="discuss-1.html">讨论</a></li>
                <li><a href="cloudfile-1.html">网盘</a></li>
        <li><a href="addon.php?r=bgmlist">片库</a></li>
        <li><a href="addon.php?r=tagcloud">更多</a></li>
            </ul>
</div>

<script type="text/javascript">btmenu.init('smenu', 'auto');</script>

<div class="clear text_center" style="margin-bottom:10px;"><style type="text/css">
    .gg_nutaku-qw, .gg_nutaku-qnzd, .gg_nutaku-qnzl {
        display: none;
    }
    .gg_nutaku-ph {
        display: block;
    }
</style>











                        <div class="gg_canvas gg_canvas-hidden" style="margin-top: 10px;">
                <a gg_url="https://l.tapdb.net/aSHVirFb"><img src="public/imgbed/gg/lsjg/qjsg_1200x90.jpg?20220123?20210410"></a>
            </div>
                                <div class="gg_canvas gg_canvas-hidden" style="margin-top: 10px;">
                <a gg_url="https://l.tyrantdb.com/gwBmUfoi"><img src="public/imgbed/gg/54647/CQ_1200x90_SC_B.jpg?20210915"></a>
            </div>
                                <div class="gg_canvas gg_canvas-hidden" style="margin-top: 10px;">
                <a gg_url="http://jg.37gowan.com/stf/visitor.html?id=146&s=3899&c={uid}"><img src="public/imgbed/gg/h5mishu/2-1200x90.png?20200910"></a>
            </div>
                                <div class="gg_canvas gg_canvas-hidden" style="margin-top: 10px;">
                <a gg_url="https://fgoatt79.apple1818.net"><img src="public/imgbed/gg/jgg/FGO_79_1200_90.png?20220121?20210410"></a>
            </div>
                                <div class="gg_canvas gg_canvas-hidden" style="margin-top: 10px;">
                <a gg_url="https://shop119340084.taobao.com/shop/view_shop.htm?spm=a313o.201708ban.category.d53.64f0197aCpRvXk&mytmenu=mdianpu&user_number_id=1965847533&mm_sycmid=1_115618_1e61a6a60d1eef8cd28b57fb4ba1d321"><img src="public/imgbed/gg/fj1157524220/1200x90_fj1157524220_20220220.jpg?20210512"></a>
            </div>
                                <div class="gg_canvas gg_canvas-hidden" style="margin-top: 10px;">
                <a gg_url="https://l.tyrantdb.com/pzT5rE6A"><img src="public/imgbed/gg/h365/JTW_20220124_1200X90_il08.jpg?20220128?20210410"></a>
            </div></div>



    <div class="clear text_center" style="margin-bottom:10px;"><div id="gg_header_link_board" class="main" style="display: none;"></div><script type="text/javascript">gg_link_board_20210903('gg_link_board', 'gg_header_link_board');</script></div>

<div class="main">





<div class="box clear rounded">
    
    <div class="clear">
        <h2 class="title list_style_table_hidden_head_title" style="display: none">番组表</h2>
        <table class="list_style table_fixed">
            <tbody class="tbody">
                        <tr class="alt1">
                <td style="width: 70px;">
                    <a href="addon.php?r=bangumi/table" target="_blank" style="padding: 5px; background-color: #F0F0F0; border-radius: 5px; text-decoration: none;">本季番</a>
                </td>
                <td style="text-align:left; width: 100%; overflow: visible;">
                                        <a href="addon.php?r=bangumi/table&bgm_id=2021q4" target="_blank" style="padding: 5px; background-color: #F0F0F0; border-radius: 5px; margin-right: 8px; text-decoration: none;">21年10月番</a>
                                        <a href="addon.php?r=bangumi/table&bgm_id=2021q3" target="_blank" style="padding: 5px; background-color: #F0F0F0; border-radius: 5px; margin-right: 8px; text-decoration: none;">21年07月番</a>
                                        <a href="addon.php?r=bangumi/table&bgm_id=2021q2" target="_blank" style="padding: 5px; background-color: #F0F0F0; border-radius: 5px; margin-right: 8px; text-decoration: none;">21年04月番</a>
                                        <a href="addon.php?r=bangumi/table&bgm_id=2021q1" target="_blank" style="padding: 5px; background-color: #F0F0F0; border-radius: 5px; margin-right: 8px; text-decoration: none;">21年01月番</a>
                                        <a href="addon.php?r=bangumi/table&bgm_id=2020q4" target="_blank" style="padding: 5px; background-color: #F0F0F0; border-radius: 5px; margin-right: 8px; text-decoration: none;">20年10月番</a>
                                        <a href="addon.php?r=bangumi/table&bgm_id=2020q3" target="_blank" style="padding: 5px; background-color: #F0F0F0; border-radius: 5px; margin-right: 8px; text-decoration: none;">20年07月番</a>
                                        <a href="addon.php?r=bangumi/table&bgm_id=2020q2" target="_blank" style="padding: 5px; background-color: #F0F0F0; border-radius: 5px; margin-right: 8px; text-decoration: none;">20年04月番</a>
                                        <a href="addon.php?r=bangumi/table&bgm_id=2020q1" target="_blank" style="padding: 5px; background-color: #F0F0F0; border-radius: 5px; margin-right: 8px; text-decoration: none;">20年01月番</a>
                                        <a href="addon.php?r=bangumi/table&bgm_id=2019q4" target="_blank" style="padding: 5px; background-color: #F0F0F0; border-radius: 5px; margin-right: 8px; text-decoration: none;">19年10月番</a>
                    
                    <a href="addon.php?r=document/view&page=guestbook/bgm-feedback" target="_blank" style="font-weight:bold; color:green; padding: 5px; background-color: #F0F0F0; border-radius: 5px; margin-right: 8px; text-decoration: none;">番组报错</a>
                    <a href="addon.php?r=document/view&page=guestbook/bgm-hitokoto" target="_blank" style="font-weight:bold; color:blue; padding: 5px; background-color: #F0F0F0; border-radius: 5px; margin-right: 8px; text-decoration: none;">短评投稿</a>
                    <a href="addon.php?r=bangumi/table" target="_blank" style="padding: 5px; background-color: #F0F0F0; border-radius: 5px; margin-right: 8px; text-decoration: none;">...更多往期▲</a>
                </td>
            </tr>
                                                            <tr class="alt2">
                            <td style="width: 70px; font-weight: bold;">
                                            <a href="today-1.html">●今天</a>
                                    </td>
                <td style="text-align:left; width: 100%; overflow: visible;">
                                        <a target="_blank" href="search.php?keyword=%E6%B5%B7%E8%B4%BC%E7%8E%8B" style="margin:0 8px 0 0; padding: 5px 0 5px 0; font-weight: bold;" class="bgm_score">
                    海贼王</a>
                                        <a target="_blank" href="search.php?keyword=1%E6%9C%88%E6%96%B0%E7%95%AA%E2%86%92" style="margin:0 8px 0 0; padding: 5px 0 5px 0; font-weight: bold;" class="bgm_score">
                    1月新番→</a>
                                        <a target="_blank" href="search.php?keyword=%E8%BF%9B%E5%87%BB%E7%9A%84%E5%B7%A8%E4%BA%BA" style="margin:0 8px 0 0; padding: 5px 0 5px 0; font-weight: bold;" class="bgm_score"                    data-balloon-length="medium" data-balloon="jsc 我吃了nm→2022年1月9日起 每周日23:05" data-balloon-pos="up">
                    进击的巨人 最终季 分割下半</a>
                                        <a target="_blank" href="search.php?keyword=%E7%8E%B0%E5%AE%9E%E4%B8%BB%E4%B9%89%E5%8B%87%E8%80%85" style="margin:0 8px 0 0; padding: 5px 0 5px 0; font-weight: bold;" class="bgm_score"                    data-balloon-length="medium" data-balloon="穿越/内政/改革→2022年1月9日起 每周日00:30" data-balloon-pos="up">
                    现实主义勇者的王国再建记 第二季</a>
                                        <a target="_blank" href="search.php?keyword=%E9%94%88%E8%89%B2%E9%93%A0%E7%94%B2" style="margin:0 8px 0 0; padding: 5px 0 5px 0; font-weight: bold;" class="bgm_score"                    data-balloon-length="medium" data-balloon="奇幻/战国/战争→2022年1月9日起 每周日21:00" data-balloon-pos="up">
                    锈色铠甲 -黎明-</a>
                                        <a target="_blank" href="search.php?keyword=%E7%8E%AB%E7%91%B0%E4%B9%8B%E7%8E%8B%E7%9A%84%E8%91%AC%E7%A4%BC" style="margin:0 8px 0 0; padding: 5px 0 5px 0; font-weight: bold;" class="bgm_score"                    data-balloon-length="medium" data-balloon="历史/战争/情感→2022年1月9日起 每周日21:30" data-balloon-pos="up">
                    玫瑰之王的葬礼</a>
                                        <a target="_blank" href="search.php?keyword=Futsal+Boys" style="margin:0 8px 0 0; padding: 5px 0 5px 0; font-weight: bold;" class="bgm_score"                    data-balloon-length="medium" data-balloon="校园/5人制足球→2022年1月9日起 每周日22:00" data-balloon-pos="up">
                    室内足球少年 Futsal Boys!!!!!</a>
                                        <a target="_blank" href="search.php?keyword=%E4%BD%90%E4%BD%90%E6%9C%A8%E4%B8%8E%E5%AE%AB%E9%87%8E" style="margin:0 8px 0 0; padding: 5px 0 5px 0; font-weight: bold;" class="bgm_score"                    data-balloon-length="medium" data-balloon="校园/日常/耽美→2021年1月9日起 每周日23:30" data-balloon-pos="up">
                    佐佐木与宫野</a>
                                        <a target="_blank" href="search.php?keyword=%E6%80%AA%E4%BA%BA%E5%BC%80%E5%8F%91%E9%83%A8" style="margin:0 8px 0 0; padding: 5px 0 5px 0; font-weight: bold;" class="bgm_score"                    data-balloon-length="medium" data-balloon="反派/职场/喜剧→2022年1月9日起 每周日01:00" data-balloon-pos="up">
                    怪人开发部的黑井津</a>
                                        <a target="_blank" href="search.php?keyword=%E2%86%901%E6%9C%88%E6%96%B0%E7%95%AA" style="margin:0 8px 0 0; padding: 5px 0 5px 0; font-weight: bold;" class="bgm_score">
                    ←1月新番</a>
                                        <a target="_blank" href="search.php?keyword=%E9%AC%BC%E7%81%AD%E4%B9%8B%E5%88%83" style="margin:0 8px 0 0; padding: 5px 0 5px 0; font-weight: bold;" class="bgm_score"                    data-balloon-length="medium" data-balloon="战斗 热血 霸权→12月5日起 每周日23:15" data-balloon-pos="up">
                    鬼灭之刃 游郭篇</a>
                                        <a target="_blank" href="search.php?keyword=%E9%B2%81%E9%82%A6%E4%B8%89%E4%B8%96%09" style="margin:0 8px 0 0; padding: 5px 0 5px 0; font-weight: bold;" class="bgm_score"                    data-balloon-length="medium" data-balloon="50周年纪念作→10月10日起 每周日00:55" data-balloon-pos="up">
                    鲁邦三世 PART6</a>
                                        <a target="_blank" href="search.php?keyword=%E6%95%B0%E7%A0%81%E5%AE%9D%E8%B4%9D" style="margin:0 8px 0 0; padding: 5px 0 5px 0; font-weight: bold;" class="bgm_score"                    data-balloon-length="medium" data-balloon="→10月3日起 每周日09:00" data-balloon-pos="up">
                    数码宝贝幽灵游戏</a>
                                        <a target="_blank" href="search.php?keyword=%E7%BE%8E%E5%A6%99%E9%AD%94%E6%B3%95" style="margin:0 8px 0 0; padding: 5px 0 5px 0; font-weight: bold;" class="bgm_score"                    data-balloon-length="medium" data-balloon="美妙系列十周年纪念作品10月3日起 每周日10:00" data-balloon-pos="up">
                    美妙魔法</a>
                                        <a target="_blank" href="search.php?keyword=%E7%89%B9%E6%96%AF%E6%8B%89%E7%AC%94%E8%AE%B0" style="margin:0 8px 0 0; padding: 5px 0 5px 0; font-weight: bold;" class="bgm_score"                    data-balloon-length="medium" data-balloon="烂  瞎眼 光污染→10月3日起 每周日23:00" data-balloon-pos="up">
                    特斯拉笔记</a>
                                        <a target="_blank" href="search.php?keyword=%E5%8D%9A%E4%BA%BA%E4%BC%A0" style="margin:0 8px 0 0; padding: 5px 0 5px 0; font-weight: bold;" class="bgm_score">
                    博人传</a>
                                        <a target="_blank" href="search.php?keyword=%E5%85%89%E4%B9%8B%E7%BE%8E%E5%B0%91%E5%A5%B3" style="margin:0 8px 0 0; padding: 5px 0 5px 0; font-weight: bold;" class="bgm_score"                    data-balloon-length="medium" data-balloon="原创/Q娃→【集数：年番?】2月28日起每周日/8:30" data-balloon-pos="up">
                    Tropical-Rouge!光之美少女</a>
                                        <a target="_blank" href="search.php?keyword=%E7%94%9C%E6%A2%A6%E7%8C%AB" style="margin:0 8px 0 0; padding: 5px 0 5px 0; font-weight: bold;" class="bgm_score"                    data-balloon-length="medium" data-balloon="原创/子供向4月11日起每周日/10:30" data-balloon-pos="up">
                    甜梦猫 MIX!</a>
                                        <a target="_blank" href="search.php?keyword=%E6%95%B0%E7%A0%81%E5%AE%9D%E8%B4%9D%E5%A4%A7%E5%86%92%E9%99%A9%EF%BC%9A" style="margin:0 8px 0 0; padding: 5px 0 5px 0; font-weight: bold;" class="bgm_score"                    data-balloon-length="medium" data-balloon="童年回忆*2→2020年4月5日起 每周日9:00" data-balloon-pos="up">
                    数码宝贝大冒险：</a>
                    
                </td>
            </tr>
                                                            <tr class="alt1">
                            <td style="width: 70px;">
                                                                    星期一                                                                                                                                                                                                                                                            </td>
                <td style="text-align:left; width: 100%; overflow: visible;">
                                        <a target="_blank" href="search.php?keyword=1%E6%9C%88%E6%96%B0%E7%95%AA%E2%86%92" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score">
                    1月新番→</a>
                                        <a target="_blank" href="search.php?keyword=%E5%85%AC%E4%B8%BB%E8%BF%9E%E6%8E%A5" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score"                    data-balloon-length="medium" data-balloon="手游改 冒险/日常/喜剧→2022年1月10日起 每周一23:00" data-balloon-pos="up">
                    公主连结Re:Dive 第二季</a>
                                        <a target="_blank" href="search.php?keyword=%E9%A3%9F%E9%94%88%E6%9C%AB%E4%B8%96%E5%BD%95" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score"                    data-balloon-length="medium" data-balloon="废土/冒险/动作→2022年1月10日起 每周一23:30" data-balloon-pos="up">
                    食锈末世录</a>
                                        <a target="_blank" href="search.php?keyword=%E7%8E%8B%E5%AD%90%E7%9A%84%E6%9C%AC%E5%91%BD" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score"                    data-balloon-length="medium" data-balloon="本季度僧侣档→2022年1月10日起 每周一00:00" data-balloon-pos="up">
                    王子的本命是恶役千金</a>
                                        <a target="_blank" href="search.php?keyword=%E6%9A%97%E8%8A%9D%E5%B1%85" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score">
                    暗芝居 10期</a>
                                        <a target="_blank" href="search.php?keyword=TRIBE+NINE" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score"                    data-balloon-length="medium" data-balloon="科幻/决斗棒球→2022年1月10日起 每周一21:30" data-balloon-pos="up">
                    夜街酷斗 TRIBE NINE</a>
                                        <a target="_blank" href="search.php?keyword=%E2%86%901%E6%9C%88%E6%96%B0%E7%95%AA" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score">
                    ←1月新番</a>
                    
                </td>
            </tr>
                                                            <tr class="alt2">
                            <td style="width: 70px;">
                                                                                            星期二                                                                                                                                                                                                                                    </td>
                <td style="text-align:left; width: 100%; overflow: visible;">
                                        <a target="_blank" href="search.php?keyword=1%E6%9C%88%E6%96%B0%E7%95%AA%E2%86%92" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score">
                    1月新番→</a>
                                        <a target="_blank" href="search.php?keyword=Fantasy+Bishoujo" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score"                    data-balloon-length="medium" data-balloon="穿越/冒险/性转→2022年1月11日起 每周二23:00" data-balloon-pos="up">
                    与变成了异世界美少女的大叔一起冒险</a>
                                        <a target="_blank" href="search.php?keyword=%E5%A4%A9%E6%89%8D%E7%8E%8B%E5%AD%90" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score"                    data-balloon-length="medium" data-balloon="喜剧/治国/恋爱→2022年1月11日起 每周二21:30" data-balloon-pos="up">
                    天才王子的赤字国家重生术</a>
                                        <a target="_blank" href="search.php?keyword=%E8%87%AA%E7%A7%B0%E8%B4%A4%E8%80%85%E5%BC%9F%E5%AD%90%E7%9A%84%E8%B4%A4%E8%80%85" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score"                    data-balloon-length="medium" data-balloon="穿越/性转/公路→2022年1月11日起 每周二23:30" data-balloon-pos="up">
                    自称贤者弟子的贤者</a>
                                        <a target="_blank" href="search.php?keyword=%E5%B9%BB%E6%83%B3%E4%B8%89%E5%9B%BD%E5%BF%97+-%E5%A4%A9%E5%85%83%E7%81%B5%E5%BF%83%E8%AE%B0-" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score"                    data-balloon-length="medium" data-balloon="奇幻/鬼怪/战斗2022年1月11日起 每周二01:00" data-balloon-pos="up">
                    幻想三国志 -天元灵心记-</a>
                                        <a target="_blank" href="search.php?keyword=%E2%86%901%E6%9C%88%E6%96%B0%E7%95%AA" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score">
                    ←1月新番</a>
                    
                </td>
            </tr>
                                                            <tr class="alt1">
                            <td style="width: 70px;">
                                                                                                                    星期三                                                                                                                                                                                                            </td>
                <td style="text-align:left; width: 100%; overflow: visible;">
                                        <a target="_blank" href="search.php?keyword=1%E6%9C%88%E6%96%B0%E7%95%AA%E2%86%92" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score">
                    1月新番→</a>
                                        <a target="_blank" href="search.php?keyword=Koroshi+Ai" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score"                    data-balloon-length="medium" data-balloon="杀手/悬疑/恋爱2022年1月12日起 每周三23:00" data-balloon-pos="up">
                    相爱相杀 （杀爱）</a>
                                        <a target="_blank" href="search.php?keyword=%E4%B8%9C%E6%96%B9%E5%B0%91%E5%B9%B4" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score"                    data-balloon-length="medium" data-balloon="架空战国时代 妖魔/冒险/动作→2022年1月5日起 每周三23:00" data-balloon-pos="up">
                    东方少年</a>
                                        <a target="_blank" href="search.php?keyword=Hakozume" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score"                    data-balloon-length="medium" data-balloon="警察/日常/恋爱2022年1月5日起 每周三22:30" data-balloon-pos="up">
                    女子警察的逆袭</a>
                                        <a target="_blank" href="search.php?keyword=Leadale" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score"                    data-balloon-length="medium" data-balloon="穿越/带娃/公路→2022年1月5日起 每周三22:30" data-balloon-pos="up">
                    里亚德录大地</a>
                                        <a target="_blank" href="search.php?keyword=%E5%BD%A9%E7%BB%BF" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score"                    data-balloon-length="medium" data-balloon="|校园/乐队/日常→2022年1月5日起 每周三00:00" data-balloon-pos="up">
                    彩绿</a>
                                        <a target="_blank" href="search.php?keyword=%E2%86%901%E6%9C%88%E6%96%B0%E7%95%AA" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score">
                    ←1月新番</a>
                    
                </td>
            </tr>
                                                            <tr class="alt2">
                            <td style="width: 70px;">
                                                                                                                                            星期四                                                                                                                                                                                    </td>
                <td style="text-align:left; width: 100%; overflow: visible;">
                                        <a target="_blank" href="search.php?keyword=1%E6%9C%88%E6%96%B0%E7%95%AA%E2%86%92" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score">
                    1月新番→</a>
                                        <a target="_blank" href="search.php?keyword=%E4%B8%9C%E4%BA%AC24%E5%8C%BA" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score"                    data-balloon-length="medium" data-balloon="悬疑/英雄剧→2022年1月6日起 每周四00:30" data-balloon-pos="up">
                    东京24区</a>
                                        <a target="_blank" href="search.php?keyword=%E6%9C%80%E6%B8%B8%E8%AE%B0" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score"                    data-balloon-length="medium" data-balloon="冒险/架空改编2022年1月6日起 每周四23:00" data-balloon-pos="up">
                    最游记RELOAD -ZEROIN-</a>
                                        <a target="_blank" href="search.php?keyword=%E5%B0%8F%E5%B0%8F%E4%B8%96%E7%95%8C" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score"                    data-balloon-length="medium" data-balloon="→2022年1月13日起 每周四18:00" data-balloon-pos="up">
                    小小世界  初音未来/迷你动画作品</a>
                                        <a target="_blank" href="search.php?keyword=%E2%86%901%E6%9C%88%E6%96%B0%E7%95%AA" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score">
                    ←1月新番</a>
                                        <a target="_blank" href="search.php?keyword=BanG+Dream" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score"                    data-balloon-length="medium" data-balloon="迷你动画第三季10月7日起 每周四22:00" data-balloon-pos="up">
                    BanG Dream! 少女乐团派对!☆PICO Fever!</a>
                                        <a target="_blank" href="search.php?keyword=%E9%80%9A%E7%81%B5%E7%8E%8B" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score"                    data-balloon-length="medium" data-balloon="漫改/热血/爷青回/安娜姐/船到桥头自然直→4月1日起每周四/17:55" data-balloon-pos="up">
                    通灵王（第二作）</a>
                    
                </td>
            </tr>
                                                            <tr class="alt1">
                            <td style="width: 70px;">
                                                                                                                                                                    星期五                                                                                                                                                            </td>
                <td style="text-align:left; width: 100%; overflow: visible;">
                                        <a target="_blank" href="search.php?keyword=1%E6%9C%88%E6%96%B0%E7%95%AA%E2%86%92" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score">
                    1月新番→</a>
                                        <a target="_blank" href="search.php?keyword=Slow+Loop" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score"                    data-balloon-length="medium" data-balloon="日常/姐妹/钓鱼→2022年1月7日起 每周五21:00" data-balloon-pos="up">
                    Slow Loop女孩的钓鱼慢活</a>
                                        <a target="_blank" href="search.php?keyword=%E5%B9%B3%E5%87%A1%E8%81%8C%E4%B8%9A" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score"                    data-balloon-length="medium" data-balloon="群穿/后宫/傲天→2022年1月14日起 每周五00:30" data-balloon-pos="up">
                    平凡职业造就世界最强 第二季</a>
                                        <a target="_blank" href="search.php?keyword=%E7%BB%88%E6%9C%AB%E7%9A%84%E5%90%8E%E5%AE%AB" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score"                    data-balloon-length="medium" data-balloon="近未来/种马/后宫→2021年1月7日起 每周五20:30" data-balloon-pos="up">
                    终末的后宫</a>
                                        <a target="_blank" href="search.php?keyword=%E5%B7%9D%E5%B0%BB%E5%B0%8F%E7%8E%89" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score"                    data-balloon-length="medium" data-balloon="日常喜剧/美食  泡面番→2022年1月14日起 每周五00:25" data-balloon-pos="up">
                    我是川尻小玉～生活黑客的糜烂生活～</a>
                                        <a target="_blank" href="search.php?keyword=%E2%86%901%E6%9C%88%E6%96%B0%E7%95%AA" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score">
                    ←1月新番</a>
                                        <a target="_blank" href="search.php?keyword=%E7%99%BD%E9%87%91%E7%BB%88%E5%B1%80" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score"                    data-balloon-length="medium" data-balloon="悬疑 推理 哲学 血腥→10月8日起 每周五01:28" data-balloon-pos="up">
                    白金终局</a>
                                        <a target="_blank" href="search.php?keyword=%E5%9B%BD%E7%8E%8B%E6%8E%92%E5%90%8D" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score"                    data-balloon-length="medium" data-balloon="治愈 成长 冒险 黑马→10月15日起 每周五01:55" data-balloon-pos="up">
                    国王排名</a>
                                        <a target="_blank" href="search.php?keyword=%E5%A6%96%E6%80%AA%E6%89%8B%E8%A1%A8" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score"                    data-balloon-length="medium" data-balloon="游戏改→4月9日起每周五/18:25" data-balloon-pos="up">
                    妖怪手表♪</a>
                                        <a target="_blank" href="search.php?keyword=%E6%96%B0%E5%B9%B2%E7%BA%BF%E5%8F%98%E5%BD%A2%E6%9C%BA%E5%99%A8%E4%BA%BA" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score"                    data-balloon-length="medium" data-balloon="原创【集数：年番?】4月9日起每周五/19:25" data-balloon-pos="up">
                    新干线变形机器人 SHINKALION Z</a>
                                        <a target="_blank" href="search.php?keyword=Doraemon" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score">
                    哆啦A梦</a>
                                        <a target="_blank" href="search.php?keyword=%E8%9C%A1%E7%AC%94%E5%B0%8F%E6%96%B0" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score">
                    蜡笔小新</a>
                                        <a target="_blank" href="search.php?keyword=%E5%AE%A0%E7%89%A9%E5%B0%8F%E7%B2%BE%E7%81%B5" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score">
                    宠物小精灵</a>
                    
                </td>
            </tr>
                                                            <tr class="alt2">
                            <td style="width: 70px;">
                                                                                                                                                                                            星期六                                                                                                                                    </td>
                <td style="text-align:left; width: 100%; overflow: visible;">
                                        <a target="_blank" href="search.php?keyword=1%E6%9C%88%E6%96%B0%E7%95%AA%E2%86%92" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score">
                    1月新番→</a>
                                        <a target="_blank" href="search.php?keyword=%E9%AB%98%E6%9C%A8%E5%90%8C%E5%AD%A6" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score"                    data-balloon-length="medium" data-balloon="校园/恋爱/狗粮→2022年1月8日起 每周六00:25" data-balloon-pos="up">
                    擅长捉弄人的高木同学 第三季</a>
                                        <a target="_blank" href="search.php?keyword=Sono+Bisque+Doll" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score"                    data-balloon-length="medium" data-balloon="校园/恋爱/Cos2022年1月8日起 每周六23:00" data-balloon-pos="up">
                    更衣人偶坠入爱河</a>
                                        <a target="_blank" href="search.php?keyword=RYMAN" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score"                    data-balloon-length="medium" data-balloon="社畜/羽毛球→2022年1月22日起 每周日00:30" data-balloon-pos="up">
                    白领羽球部 RYMANS CLUB</a>
                                        <a target="_blank" href="search.php?keyword=%E5%B0%91%E5%A5%B3%E5%89%8D%E7%BA%BF" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score"                    data-balloon-length="medium" data-balloon="手游改 科幻/废土/枪械→2022年1月8日起 每周六00:00" data-balloon-pos="up">
                    少女前线</a>
                                        <a target="_blank" href="search.php?keyword=%E6%98%8E%E6%97%A5%E5%90%8C%E5%AD%A6%E7%9A%84%E6%B0%B4%E6%89%8B%E6%9C%8D" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score"                    data-balloon-length="medium" data-balloon="校园/乡村/日常→2022年1月8日起 每周六23:30" data-balloon-pos="up">
                    明日同学的水手服</a>
                                        <a target="_blank" href="search.php?keyword=%E5%A4%B1%E6%A0%BC%E7%BA%B9" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score"                    data-balloon-length="medium" data-balloon="奇幻/转世/傲天→2022年1月8日起 每周六21:00" data-balloon-pos="up">
                    失格纹的最强贤者</a>
                                        <a target="_blank" href="search.php?keyword=CUE" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score"                    data-balloon-length="medium" data-balloon="声优/业界/成长→2022年1月8日起 每周六00:55" data-balloon-pos="up">
                    CUE!</a>
                                        <a target="_blank" href="search.php?keyword=%E2%86%901%E6%9C%88%E6%96%B0%E7%95%AA" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score">
                    ←1月新番</a>
                                        <a target="_blank" href="search.php?keyword=%E5%8D%8A%E5%A6%96%E7%9A%84%E5%A4%9C%E5%8F%89%E5%A7%AC+" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score"                    data-balloon-length="medium" data-balloon="圈钱 犬薇 消费情怀→10月2日起 每周六17:30" data-balloon-pos="up">
                    半妖的夜叉姬 贰之章</a>
                                        <a target="_blank" href="search.php?keyword=%E5%8B%87%E8%80%85%E6%96%97%E6%81%B6%E9%BE%99" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score"                    data-balloon-length="medium" data-balloon="漫改/奇幻/经典→【集数：100】10月3日起每周六09:30" data-balloon-pos="up">
                    勇者斗恶龙 达伊的大冒险</a>
                                        <a target="_blank" href="search.php?keyword=%E5%90%8D%E4%BE%A6%E6%8E%A2%E6%9F%AF%E5%8D%97" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score"                    data-balloon-length="medium" data-balloon="日本人口还够吗？" data-balloon-pos="up">
                    名侦探柯南</a>
                    
                </td>
            </tr>
                                                            <tr class="alt1">
                            <td style="width: 70px;">
                                                                                                                                                                                                                                                                    剧场版                                                            </td>
                <td style="text-align:left; width: 100%; overflow: visible;">
                                        <a target="_blank" href="search.php?keyword=%E5%89%A7%E5%9C%BA%E7%89%88" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score">
                    BD已发售→</a>
                                        <a target="_blank" href="search.php?keyword=%E5%89%A7%E5%9C%BA%E7%89%88" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score">
                    即将发售BD→</a>
                    
                </td>
            </tr>
                                                            <tr class="alt2">
                            <td style="width: 70px;">
                                                                                                                                                                                                                                                                                            特殊放送                                    </td>
                <td style="text-align:left; width: 100%; overflow: visible;">
                                        <a target="_blank" href="search.php?keyword=%E7%A9%BA%E8%89%B2Utility" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score"                    data-balloon-length="medium" data-balloon="日常/高尔夫 共1话15分钟动画本篇+15分钟声优谈话→2021年12月31日18时30分" data-balloon-pos="up">
                    空色Utility</a>
                                        <a target="_blank" href="search.php?keyword=JOJO" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score"                    data-balloon-length="medium" data-balloon="第六季 石之海 →2021年12月1日Netflix（1-12话）" data-balloon-pos="up">
                    JOJO的奇妙冒险 石之海</a>
                                        <a target="_blank" href="search.php?keyword=%E6%B0%B8%E8%BF%9C%E7%9A%84831" style="margin:0 8px 0 0; padding: 5px 0 5px 0;" class="bgm_score"                    data-balloon-length="medium" data-balloon="寄了 神山你在干甚么 社会派 犯罪 青春 神山健治负责监督和脚本的原创动画→2022年1月30日起 每周日19:00" data-balloon-pos="up">
                    永远的831</a>
                    
                </td>
            </tr>
                        
            <!-- reset the cycle -->
                        
                        </tbody>
        </table>
    </div>
    
</div>


<div class="box clear rounded">
<h2 class="title">
        <div class="left">全部&nbsp;&nbsp;<a href="rss.xml" target="_blank"><img src="images/icon_rss2.png" alt="RSS" align="absmiddle" /></a></div>
    <div class="right text_normal"><a href="today-1.html">只看今天</a>&nbsp;&nbsp;|&nbsp;&nbsp;<a href="weekly-1.html">只看本周</a></div>
    </h2>

<div class="clear">
<table id="listTable" class="list_style table_fixed">
  <thead class="tcat">
      <tr>
        <th axis="string" class="l1">发表时间</th>
        <th axis="string" class="l2">类别</th>
        <th axis="string" class="l3">标题</th>
        <th axis="size" class="l4">大小</th>
        <th axis="number" class="l5">种子</th>
        <th axis="number" class="l6">下载</th>
        <th axis="number" class="l7">完成</th>
        <th axis="string" class="l8">UP主 / 代号</th>
      </tr>
  </thead>
  <tbody class="tbody" id="data_list">
              <tr class="alt1">
        <td nowrap="nowrap">
                            今天 23:09
                    </td>
        <td><a href="sort-1-1.html">动画</a></td>
        <td style="text-align:left;">
                                <a href="show-e83f7a6092c662faa16003769114abf0a5ef2148.html" target="_blank">
        [orion origin发布组] 更衣人偶坠入爱河/恋上换装娃娃 Sono Bisque Doll wa Koi wo Suru [07] [1440p] [GB] [网盘] [2022年1月番]</a>
                            </td>
        <td>874.4MB</td>
        <td nowrap="nowrap">
            <span class="bts_3">0</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_3">0</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_3">0</span>
        </td>
        <td><a href="team-139-1.html" class="text_green">orion个人发布</a></td>
      </tr>
              <tr class="alt2">
        <td nowrap="nowrap">
                            今天 23:04
                    </td>
        <td><a href="sort-1-1.html">动画</a></td>
        <td style="text-align:left;">
                                <a href="show-0c692371a9e9992c3c4e3d04d2555d39a34424b3.html" target="_blank">
        [orion origin发布组] 更衣人偶坠入爱河/恋上换装娃娃 Sono Bisque Doll wa Koi wo Suru [07] [1080p] [GB] [网盘] [2022年1月番]</a>
                            </td>
        <td>458.5MB</td>
        <td nowrap="nowrap">
            <span class="bts_3">0</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_3">0</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_3">0</span>
        </td>
        <td><a href="team-139-1.html" class="text_green">orion个人发布</a></td>
      </tr>
              <tr class="alt1">
        <td nowrap="nowrap">
                            今天 22:58
                    </td>
        <td><a href="sort-1-1.html">动画</a></td>
        <td style="text-align:left;">
                                <a href="show-b85f1de7a776504ac7063c0aab216507a1a2f5a0.html" target="_blank">
        非凡公主希瑞.She-Ra.and.the.Princesses.of.Power.S05E12.中英字幕.WEB.720P.甜饼字幕组.mp4</a>
                            </td>
        <td>281.7MB</td>
        <td nowrap="nowrap">
            <span class="bts_3">0</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_2">1</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_2">1</span>
        </td>
        <td><a href="team-29-1.html" class="text_green">甜饼字幕组</a></td>
      </tr>
              <tr class="alt2">
        <td nowrap="nowrap">
                            今天 22:56
                    </td>
        <td><a href="sort-1-1.html">动画</a></td>
        <td style="text-align:left;">
                                <a href="show-4e86ad12b659826d1248a7bd5bfb0dfe2787b280.html" target="_blank">
        [霜庭云花Sub&amp;森之屋动画组][平凡職業造就世界最強 第二季 / Arifureta Shokugyou de Sekai Saikyou S2][06][1080P][AVC][繁日内嵌]</a>
                            </td>
        <td>367.3MB</td>
        <td nowrap="nowrap">
            <span class="bts_3">0</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_3">0</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_3">0</span>
        </td>
        <td><a href="search.php?keyword=%E9%9C%9C%E5%BA%AD%E4%BA%91%E8%8A%B1Sub">霜庭云花Sub</a></td>
      </tr>
              <tr class="alt1">
        <td nowrap="nowrap">
                            今天 22:46
                    </td>
        <td><a href="sort-1-1.html">动画</a></td>
        <td style="text-align:left;">
                                <a href="show-96c18b8f26ff079793c6baca87c4d98379b24d5b.html" target="_blank">
        [霜庭云花Sub&amp;森之屋动画组]平凡职业成就世界最强 第二季 / Arifureta Shokugyou de Sekai Saikyou S2 [06][WebRip 1080P][简繁日内封]</a>
                            </td>
        <td>340.8MB</td>
        <td nowrap="nowrap">
            <span class="bts_1">6</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_1">20</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_1">10</span>
        </td>
        <td><a href="search.php?keyword=%E9%9C%9C%E5%BA%AD%E4%BA%91%E8%8A%B1Sub">霜庭云花Sub</a></td>
      </tr>
              <tr class="alt2">
        <td nowrap="nowrap">
                            今天 22:45
                    </td>
        <td><a href="sort-4-1.html">周边</a></td>
        <td style="text-align:left;">
                                <a href="show-746238e27a5b10e789ce1ad1d9d48eba6cb5dc8a.html" target="_blank">
        [MagicStar] 真凶标签 / 真犯人フラグ EP17 [WEBDL] [1080p] [HULU]【生】【附日字】</a>
                            </td>
        <td>1.3GB</td>
        <td nowrap="nowrap">
            <span class="bts_3">0</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_1">5</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_2">2</span>
        </td>
        <td><a href="search.php?keyword=%E9%AD%94%E6%98%9F%E5%AD%97%E5%B9%95%E5%9B%A2">魔星字幕团</a></td>
      </tr>
              <tr class="alt1">
        <td nowrap="nowrap">
                            今天 22:34
                    </td>
        <td><a href="sort-2-1.html">漫画</a></td>
        <td style="text-align:left;">
                                <a href="show-d8d1801b19c169258a5bdd6a037c3a38b7507607.html" target="_blank">
        [ARIA吧汉化组]Aria Special Navigation AQUARIA III(Aria The Benedizione)</a>
                            </td>
        <td>335.4MB</td>
        <td nowrap="nowrap">
            <span class="bts_3">0</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_1">11</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_2">1</span>
        </td>
        <td><a href="search.php?keyword=ARIA%E5%90%A7%E6%B1%89%E5%8C%96%E7%BB%84">ARIA吧汉化组</a></td>
      </tr>
              <tr class="alt2">
        <td nowrap="nowrap">
                            今天 22:31
                    </td>
        <td><a href="sort-1-1.html">动画</a></td>
        <td style="text-align:left;">
                                <a href="show-630c55583b4343cbc61332104748a8866fc40a64.html" target="_blank">
        [NC-Raws] FUTSAL BOYS!! - 07 (B-Global 1920x1080 HEVC AAC MKV)</a>
                            </td>
        <td>647.7MB</td>
        <td nowrap="nowrap">
            <span class="bts_2">3</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_1">12</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_1">4</span>
        </td>
        <td><a href="search.php?keyword=NC-Raws">NC-Raws</a></td>
      </tr>
              <tr class="alt1">
        <td nowrap="nowrap">
                            今天 22:30
                    </td>
        <td><a href="sort-1-1.html">动画</a></td>
        <td style="text-align:left;">
                                <a href="show-44a3a78e0d2497e6d0b2e4ad311f6426875efe86.html" target="_blank">
        [NC-Raws] Futsal Boys！！！！！ - 07 (Baha 1920x1080 AVC AAC MP4)</a>
                            </td>
        <td>380.8MB</td>
        <td nowrap="nowrap">
            <span class="bts_1">17</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_1">35</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_1">22</span>
        </td>
        <td><a href="search.php?keyword=NC-Raws">NC-Raws</a></td>
      </tr>
              <tr class="alt2">
        <td nowrap="nowrap">
                            今天 22:30
                    </td>
        <td><a href="sort-1-1.html">动画</a></td>
        <td style="text-align:left;">
                                <a href="show-a78ffdc7a571e8ccaeecf893f098f34cd1600c22.html" target="_blank">
        [Lilith-Raws] 薔薇王的葬列 / Baraou no Souretsu - 07 [Baha][WEB-DL][1080p][AVC AAC][CHT][MP4]</a>
                            </td>
        <td>401.4MB</td>
        <td nowrap="nowrap">
            <span class="bts_3">0</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_3">0</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_3">0</span>
        </td>
        <td><a href="search.php?keyword=Lilith-Raws">Lilith-Raws</a></td>
      </tr>
              <tr class="alt1">
        <td nowrap="nowrap">
                            今天 22:25
                    </td>
        <td><a href="sort-5-1.html">其它</a></td>
        <td style="text-align:left;">
                                <a href="show-c5f0f2c3456183eb52b933cf93e0cd460025cc14.html" target="_blank">
        自由之缰.Free.Rein.S01E09.中英字幕.WEB.720P.甜饼字幕组.mp4</a>
                            </td>
        <td>293.8MB</td>
        <td nowrap="nowrap">
            <span class="bts_3">0</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_3">0</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_3">0</span>
        </td>
        <td><a href="team-29-1.html" class="text_green">甜饼字幕组</a></td>
      </tr>
              <tr class="alt2">
        <td nowrap="nowrap">
                            今天 22:25
                    </td>
        <td><a href="sort-5-1.html">其它</a></td>
        <td style="text-align:left;">
                                <a href="show-5d2a9f09737dcf0e8a91c19fd4ffcb2af7878e22.html" target="_blank">
        自由之缰.Free.Rein.S01E08.中英字幕.WEB.720P.甜饼字幕组.mp4</a>
                            </td>
        <td>304.2MB</td>
        <td nowrap="nowrap">
            <span class="bts_3">0</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_3">0</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_3">0</span>
        </td>
        <td><a href="team-29-1.html" class="text_green">甜饼字幕组</a></td>
      </tr>
              <tr class="alt1">
        <td nowrap="nowrap">
                            今天 22:21
                    </td>
        <td><a href="sort-1-1.html">动画</a></td>
        <td style="text-align:left;">
                                <a href="show-c7135a681323e2a1e037431e1165c9f3422c762b.html" target="_blank">
        [c.c動漫][1月新番][Ryman's Club / 白領羽球部][04][簡繁內掛][AVC_AAC][1080P][MKV]</a>
                            </td>
        <td>146.6MB</td>
        <td nowrap="nowrap">
            <span class="bts_1">16</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_1">40</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_1">22</span>
        </td>
        <td><a href="search.php?keyword=c.c%E5%8A%A8%E6%BC%AB">c.c动漫</a></td>
      </tr>
              <tr class="alt2">
        <td nowrap="nowrap">
                            今天 22:18
                    </td>
        <td><a href="sort-1-1.html">动画</a></td>
        <td style="text-align:left;">
                                <a href="show-9cf2df4024bd17a09386c0a0010471f2bdd59f4f.html" target="_blank">
        ​[c.c動漫][4月番][慕留人 -火影忍者新世代-][Boruto -Naruto Next Generations-][237][BIG5][720P][MP4]</a>
                            </td>
        <td>161.5MB</td>
        <td nowrap="nowrap">
            <span class="bts_1">51</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_1">103</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_1">72</span>
        </td>
        <td><a href="search.php?keyword=c.c%E5%8A%A8%E6%BC%AB">c.c动漫</a></td>
      </tr>
              <tr class="alt1">
        <td nowrap="nowrap">
                            今天 22:13
                    </td>
        <td><a href="sort-4-1.html">周边</a></td>
        <td style="text-align:left;">
                                <a href="show-8b71338f0727200da91ba75b2b0b68b962ff0d35.html" target="_blank">
        [MagicStar] DCU ~持有手铐的潜水员~ / DCU ~手錠を持ったダイバー~ EP05 [WEBDL] [1080p]【生】【附日字】</a>
                            </td>
        <td>1.2GB</td>
        <td nowrap="nowrap">
            <span class="bts_3">0</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_3">0</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_3">0</span>
        </td>
        <td><a href="search.php?keyword=%E9%AD%94%E6%98%9F%E5%AD%97%E5%B9%95%E5%9B%A2">魔星字幕团</a></td>
      </tr>
              <tr class="alt2">
        <td nowrap="nowrap">
                            今天 22:11
                    </td>
        <td><a href="sort-4-1.html">周边</a></td>
        <td style="text-align:left;">
                                <a href="show-7b824ad2a7ef3a23906fdd2b2faf6bd3f55de96c.html" target="_blank">
        [MagicStar] 逃亡医生F / 逃亡医F EP06 [WEBDL] [1080p] [HULU]【生】【附日字】</a>
                            </td>
        <td>1.3GB</td>
        <td nowrap="nowrap">
            <span class="bts_1">4</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_1">12</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_1">7</span>
        </td>
        <td><a href="search.php?keyword=%E9%AD%94%E6%98%9F%E5%AD%97%E5%B9%95%E5%9B%A2">魔星字幕团</a></td>
      </tr>
              <tr class="alt1">
        <td nowrap="nowrap">
                            今天 22:01
                    </td>
        <td><a href="sort-1-1.html">动画</a></td>
        <td style="text-align:left;">
                                <a href="show-70502f5f4138b7b2a92d0b18413802c158a6c414.html" target="_blank">
        [NC-Raws] 锈色铠甲 黎明 / Sabiiro no Armor: Reimei - 07 (B-Global 1920x1080 HEVC AAC MKV)</a>
                            </td>
        <td>651.1MB</td>
        <td nowrap="nowrap">
            <span class="bts_1">10</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_1">17</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_1">9</span>
        </td>
        <td><a href="search.php?keyword=NC-Raws">NC-Raws</a></td>
      </tr>
              <tr class="alt2">
        <td nowrap="nowrap">
                            今天 22:01
                    </td>
        <td><a href="sort-1-1.html">动画</a></td>
        <td style="text-align:left;">
                                <a href="show-6cb876fa76ab853b630856ba34402765cb0ecbb1.html" target="_blank">
        [NC-Raws] 薔薇王的葬列 / Baraou no Souretsu - 07 (B-Global 3840x2160 HEVC AAC MKV)</a>
                            </td>
        <td>566.2MB</td>
        <td nowrap="nowrap">
            <span class="bts_3">0</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_3">0</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_3">0</span>
        </td>
        <td><a href="search.php?keyword=NC-Raws">NC-Raws</a></td>
      </tr>
              <tr class="alt1">
        <td nowrap="nowrap">
                            今天 22:01
                    </td>
        <td><a href="sort-1-1.html">动画</a></td>
        <td style="text-align:left;">
                                <a href="show-f0ba2fe747e6edbafc15e07ae969228a0ce4244e.html" target="_blank">
        [NC-Raws] 薔薇王的葬列 / Baraou no Souretsu - 07 (Baha 1920x1080 AVC AAC MP4)</a>
                            </td>
        <td>401.5MB</td>
        <td nowrap="nowrap">
            <span class="bts_3">0</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_3">0</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_3">0</span>
        </td>
        <td><a href="search.php?keyword=NC-Raws">NC-Raws</a></td>
      </tr>
              <tr class="alt2">
        <td nowrap="nowrap">
                            今天 21:40
                    </td>
        <td><a href="sort-4-1.html">周边</a></td>
        <td style="text-align:left;">
                                <a href="show-e5a83a99e00f3fd03565dabcc89df75fab481397.html" target="_blank">
        【幻月字幕组】【22年日剧】【如果 有一所只有帅哥的高中】【06】【1080P】【中日双语】</a>
                            </td>
        <td>541.9MB</td>
        <td nowrap="nowrap">
            <span class="bts_1">12</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_1">12</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_1">13</span>
        </td>
        <td><a href="search.php?keyword=%E5%B9%BB%E6%9C%88%E5%AD%97%E5%B9%95%E7%BB%84">幻月字幕组</a></td>
      </tr>
              <tr class="alt1">
        <td nowrap="nowrap">
                            今天 21:00
                    </td>
        <td><a href="sort-1-1.html">动画</a></td>
        <td style="text-align:left;">
                                <a href="show-79166482e2c94b941dda49ef2e84abc5139471fe.html" target="_blank">
        [OPFans楓雪動漫][ONE PIECE 海賊王][第1011話][1080p][周日版][MKV]</a>
                            </td>
        <td>1.4GB</td>
        <td nowrap="nowrap">
            <span class="bts_1">74</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_1">138</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_1">90</span>
        </td>
        <td><a href="search.php?keyword=%E9%81%8E%E5%8E%BB%E3%83%8E%E8%A8%98%E6%86%B6">過去ノ記憶</a></td>
      </tr>
              <tr class="alt2">
        <td nowrap="nowrap">
                            今天 20:54
                    </td>
        <td><a href="sort-1-1.html">动画</a></td>
        <td style="text-align:left;">
                                <a href="show-a28bd20be4cc930341698d8d63f035baecf4e0a5.html" target="_blank">
        [orion origin发布组] 明日酱的水手服/明日同学的水手服 Akebi-chan no Sailor-fuku [07] [1440p] [GB] [mkv] [网盘] [2022年1月番]</a>
                            </td>
        <td>1.0GB</td>
        <td nowrap="nowrap">
            <span class="bts_1">10</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_1">31</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_1">16</span>
        </td>
        <td><a href="team-139-1.html" class="text_green">orion个人发布</a></td>
      </tr>
              <tr class="alt1">
        <td nowrap="nowrap">
                            今天 20:48
                    </td>
        <td><a href="sort-1-1.html">动画</a></td>
        <td style="text-align:left;">
                                <a href="show-59e4041f2a47cdc611cc0ff541d3246e0c54823d.html" target="_blank">
        [orion origin发布组] 明日酱的水手服/明日同学的水手服 Akebi-chan no Sailor-fuku [07] [1080p] [GB] [mkv] [网盘] [2022年1月番]</a>
                            </td>
        <td>543.3MB</td>
        <td nowrap="nowrap">
            <span class="bts_1">7</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_1">33</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_1">19</span>
        </td>
        <td><a href="team-139-1.html" class="text_green">orion个人发布</a></td>
      </tr>
              <tr class="alt2">
        <td nowrap="nowrap">
                            今天 20:25
                    </td>
        <td><a href="sort-1-1.html">动画</a></td>
        <td style="text-align:left;">
                                <a href="show-4c6d1f5ad211685cface420a178b9969dd7a62d8.html" target="_blank">
        [Skymoon-Raws][One Piece 海賊王][1011][ViuTV][WEB-RIP][CHT][SRTx2][1080p][MKV](先行版本) V2</a>
                            </td>
        <td>451.5MB</td>
        <td nowrap="nowrap">
            <span class="bts_3">0</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_3">0</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_3">0</span>
        </td>
        <td><a href="search.php?keyword=%E5%A4%A9%E6%9C%88%E5%8B%95%E6%BC%AB%26amp%3B%E7%99%BC%E4%BD%88%E7%B5%84">天月動漫&amp;發佈組</a></td>
      </tr>
              <tr class="alt1">
        <td nowrap="nowrap">
                            今天 20:25
                    </td>
        <td><a href="sort-1-1.html">动画</a></td>
        <td style="text-align:left;">
                                <a href="show-6d0148813b079b3bc54d65c1782bee9de4720242.html" target="_blank">
        [OPFans枫雪动漫][ONE PIECE 海贼王][第1011话][720p][周日版][MP4]</a>
                            </td>
        <td>225.2MB</td>
        <td nowrap="nowrap">
            <span class="bts_1">216</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_1">162</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_1">170</span>
        </td>
        <td><a href="search.php?keyword=%E9%81%8E%E5%8E%BB%E3%83%8E%E8%A8%98%E6%86%B6">過去ノ記憶</a></td>
      </tr>
              <tr class="alt2">
        <td nowrap="nowrap">
                            今天 20:23
                    </td>
        <td><a href="sort-1-1.html">动画</a></td>
        <td style="text-align:left;">
                                <a href="show-6e7a4a22779aba647523da36bd1cdf6eb75a668e.html" target="_blank">
        [Skymoon-Raws][One Piece 海賊王][1011][ViuTV][WEB-DL][1080p][AVC AAC][CHT][MP4+ASS](正式版本) V2</a>
                            </td>
        <td>466.9MB</td>
        <td nowrap="nowrap">
            <span class="bts_3">0</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_3">0</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_3">0</span>
        </td>
        <td><a href="search.php?keyword=%E5%A4%A9%E6%9C%88%E5%8B%95%E6%BC%AB%26amp%3B%E7%99%BC%E4%BD%88%E7%B5%84">天月動漫&amp;發佈組</a></td>
      </tr>
              <tr class="alt1">
        <td nowrap="nowrap">
                            今天 20:01
                    </td>
        <td><a href="sort-1-1.html">动画</a></td>
        <td style="text-align:left;">
                                <a href="show-1b01890c4acc775f907162062c1bf49b0a8c342d.html" target="_blank">
        [NC-Raws] 龙蛇演义 / Dragon's Disciple - 06 (B-Global Donghua 1920x1080 HEVC AAC MKV)</a>
                            </td>
        <td>548.8MB</td>
        <td nowrap="nowrap">
            <span class="bts_1">39</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_1">24</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_1">33</span>
        </td>
        <td><a href="search.php?keyword=NC-Raws">NC-Raws</a></td>
      </tr>
              <tr class="alt2">
        <td nowrap="nowrap">
                            今天 19:47
                    </td>
        <td><a href="sort-1-1.html">动画</a></td>
        <td style="text-align:left;">
                                <a href="show-f660492b2d1970aafd1783c0f41fb8347fa16f57.html" target="_blank">
        [神滨市立大学勇者部&amp;LoliHouse] 魔法纪录 魔法少女小圆外传 第二季 -觉醒前夜- / Magia Record S2 [06-08][BDRip 1080p HEVC-10bit FLAC][简繁内封]</a>
                            </td>
        <td>4.1GB</td>
        <td nowrap="nowrap">
            <span class="bts_3">0</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_1">150</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_1">6</span>
        </td>
        <td><a href="search.php?keyword=LoliHouse">LoliHouse</a></td>
      </tr>
              <tr class="alt1">
        <td nowrap="nowrap">
                            今天 19:32
                    </td>
        <td><a href="sort-1-1.html">动画</a></td>
        <td style="text-align:left;">
                                <a href="show-691f4874590565226e5bb37f7bac6bbd08a254e9.html" target="_blank">
        [离谱Sub] 杀爱 / 殺し愛 / Koroshi Ai [06][HEVC AAC][1080p][简繁内封]</a>
                            </td>
        <td>244.3MB</td>
        <td nowrap="nowrap">
            <span class="bts_3">0</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_3">0</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_3">0</span>
        </td>
        <td><a href="team-138-1.html" class="text_green">离谱Sub</a></td>
      </tr>
              <tr class="alt2">
        <td nowrap="nowrap">
                            今天 19:32
                    </td>
        <td><a href="sort-1-1.html">动画</a></td>
        <td style="text-align:left;">
                                <a href="show-1dc4f1ffe83027e348339aa5e78ac8f81bc21dee.html" target="_blank">
        [离谱Sub] 相愛相殺 / 殺し愛 / Koroshi Ai [03][AVC AAC][1080p][繁體內嵌]</a>
                            </td>
        <td>431.9MB</td>
        <td nowrap="nowrap">
            <span class="bts_3">0</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_3">0</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_3">0</span>
        </td>
        <td><a href="team-138-1.html" class="text_green">离谱Sub</a></td>
      </tr>
              <tr class="alt1">
        <td nowrap="nowrap">
                            今天 19:32
                    </td>
        <td><a href="sort-1-1.html">动画</a></td>
        <td style="text-align:left;">
                                <a href="show-b5889e97bd6dfd8015f04c3356badcb8c9860c72.html" target="_blank">
        [离谱Sub] 杀爱 / 殺し愛 / Koroshi Ai [06][AVC AAC][1080p][简体内嵌]</a>
                            </td>
        <td>431.4MB</td>
        <td nowrap="nowrap">
            <span class="bts_1">10</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_1">13</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_1">10</span>
        </td>
        <td><a href="team-138-1.html" class="text_green">离谱Sub</a></td>
      </tr>
              <tr class="alt2">
        <td nowrap="nowrap">
                            今天 19:21
                    </td>
        <td><a href="sort-1-1.html">动画</a></td>
        <td style="text-align:left;">
                                <a href="show-5de921441b85392adebc2a605ab4fa7f04803e6c.html" target="_blank">
        [桜都字幕组] 终末的后宫 / Shuumatsu no Harem [07][1080P][简体内嵌]</a>
                            </td>
        <td>480.4MB</td>
        <td nowrap="nowrap">
            <span class="bts_1">243</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_1">96</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_1">156</span>
        </td>
        <td><a href="search.php?keyword=%E6%A1%9C%E9%83%BD%E5%AD%97%E5%B9%95%E7%BB%84">桜都字幕组</a></td>
      </tr>
              <tr class="alt1">
        <td nowrap="nowrap">
                            今天 19:20
                    </td>
        <td><a href="sort-1-1.html">动画</a></td>
        <td style="text-align:left;">
                                <a href="show-9fae0c993cea0ff60f8773deb8d897a5109daecf.html" target="_blank">
        [桜都字幕组] 终末的后宫 / Shuumatsu no Harem [07][1080P][简体内嵌]</a>
                            </td>
        <td>480.6MB</td>
        <td nowrap="nowrap">
            <span class="bts_1">106</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_1">83</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_1">84</span>
        </td>
        <td><a href="search.php?keyword=%E6%A1%9C%E9%83%BD%E5%AD%97%E5%B9%95%E7%BB%84">桜都字幕组</a></td>
      </tr>
              <tr class="alt2">
        <td nowrap="nowrap">
                            今天 19:20
                    </td>
        <td><a href="sort-1-1.html">动画</a></td>
        <td style="text-align:left;">
                                <a href="show-92ef70346cbde06655e5b849e18a8a2d884172ab.html" target="_blank">
        [桜都字幕组] 终末的后宫 / Shuumatsu no Harem [07][1080P][简繁内封]</a>
                            </td>
        <td>470.5MB</td>
        <td nowrap="nowrap">
            <span class="bts_1">154</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_1">34</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_1">117</span>
        </td>
        <td><a href="search.php?keyword=%E6%A1%9C%E9%83%BD%E5%AD%97%E5%B9%95%E7%BB%84">桜都字幕组</a></td>
      </tr>
              <tr class="alt1">
        <td nowrap="nowrap">
                            今天 19:17
                    </td>
        <td><a href="sort-1-1.html">动画</a></td>
        <td style="text-align:left;">
                                <a href="show-b9a2ba27123fd0c8a394cfb8774089c8920bc9b0.html" target="_blank">
        [神滨市立大学勇者部&amp;LoliHouse] 魔法纪录 魔法少女小圆外传 第二季 -觉醒前夜- / Magia Record S2 [01-03][BDRip 1080p HEVC-10bit FLAC][简繁内封] v2</a>
                            </td>
        <td>5.1GB</td>
        <td nowrap="nowrap">
            <span class="bts_3">0</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_3">0</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_3">0</span>
        </td>
        <td><a href="search.php?keyword=LoliHouse">LoliHouse</a></td>
      </tr>
              <tr class="alt2">
        <td nowrap="nowrap">
                            今天 18:46
                    </td>
        <td><a href="sort-1-1.html">动画</a></td>
        <td style="text-align:left;">
                                <a href="show-68322434dea2b9c09c05bedcfda1992403356e8e.html" target="_blank">
        [雪飘工作室][ワッチャプリマジ！/Whatcha _Pri-maji!/哇恰美妙魔法！][19][简](检索:/美妙旋律/美妙天堂/美妙频道) 附外挂字幕</a>
                            </td>
        <td>258.7MB</td>
        <td nowrap="nowrap">
            <span class="bts_1">60</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_1">14</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_1">32</span>
        </td>
        <td><a href="search.php?keyword=%E9%9B%AA%E9%A3%98%E5%B7%A5%E4%BD%9C%E5%AE%A4">雪飘工作室</a></td>
      </tr>
              <tr class="alt1">
        <td nowrap="nowrap">
                            今天 18:08
                    </td>
        <td><a href="sort-4-1.html">周边</a></td>
        <td style="text-align:left;">
                                <a href="show-e2c1d538a83ff1b0348f0a4df5624ec34fab52b5.html" target="_blank">
        【幻月字幕组】【22年日剧】【鹿枫堂四色日和】【06】【1080P】【中日双语】</a>
                            </td>
        <td>542.9MB</td>
        <td nowrap="nowrap">
            <span class="bts_1">50</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_1">15</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_1">26</span>
        </td>
        <td><a href="search.php?keyword=%E5%B9%BB%E6%9C%88%E5%AD%97%E5%B9%95%E7%BB%84">幻月字幕组</a></td>
      </tr>
              <tr class="alt2">
        <td nowrap="nowrap">
                            今天 18:04
                    </td>
        <td><a href="sort-1-1.html">动画</a></td>
        <td style="text-align:left;">
                                <a href="show-7babb62f3af830c6bac61ff0d465bbfe2fea0088.html" target="_blank">
        【飞沐team】大头儿子和小头爸爸 156集全(1995)[国语无字][DVD修复版](1080P)[MKV]</a>
                            </td>
        <td>202.7GB</td>
        <td nowrap="nowrap">
            <span class="bts_3">0</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_3">0</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_3">0</span>
        </td>
        <td><a href="team-117-1.html" class="text_green">飞沐team</a></td>
      </tr>
              <tr class="alt1">
        <td nowrap="nowrap">
                            今天 18:03
                    </td>
        <td><a href="sort-1-1.html">动画</a></td>
        <td style="text-align:left;">
                                <a href="show-51f6ca83bd8af986618e22348f781909f2c94f5b.html" target="_blank">
        [星空字幕組][白領羽球部 / Ryman's Club][02][繁日雙語][1080P][WEBrip][MP4]（急招翻譯、校對）</a>
                            </td>
        <td>354.4MB</td>
        <td nowrap="nowrap">
            <span class="bts_3">0</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_3">0</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_3">0</span>
        </td>
        <td><a href="search.php?keyword=%E6%98%9F%E7%A9%BA%E5%AD%97%E5%B9%95%E7%BB%84">星空字幕组</a></td>
      </tr>
              <tr class="alt2">
        <td nowrap="nowrap">
                            今天 18:02
                    </td>
        <td><a href="sort-1-1.html">动画</a></td>
        <td style="text-align:left;">
                                <a href="show-8756a6497f5961a27a28aebec83a8ce952fa78d5.html" target="_blank">
        [星空字幕组][白领羽球部 / Ryman's Club][02][简日双语][1080P][WEBrip][MP4]（急招翻译、校对）</a>
                            </td>
        <td>354.6MB</td>
        <td nowrap="nowrap">
            <span class="bts_1">16</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_1">13</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_1">14</span>
        </td>
        <td><a href="search.php?keyword=%E6%98%9F%E7%A9%BA%E5%AD%97%E5%B9%95%E7%BB%84">星空字幕组</a></td>
      </tr>
              <tr class="alt1">
        <td nowrap="nowrap">
                            今天 17:31
                    </td>
        <td><a href="sort-1-1.html">动画</a></td>
        <td style="text-align:left;">
                                <a href="show-c4f911fbaf872be72f7bec7eb38132b8ccaba6d2.html" target="_blank">
        [NC-Raws] 博人传之火影忍者次世代 / Boruto - Naruto Next Generations - 237 (B-Global 1920x1080 HEVC AAC MKV)</a>
                            </td>
        <td>653.3MB</td>
        <td nowrap="nowrap">
            <span class="bts_1">52</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_1">17</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_1">29</span>
        </td>
        <td><a href="search.php?keyword=NC-Raws">NC-Raws</a></td>
      </tr>
              <tr class="alt2">
        <td nowrap="nowrap">
                            今天 17:30
                    </td>
        <td><a href="sort-1-1.html">动画</a></td>
        <td style="text-align:left;">
                                <a href="show-b36b5ae62040a621049e3fb762b09fcfae25ce61.html" target="_blank">
        [神滨市立大学勇者部&amp;LoliHouse] 魔法纪录 魔法少女小圆外传 第二季 -觉醒前夜- / Magia Record S2 [04-05][BDRip 1080p HEVC-10bit FLAC][简繁内封]</a>
                            </td>
        <td>2.9GB</td>
        <td nowrap="nowrap">
            <span class="bts_1">45</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_1">62</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_1">69</span>
        </td>
        <td><a href="search.php?keyword=LoliHouse">LoliHouse</a></td>
      </tr>
              <tr class="alt1">
        <td nowrap="nowrap">
                            今天 17:24
                    </td>
        <td><a href="sort-3-1.html">音乐</a></td>
        <td style="text-align:left;">
                                <a href="show-a4e944be5f22c97e983990c3b73215d9f283e139.html" target="_blank">
        [220216][少女与战车]『ガールズ&amp;パンツァー 劇場版』ボーカルミニアルバム「音楽道、まだまだ邁進中です!!!」ChouCho×佐咲紗花、渕上舞、田中理恵、喜多村英梨、川澄綾子、吉岡麻耶、金元寿子、瀬戸麻沙美、能登麻美子、竹達彩奈[FLAC]</a>
                            </td>
        <td>279.7MB</td>
        <td nowrap="nowrap">
            <span class="bts_3">0</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_3">0</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_3">0</span>
        </td>
        <td><a href="search.php?keyword=%E5%A4%A9%E4%BD%BF%E5%8A%A8%E6%BC%AB%E8%AE%BA%E5%9D%9B">天使动漫论坛</a></td>
      </tr>
              <tr class="alt2">
        <td nowrap="nowrap">
                            今天 17:11
                    </td>
        <td><a href="sort-3-1.html">音乐</a></td>
        <td style="text-align:left;">
                                <a href="show-87d2fccb09c49ab6abf1c9926c23d452aaa89bc3.html" target="_blank">
        [220216][少女与战车]『ガールズ&amp;パンツァー 劇場版』ボーカルミニアルバム「音楽道、まだまだ邁進中です!!!」ChouCho×佐咲紗花、渕上舞、田中理恵、喜多村英梨、川澄綾子、吉岡麻耶、金元寿子、瀬戸麻沙美、能登麻美子、竹達彩奈[320K]</a>
                            </td>
        <td>79.9MB</td>
        <td nowrap="nowrap">
            <span class="bts_1">297</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_1">26</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_1">107</span>
        </td>
        <td><a href="search.php?keyword=%E5%A4%A9%E4%BD%BF%E5%8A%A8%E6%BC%AB%E8%AE%BA%E5%9D%9B">天使动漫论坛</a></td>
      </tr>
              <tr class="alt1">
        <td nowrap="nowrap">
                            今天 17:11
                    </td>
        <td><a href="sort-1-1.html">动画</a></td>
        <td style="text-align:left;">
                                <a href="show-cca8da9e9d7f2b7f0f8696797d54de884677b069.html" target="_blank">
        [DBD-Raws][枪神斯托拉塔斯/Gunslinger Stratos The Animation/ガンスリンガー ストラトス THE ANIMATION][01-12TV全集+SP+特典映像][1080P][BDRip][HEVC-10bit][FLAC][MKV]</a>
                            </td>
        <td>28.2GB</td>
        <td nowrap="nowrap">
            <span class="bts_1">10</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_1">101</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_1">14</span>
        </td>
        <td><a href="search.php?keyword=DBD%E5%88%B6%E4%BD%9C%E7%BB%84">DBD制作组</a></td>
      </tr>
              <tr class="alt2">
        <td nowrap="nowrap">
                            今天 16:42
                    </td>
        <td><a href="sort-1-1.html">动画</a></td>
        <td style="text-align:left;">
                                <a href="show-828c0da93a199526e35a0a8d518d805221f683e1.html" target="_blank">
        [Skymoon-Raws][One Piece 海賊王][1011][ViuTV][WEB-DL][1080p][AVC AAC][CHT][MP4+ASS](正式版本)</a>
                            </td>
        <td>466.9MB</td>
        <td nowrap="nowrap">
            <span class="bts_1">169</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_1">16</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_1">77</span>
        </td>
        <td><a href="search.php?keyword=%E5%A4%A9%E6%9C%88%E5%8B%95%E6%BC%AB%26amp%3B%E7%99%BC%E4%BD%88%E7%B5%84">天月動漫&amp;發佈組</a></td>
      </tr>
              <tr class="alt1">
        <td nowrap="nowrap">
                            今天 16:42
                    </td>
        <td><a href="sort-1-1.html">动画</a></td>
        <td style="text-align:left;">
                                <a href="show-b1e762663a3257189775aeb2001331ea31e894df.html" target="_blank">
        [Skymoon-Raws][One Piece 海賊王][1011][ViuTV][WEB-RIP][CHT][SRT][1080p][MKV](先行版本)</a>
                            </td>
        <td>451.5MB</td>
        <td nowrap="nowrap">
            <span class="bts_3">0</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_3">0</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_3">0</span>
        </td>
        <td><a href="search.php?keyword=%E5%A4%A9%E6%9C%88%E5%8B%95%E6%BC%AB%26amp%3B%E7%99%BC%E4%BD%88%E7%B5%84">天月動漫&amp;發佈組</a></td>
      </tr>
              <tr class="alt2">
        <td nowrap="nowrap">
                            今天 15:46
                    </td>
        <td><a href="sort-1-1.html">动画</a></td>
        <td style="text-align:left;">
                                <a href="show-bc4c90d11999aaa50e2674a308685dc8aa1b97aa.html" target="_blank">
        [桜都字幕组] 自称贤者弟子的贤者 / Kenja no Deshi o Nanoru Kenja [06][1080p][简体内嵌]</a>
                            </td>
        <td>517.8MB</td>
        <td nowrap="nowrap">
            <span class="bts_1">93</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_1">27</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_1">46</span>
        </td>
        <td><a href="search.php?keyword=%E6%A1%9C%E9%83%BD%E5%AD%97%E5%B9%95%E7%BB%84">桜都字幕组</a></td>
      </tr>
              <tr class="alt1">
        <td nowrap="nowrap">
                            今天 15:45
                    </td>
        <td><a href="sort-1-1.html">动画</a></td>
        <td style="text-align:left;">
                                <a href="show-98d9dec39cb71fa949d26e139f159e3fb92a1489.html" target="_blank">
        [桜都字幕組] 自稱賢者弟子的賢者 / Kenja no Deshi o Nanoru Kenja [06][1080p][繁體內嵌]</a>
                            </td>
        <td>517.8MB</td>
        <td nowrap="nowrap">
            <span class="bts_1">188</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_1">13</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_1">67</span>
        </td>
        <td><a href="search.php?keyword=%E6%A1%9C%E9%83%BD%E5%AD%97%E5%B9%95%E7%BB%84">桜都字幕组</a></td>
      </tr>
              <tr class="alt2">
        <td nowrap="nowrap">
                            今天 15:45
                    </td>
        <td><a href="sort-1-1.html">动画</a></td>
        <td style="text-align:left;">
                                <a href="show-abcd915c85f5dc4af6a60124adebf6b2eaceb1eb.html" target="_blank">
        [桜都字幕组] 自称贤者弟子的贤者 / Kenja no Deshi o Nanoru Kenja [06][1080p][简繁内封]</a>
                            </td>
        <td>536.6MB</td>
        <td nowrap="nowrap">
            <span class="bts_3">0</span>
        </td>
        <td nowrap="nowrap">
            <span class="btl_3">0</span>
        </td>
        <td nowrap="nowrap">
            <span class="btc_3">0</span>
        </td>
        <td><a href="search.php?keyword=%E6%A1%9C%E9%83%BD%E5%AD%97%E5%B9%95%E7%BB%84">桜都字幕组</a></td>
      </tr>
    
  </tbody>
</table>
</div>

</div>
<div class="pages clear"><span class="nextprev">〈</span><span class="current">1</span><a href="/2.html" target="_self">2</a><a href="/3.html" target="_self">3</a><a href="/4.html" target="_self">4</a><a href="/5.html" target="_self">5</a><span>&#8230;.</span><a href="/99.html" class="pager-last active" target="_self">99</a><a href="/2.html" class="nextprev" target="_self">〉</a></div>
    <div class="clear text_center" style="margin-bottom:10px;"></div>

<script type="text/javascript" src="javascripts/tablesort.js?v20081028"></script>
<script type="text/javascript">
window.addEvent('domready', function(){new SortableTable($('listTable'))});
</script>

<div class="box clear rounded">
    <h2 class="title"><span class="left">联萌成员</span><a href="javascript:void(0);" onclick="panelDeploy('team', 1);return false;" class="right"><img src="images/collapse.gif" id="panel_team_button" alt="icon" /></a></h2>
    <table class="list_style table_fixed">
        <tbody id="panel_team">
                <tr class="alt1">            <td width="12%"><a href="team-1-1.html">爱恋字幕社</a></td>
                                        <td width="12%"><a href="team-2-1.html">漫猫字幕组</a></td>
                                        <td width="12%"><a href="team-3-1.html">风之圣殿字幕组</a></td>
                                        <td width="12%"><a href="team-4-1.html">梦星字幕组</a></td>
                                        <td width="12%"><a href="team-5-1.html">盲点字幕组</a></td>
                                        <td width="12%"><a href="team-6-1.html">HKACG+</a></td>
                                        <td width="12%"><a href="team-7-1.html">汐空字幕组</a></td>
                                        <td width="12%"><a href="team-8-1.html">极影字幕社</a></td>
            </tr>                <tr class="alt2">            <td width="12%"><a href="team-9-1.html">動漫國字幕組</a></td>
                                        <td width="12%"><a href="team-10-1.html">清蓝字幕组</a></td>
                                        <td width="12%"><a href="team-11-1.html">炸天团队</a></td>
                                        <td width="12%"><a href="team-12-1.html">雪梦字幕组</a></td>
                                        <td width="12%"><a href="team-13-1.html">吐槽字幕组</a></td>
                                        <td width="12%"><a href="team-14-1.html">银色子弹字幕组</a></td>
                                        <td width="12%"><a href="team-15-1.html">轻之国度字幕组</a></td>
                                        <td width="12%"><a href="team-16-1.html">风车字幕组</a></td>
            </tr>                <tr class="alt1">            <td width="12%"><a href="team-17-1.html">SSK字幕组</a></td>
                                        <td width="12%"><a href="team-18-1.html">风翼字幕组</a></td>
                                        <td width="12%"><a href="team-19-1.html">傲娇零字幕组</a></td>
                                        <td width="12%"><a href="team-20-1.html">佳芸字幕组</a></td>
                                        <td width="12%"><a href="team-21-1.html">星辰字幕组</a></td>
                                        <td width="12%"><a href="team-22-1.html">繁星字幕组</a></td>
                                        <td width="12%"><a href="team-23-1.html">矢车菊影音工作室</a></td>
                                        <td width="12%"><a href="team-24-1.html">宅死在家下辈子还宅</a></td>
            </tr>                <tr class="alt2">            <td width="12%"><a href="team-25-1.html">樱翼汉化组</a></td>
                                        <td width="12%"><a href="team-26-1.html">追放字幕组</a></td>
                                        <td width="12%"><a href="team-27-1.html">ZERO字幕组</a></td>
                                        <td width="12%"><a href="team-28-1.html">脸肿字幕组</a></td>
                                        <td width="12%"><a href="team-29-1.html">甜饼字幕组</a></td>
                                        <td width="12%"><a href="team-30-1.html">TSDM字幕組</a></td>
                                        <td width="12%"><a href="team-31-1.html">星火字幕组</a></td>
                                        <td width="12%"><a href="team-32-1.html">光荣字幕组</a></td>
            </tr>                <tr class="alt1">            <td width="12%"><a href="team-33-1.html">梦幻恋樱字幕组</a></td>
                                        <td width="12%"><a href="team-34-1.html">安良城红吧字幕组</a></td>
                                        <td width="12%"><a href="team-35-1.html">驯兽师联盟</a></td>
                                        <td width="12%"><a href="team-36-1.html">个人发布</a></td>
                                        <td width="12%"><a href="team-37-1.html">花語字幕組</a></td>
                                        <td width="12%"><a href="team-38-1.html">东京第七区字幕组</a></td>
                                        <td width="12%"><a href="team-39-1.html">CE家族社字幕组</a></td>
                                        <td width="12%"><a href="team-40-1.html">Ｆ宅字幕组</a></td>
            </tr>                <tr class="alt2">            <td width="12%"><a href="team-42-1.html">咪路fans制作组</a></td>
                                        <td width="12%"><a href="team-43-1.html">仲夏动漫字幕组</a></td>
                                        <td width="12%"><a href="team-44-1.html">动漫花园字幕组</a></td>
                                        <td width="12%"><a href="team-45-1.html">WOLF字幕组</a></td>
                                        <td width="12%"><a href="team-48-1.html">光之园字幕组</a></td>
                                        <td width="12%"><a href="team-49-1.html">光之园字幕</a></td>
                                        <td width="12%"><a href="team-50-1.html">紫音動漫組</a></td>
                                        <td width="12%"><a href="team-51-1.html">花语发布</a></td>
            </tr>                <tr class="alt1">            <td width="12%"><a href="team-52-1.html">波洛咖啡厅字幕组</a></td>
                                        <td width="12%"><a href="team-53-1.html">c.c动漫</a></td>
                                        <td width="12%"><a href="team-54-1.html">泷沉琉璃MAD资源组</a></td>
                                        <td width="12%"><a href="team-55-1.html">動音漫影</a></td>
                                        <td width="12%"><a href="team-56-1.html">众神之王字幕组</a></td>
                                        <td width="12%"><a href="team-57-1.html">RH字幕组</a></td>
                                        <td width="12%"><a href="team-58-1.html">幻次元字幕组</a></td>
                                        <td width="12%"><a href="team-59-1.html">MT字幕组</a></td>
            </tr>                <tr class="alt2">            <td width="12%"><a href="team-60-1.html">SweetSub</a></td>
                                        <td width="12%"><a href="team-61-1.html">The ARC-V Project</a></td>
                                        <td width="12%"><a href="team-62-1.html">蓝白条论坛·玖组</a></td>
                                        <td width="12%"><a href="team-63-1.html">LoliHouse</a></td>
                                        <td width="12%"><a href="team-64-1.html">YWCN字幕组</a></td>
                                        <td width="12%"><a href="team-65-1.html"><span class="__cf_email__" data-cfemail="384f5d515a5778">[email&#160;protected]</span>海贼王微圈</a></td>
                                        <td width="12%"><a href="team-66-1.html">Ican字幕组</a></td>
                                        <td width="12%"><a href="team-67-1.html">虚数学区研究协会</a></td>
            </tr>                <tr class="alt1">            <td width="12%"><a href="team-68-1.html">星云字幕组</a></td>
                                        <td width="12%"><a href="team-69-1.html">天使动漫论坛</a></td>
                                        <td width="12%"><a href="team-70-1.html">冷番补完字幕组</a></td>
                                        <td width="12%"><a href="team-71-1.html">天行搬运</a></td>
                                        <td width="12%"><a href="team-72-1.html">铜锣字幕组</a></td>
                                        <td width="12%"><a href="team-73-1.html">雪恋</a></td>
                                        <td width="12%"><a href="team-74-1.html">狐狸小宮</a></td>
                                        <td width="12%"><a href="team-75-1.html">幻月字幕组</a></td>
            </tr>                <tr class="alt2">            <td width="12%"><a href="team-76-1.html">喵萌奶茶屋</a></td>
                                        <td width="12%"><a href="team-77-1.html">小愿8压制组</a></td>
                                        <td width="12%"><a href="team-78-1.html">歐克勒亞</a></td>
                                        <td width="12%"><a href="team-79-1.html">科学字幕组</a></td>
                                        <td width="12%"><a href="team-81-1.html">果冻字幕组</a></td>
                                        <td width="12%"><a href="team-82-1.html">NAZOrip</a></td>
                                        <td width="12%"><a href="team-83-1.html">树屋字幕组</a></td>
                                        <td width="12%"><a href="team-85-1.html">GiliGili</a></td>
            </tr>                <tr class="alt1">            <td width="12%"><a href="team-86-1.html">SmallNoob</a></td>
                                        <td width="12%"><a href="team-87-1.html">K&amp;W-RAWS</a></td>
                                        <td width="12%"><a href="team-88-1.html">咪梦动漫组</a></td>
                                        <td width="12%"><a href="team-89-1.html">萌FUN字幕组</a></td>
                                        <td width="12%"><a href="team-90-1.html">闲人字幕联萌</a></td>
                                        <td width="12%"><a href="team-91-1.html">青森小镇</a></td>
                                        <td width="12%"><a href="team-92-1.html">兔叽</a></td>
                                        <td width="12%"><a href="team-95-1.html">AcgN-樱花飘落</a></td>
            </tr>                <tr class="alt2">            <td width="12%"><a href="team-96-1.html">新番字幕组</a></td>
                                        <td width="12%"><a href="team-97-1.html">幻之字幕组</a></td>
                                        <td width="12%"><a href="team-98-1.html">ACG门户</a></td>
                                        <td width="12%"><a href="team-99-1.html">QS-Raws</a></td>
                                        <td width="12%"><a href="team-100-1.html">云光字幕组</a></td>
                                        <td width="12%"><a href="team-101-1.html">虾狐影视论坛</a></td>
                                        <td width="12%"><a href="team-102-1.html">404GROUP</a></td>
                                        <td width="12%"><a href="team-103-1.html">樱都字幕组</a></td>
            </tr>                <tr class="alt1">            <td width="12%"><a href="team-104-1.html">栗子动漫</a></td>
                                        <td width="12%"><a href="team-105-1.html">HFRgroup</a></td>
                                        <td width="12%"><a href="team-106-1.html">尽梨了字幕组</a></td>
                                        <td width="12%"><a href="team-107-1.html">MCE汉化组</a></td>
                                        <td width="12%"><a href="team-108-1.html">Niconeiko Works</a></td>
                                        <td width="12%"><a href="team-109-1.html">熱力社HOTNET</a></td>
                                        <td width="12%"><a href="team-111-1.html">神枫字幕组</a></td>
                                        <td width="12%"><a href="team-112-1.html">GMTeam</a></td>
            </tr>                <tr class="alt2">            <td width="12%"><a href="team-113-1.html">DAY字幕组</a></td>
                                        <td width="12%"><a href="team-114-1.html">BYYM发布组</a></td>
                                        <td width="12%"><a href="team-115-1.html">红鸟窝字幕组</a></td>
                                        <td width="12%"><a href="team-116-1.html">英配部落</a></td>
                                        <td width="12%"><a href="team-117-1.html">飞沐team</a></td>
                                        <td width="12%"><a href="team-118-1.html">溪流书房</a></td>
                                        <td width="12%"><a href="team-119-1.html">野比家字幕组</a></td>
                                        <td width="12%"><a href="team-120-1.html">膜鱼字幕组</a></td>
            </tr>                <tr class="alt1">            <td width="12%"><a href="team-121-1.html">爱咕字幕组</a></td>
                                        <td width="12%"><a href="team-122-1.html">仲夏动漫社</a></td>
                                        <td width="12%"><a href="team-123-1.html">夜莺家族字幕组</a></td>
                                        <td width="12%"><a href="team-124-1.html">STL字幕组</a></td>
                                        <td width="12%"><a href="team-125-1.html">亿万同人字幕组</a></td>
                                        <td width="12%"><a href="team-126-1.html">鹰小队翻译组</a></td>
                                        <td width="12%"><a href="team-127-1.html">NEX字幕组</a></td>
                                        <td width="12%"><a href="team-128-1.html">臭臭动漫整合</a></td>
            </tr>                <tr class="alt2">            <td width="12%"><a href="team-129-1.html">风筝字幕社</a></td>
                                        <td width="12%"><a href="team-130-1.html">嗷呜字幕组</a></td>
                                        <td width="12%"><a href="team-131-1.html">Yosuga字幕组</a></td>
                                        <td width="12%"><a href="team-132-1.html">空（Zero.sub）</a></td>
                                        <td width="12%"><a href="team-133-1.html">森之屋动画组</a></td>
                                        <td width="12%"><a href="team-134-1.html">SW字幕组</a></td>
                                        <td width="12%"><a href="team-135-1.html">霜庭云花Sub</a></td>
                                        <td width="12%"><a href="team-136-1.html">丝龙傲天二次元绝对领域</a></td>
            </tr>                <tr class="alt1">            <td width="12%"><a href="team-137-1.html">Amor</a></td>
                                        <td width="12%"><a href="team-138-1.html">离谱Sub</a></td>
                                        <td width="12%"><a href="team-139-1.html">orion个人发布</a></td>
                                        <td width="12%"><a href="team-140-1.html">幻城字幕组</a></td>
                                        <td width="12%"></td>
                                        <td width="12%"></td>
                                        <td width="12%"></td>
                                        <td width="12%"></td>
            </tr>        
        </tbody>
    </table>
</div>
<script data-cfasync="false" src="/cdn-cgi/scripts/5c5dd728/cloudflare-static/email-decode.min.js"></script><script type="text/javascript">panelDeploy('team', 0);</script>

<div class="box clear flink rounded">
    <h2 class="title">
        <span class="left">友情链接</span>
        <span class="right"><a href="javascript:void(0);" onclick="panelDeploy('flink', true);return false;" class="right"><img src="images/collapse.gif" id="panel_flink_button" /></a></span>
    </h2>
    <table class="list_style" cellpadding="0" cellspacing="0">
        <tbody id="panel_flink">
            
                                      <tr class="alt1">
                
                <td class="ftext">
                                                                        <a href="http://acgtracker.com/" target="_blank" title="ACGTracker">ACGTracker</a>&nbsp;&nbsp;
                                            
                </td>
              </tr>
                    </tbody>
    </table>
</div>
<script type="text/javascript">panelDeploy('flink', 0);</script>



</div>

<div class="clear text_center" style="margin-bottom:10px;"><style>
    .gg_download .label {
        display: inline;
        padding: .2em .4em .2em .4em;
        font-size: 12px;
        margin-right: 0.5em;
        border-radius: 5px;
        line-height: 1;
        text-align: center;
        vertical-align: baseline;
    }
    .gg_download .bg_grey {
        background-color: #d8d8d8;
    }
</style>

<script type="text/javascript">
    (function ($) {
        if (!Config) {
            return false;
        }
        if (Config['in_script'] !== 'show') {
            return false;
        }
        if (Config['mika_mode']['enabled'] && Config['user_script']['enabled'] && (Config['user_script']['revision'] !== $.cookie('user_script_rev'))) {
            $('#box_download div.basic_info ul')
                // .append('<li class="gg_download"><a href="https://union.115.com/?ac=invite_reg_share&uid=27720547" target="_blank" style="cursor: pointer;">115网盘</a></li>')
                .append('<li class="gg_download"><a href="https://look556.com/?from=kiss" target="_blank" style="cursor: pointer;">云端试播</a></li>');
        }
        else{
            var host = 'http://115.com/?tab=offline&mode=wangpan&union=27720547&download=';
            var magnet = 'magnet:?xt=urn:btih:' + Config['hash_id'];
            var link = host + encodeURIComponent(magnet);

            $('#box_download div.basic_info ul')
                // .append('<li><a href="' + link + '" target="_blank" style="cursor: pointer;">+115离线</a></li>')
                .append('<li><a href="https://look556.com/?from=kiss" target="_blank" style="cursor: pointer;">+在线预览</a></li>');
                // .append('<li><a href="http://app.tianbo17.com/?hmsr=kiss" target="_blank" style="cursor: pointer;">+在线看番</a></li>');
        }
    })(jQuery);
</script></div>


<div class="footer">
    <div class="left" style="font-weight: bold;">
        <p>＊MioBT＊联盟</p>
        <p>联系邮箱：<a href="/cdn-cgi/l/email-protection" class="__cf_email__" data-cfemail="503d393f32247e333f3d10373d31393c7e333f3d">[email&#160;protected]</a></p>
    </div>
    <div class="right">
        <p>
            
            <!--技术支持: <a href="http://www.miobt.com/addon.php?r=document/view&page=miobt-introduction" target="_blank">MioBT</a> |-->
            <!--技术支持: MioBT |-->
            技术支持: MioBT |
            <a href="javascript:void(0);" onclick="window.scroll(0,0);return false;">已到页末↑</a>
        </p>

                    <p>Copyright &copy;2008-2022&nbsp;<span style="font-weight: bold;"><a href="http://www.miobt.com/" target="_blank" class="copyright">WWW.MIOBT.COM</a></span>. All rights reserved.</p>
        

        <!-- Back to top start -->
        <script data-cfasync="false" src="/cdn-cgi/scripts/5c5dd728/cloudflare-static/email-decode.min.js"></script><script type="text/javascript">
            (function ($) {
                var max_width = 1420;
                var min_top = 300;
                var layer_hidden = true;
                var update_layer = function () {
                    if (($(window).width() >= max_width) && ($(window).scrollTop() > min_top)) {
//                        layer_hidden && $('#backTop').fadeIn(1000);
                        layer_hidden && $('#backTop').show();
                        layer_hidden = false;
                    }
                    else {
//                        layer_hidden || $('#backTop').fadeOut(1000);
                        layer_hidden || $('#backTop').hide();
                        layer_hidden = true;
                    }
                };

                $(window).resize(function () {
                    update_layer();
                });

                $(window).scroll(function () {
                    update_layer();
                });

                $(document).ready(function () {
                    update_layer();

                    $("#backTop").click(function () {
                        $("body,html").animate({scrollTop: 0}, 1000);
                        return false;
                    });
                });
            })(jQuery)
        </script>
        
        <style type="text/css">
            #backTop { opacity: 1;filter: alpha(opacity=100);position: fixed; _position:absolute; z-index:9999; bottom:0; right:0; display:none; height: 278px;width: 130px;background: url("images/back2top/back2top1.png") no-repeat; cursor:pointer}
            #backTop a { display:block; overflow:hidden}
            #backTop:hover {opacity: 1;filter: alpha(opacity=100)}
            #scrolltop {display:none!important}
        </style>
        <div id="backTop"><a href="javascript:;"></a></div>
        <!-- Back to top end -->
    </div>
</div>


</div>

<script type="text/javascript">
    (function ($) {
        $(".gg_canvas a[gg_url]").each(function () {
            $(this).attr({
                'target': '_blank',
                'rel': 'nofollow',
                'href': $(this).attr('gg_url')
            });
        });

        $(".gg_canvas-hidden").show();
    })(jQuery);
</script>

<script type="text/javascript">
    (function ($) {
        if (window.location.hash === '#disable_new_tab') {
            $.cookie('disable_new_tab', 1, {expires: 365, path: '/'});
        }

        $(document).ready(function () {
            if ($.cookie('disable_new_tab')) {
                $('a[target="_blank"]').removeAttr('target');
            }
        });
    })(jQuery);
</script>


<div style="display: none;">
    <script>
        var _hmt = _hmt || [];
        (function() {
            var hm = document.createElement("script");
            hm.src = "https://hm.baidu.com/hm.js?1fbd709fad69a2a27b0423b4ddb44b9b";
            var s = document.getElementsByTagName("script")[0];
            s.parentNode.insertBefore(hm, s);
        })();
    </script>
</div>

<script type="text/javascript">
    (function ($) {
        $(document).ready(function () {
            $("a[gg_url]").click(function () {
                _hmt.push(['_trackEvent', 'gg', 'click', $(this).attr('gg_url')]);
            });
            if (Config['gg_blocked']) {
                _hmt.push(['_trackEvent', 'gg', 'block']);
            } else {
                _hmt.push(['_trackEvent', 'gg', 'allow']);
            }
        });
    })(jQuery);
</script>
</body>
</html>
`

const dmhyHtml = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<title>首頁 - 動漫花園資源網 - 動漫愛好者的自由交流平台</title><meta name="keywords" content="動畫 漫畫 动漫音乐 动漫下載 动漫社区 游戏 animation comic game" />
<meta name="description" content="動漫花園資訊網是一個動漫愛好者交流的平台,提供最及時,最全面的動畫,漫畫,動漫音樂,動漫下載,BT,ED,動漫遊戲,資訊,分享,交流,讨论." />
<meta http-equiv="Content-Language" content="zh-TW" />
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" /><link href="/min/g=css&amp;v=10" media="screen" rel="stylesheet" type="text/css" /><script type="text/javascript" src="/min/g=js&amp;v=9"></script><!--[if IE]>
<link href="/css/iestyles.css" media="screen" rel="stylesheet" type="text/css" />
<![endif]-->
<link rel="shortcut icon" href="/favicon.ico" type="image/x-icon"/> 
<link rel="alternate" type="application/rss+xml" title="RSS Feeds" href="topics/rss/rss.xml" />
<link rel="search" type="application/opensearchdescription+xml"
	title="動漫花園" href="/js/dmhy.xml" />
</head>
<body>
<div class="container">
<div class="bg">
<div class="header">
<div class="headerleft" id="top"><a href="/" title="動漫花園資源網"><img alt="動漫花園資源網logo" src="/images/sitelogo.gif"/></a></div>
<div class="headerright">
<div class="links">
<span id="user_cm"></span>
<a href="/cms/page/name/programme.html" title="連載節目列表(圖像+字幕組分類)">每週番組</a>&nbsp;|
<a href="/cms/page/name/owari_bangumi.html" title="完結節目列表(圖像+字幕組,按首播季度分類)">完結番組</a>&nbsp;|
<a href="/cms/page/name/faq.html" title="發問前請先瀏覽" >常見問題<!--<font color=#ff9a4d"><i>更新</i></font>--></a>&nbsp;|
<a href="/cms/page/name/tos.html" title="請遵守" >服務條款</a>&nbsp;|
<a href="https://bbs.dmhy.org/thread.php?fid-504.html" target="_blank" title="請勿發表無關問題">意見建議</a>&nbsp;|
<a href="/cms/page/name/contact-us.html" title="有問題可聯系">聯繫我們</a>&nbsp;|
<a href="javascript:document.all?window.external.AddFavorite('https://share.dmhy.org', '動漫花園資源網'):(window.sidebar?window.sidebar.addPanel('動漫花園資源網', 'https://share.dmhy.org', ''):alert('Ctrl + D'))" title="添加收藏">收藏花園</a>&nbsp;|&nbsp;<span id="login_cm"></span>
<script>
var _hmt = _hmt || [];
(function() {
  var hm = document.createElement("script");
  hm.src = "https://hm.baidu.com/hm.js?e4918ccc327a268ee93dac21d5a7d53c";
  var s = document.getElementsByTagName("script")[0]; 
  s.parentNode.insertBefore(hm, s);
})();
</script>
<script type="text/javascript">
if($.cookie('uname')){
	str = '歡迎您!&nbsp;&nbsp;<span style="color:#FF9A4D;font-weight: bold;">'+decodeURIComponent($.cookie('uname'))+'</span>&nbsp;|&nbsp;';
	str += '<a href="/user">個人中心</a>&nbsp;|&nbsp;'
	$('#user_cm').append(str);
	login = '<a href="/user/logout">退出</a>';
	$('#login_cm').append(login);
}else{
	login = '<a href="/user/login">登入/註冊</a>';
	$('#login_cm').append(login);
}
</script>
</div>
<br />
<div class="ad">
<!--
<div id="950_ad" align="center">
<script> 
var pic_id = "/95080";
var random = Math.floor(Math.random()*1+1);
pic_id = pic_id + random + '.';
var extension = random == 1 ? 'jpg?6' : 'jpg?6';
var pic_html = pic_id + extension;
if(random == 1){
var href = "#";
}else if(random == 2){
var href = "#";
}else if(random == 3){
var href = "#";
}else{
var href = "#";
}
document.getElementById('950_ad').innerHTML='<a onclick="_hmt.push([\'_trackEvent\', \'ad\', \'click\', \'950\'])" href="'+href+'" target="_blank"><img src="' + pic_html + '" width="950" height="80" /></a>';
</script>
</div>
-->
<script src="https://js.kiwihk.net/?id=dmhy" type="text/javascript"></script>
<div class="kiwi-ad-wrapper-950x80 kiwi-banner1"></div></div>
</div>
<div class="top_sort">
<div class="menu">
<ul class="nav">
	<li> 
	<a class="sort-2" title="動畫" 
		href="/topics/list/sort_id/2"><font color=red>動畫</font></a>
		<ul>
		<li>
	<a class="sort-31" title="季度全集" 
		href="/topics/list/sort_id/31"><b><font color=red>季度全集</font></b></a>
	</li>
		</ul>
		</li>
	<li> 
	<a class="sort-3" title="漫畫" 
		href="/topics/list/sort_id/3"><font color=green>漫畫</font></a>
		<ul>
		<li>
	<a class="sort-41" title="港台原版" 
		href="/topics/list/sort_id/41"><b><font color=green>港台原版</font></b></a>
	</li>
		<li>
	<a class="sort-42" title="日文原版" 
		href="/topics/list/sort_id/42"><b><font color=green>日文原版</font></b></a>
	</li>
		</ul>
		</li>
	<li> 
	<a class="sort-4" title="音樂" 
		href="/topics/list/sort_id/4"><font color=purple>音樂</font></a>
		<ul>
		<li>
	<a class="sort-43" title="動漫音樂" 
		href="/topics/list/sort_id/43"><b><font color=purple>動漫音樂</font></b></a>
	</li>
		<li>
	<a class="sort-44" title="同人音樂" 
		href="/topics/list/sort_id/44"><b><font color=purple>同人音樂</font></b></a>
	</li>
		<li>
	<a class="sort-15" title="流行音樂" 
		href="/topics/list/sort_id/15"><b><font color=purple>流行音樂</font></b></a>
	</li>
		</ul>
		</li>
	<li> 
	<a class="sort-6" title="日劇" 
		href="/topics/list/sort_id/6"><font color=blue>日劇</font></a>
		</li>
	<li> 
	<a class="sort-7" title="ＲＡＷ" 
		href="/topics/list/sort_id/7"><font color=orange>ＲＡＷ</font></a>
		</li>
	<li> 
	<a class="sort-9" title="遊戲" 
		href="/topics/list/sort_id/9"><font color=#0eb9e7>遊戲</font></a>
		<ul>
		<li>
	<a class="sort-17" title="電腦遊戲" 
		href="/topics/list/sort_id/17"><font color=#0eb9e7>電腦遊戲</font></a>
	</li>
		<li>
	<a class="sort-18" title="電視遊戲" 
		href="/topics/list/sort_id/18"><font color=#0eb9e7>電視遊戲</font></a>
	</li>
		<li>
	<a class="sort-19" title="掌機遊戲" 
		href="/topics/list/sort_id/19"><font color=#0eb9e7>掌機遊戲</font></a>
	</li>
		<li>
	<a class="sort-20" title="網絡遊戲" 
		href="/topics/list/sort_id/20"><font color=#0eb9e7>網絡遊戲</font></a>
	</li>
		<li>
	<a class="sort-21" title="遊戲周邊" 
		href="/topics/list/sort_id/21"><font color=#0eb9e7>遊戲周邊</font></a>
	</li>
		</ul>
		</li>
	<li> 
	<a class="sort-12" title="特攝" 
		href="/topics/list/sort_id/12"><font color=brown>特攝</font></a>
		</li>
	<li> 
	<a class="sort-1" title="其他" 
		href="/topics/list/sort_id/1"><font color=black>其他</font></a>
		</li>
</ul>
</div>

<div id="expand-button" >
	<div id="team-expand" >&nbsp;</div>
	<div id="team-collapse" >&nbsp;</div>
</div>
<div id="team-dialog" ></div>

<script type="text/javascript">
jQuery('#expand-button').click(function(){
	jQuery('#team-expand').toggle();
	jQuery('#team-collapse').toggle();
	jQuery("#team-dialog").toggle();

	if(jQuery('#team-dialog').html() == ""){
	  var rc_url = '/team/navigate/';
	  jQuery('#team-dialog').load(rc_url,null,function(){});
	}
}); 
</script><div style="clear:both;"></div>
</div>
</div>
<div class="main">
<div id="1280_adv" align="center">
<script> 
var pic_id = "/VA";
var random = Math.floor(Math.random()*3+1);
pic_id = pic_id + random + '.';
var extension = random == 1 ? 'jpg?6' : 'jpg?6';
var pic_html = pic_id + extension;
if(random == 1){
var href = "https://l.tyrantdb.com/GJvGK4lq";
}else if(random == 2){
var href = "https://l.tyrantdb.com/LZ5apkbt";
}else if(random == 3){
var href = "https://l.tyrantdb.com/ja9W0js0";
}else{
var href = "https://l.tyrantdb.com/0JcV5avj";
}
document.getElementById('1280_adv').innerHTML='<a onclick="_hmt.push([\'_trackEvent\', \'ad\', \'click\', \'1280v\'])" href="'+href+'" target="_blank"><img src="' + pic_html + '" width="1280" height="120" /></a>';
</script>
<!---
<div class="kiwi-ad-wrapper-1280x120"></div>
--->
</div>
<script src="/js/jquery-1.11.1.min.js"></script>
<script src="/js/bd.js?2"></script>
<script type="text/javascript">
sunword='';monword='';tueword='';wedword='';thuword='';friword='';satword='';
sunarray = new Array();
monarray = new Array();
tuearray = new Array();
wedarray = new Array();
thuarray = new Array();
friarray = new Array();
satarray = new Array();
longarray = new Array();

/*******修改開始******/

/*
['圖片','動畫名','直接搜索連結','字幕組','官方公式']  範例：
___array.push(['http://share.dmhy.org/images/weekly/Gintama2017.gif','銀魂','%E9%8A%80%E9%AD%82%7CGintama','<a href="topics/list?keyword=%E9%93%B6%E9%AD%82+team_id%3A117">動漫花園</a>','http://www.sunrise-inc.co.jp/gintama/']);
*/

//其它

//星期日 (日)

sunarray.push(['https://share.dmhy.org/images/weekly/chibimaru.jpg','櫻桃小丸子','%E5%B0%8F%E4%B8%B8%E5%AD%90','<a href="/topics/list?keyword=%E5%B0%8F%E4%B8%B8%E5%AD%90+team_id%3A488">丸子家族</a><a href="/topics/list?keyword=%E5%B0%8F%E4%B8%B8%E5%AD%90+team_id%3A506">銀光</a>','http://chibimaru.tv/'] );
sunarray.push(['http://share.dmhy.org/images/weekly/pocket.gif','寵物小精靈','%E5%AF%B5%E7%89%A9%E5%B0%8F%E7%B2%BE%E9%9D%88','<a href="/topics/list?keyword=%E5%AF%B5%E7%89%A9%E5%B0%8F%E7%B2%BE%E9%9D%88+team_id%3A94">C2Club</a><a href="/topics/list?keyword=%E5%AF%B5%E7%89%A9%E5%B0%8F%E7%B2%BE%E9%9D%88+team_id%3A117">動漫花園</a><a href="/topics/list?keyword=%E5%AF%B5%E7%89%A9%E5%B0%8F%E7%B2%BE%E9%9D%88+team_id%3A88">Dymy</a><a href="/topics/list?keyword=%E5%AF%B5%E7%89%A9%E5%B0%8F%E7%B2%BE%E9%9D%88+team_id%3A376">ACB</a><a href="/topics/list?keyword=%E5%AF%B5%E7%89%A9%E5%B0%8F%E7%B2%BE%E9%9D%88+team_id%3A57">月光戀曲</a><a href="/topics/list?keyword=%E5%AF%B5%E7%89%A9%E5%B0%8F%E7%B2%BE%E9%9D%88+team_id%3A506">銀光</a><a href="/topics/list?keyword=%E5%AF%B5%E7%89%A9%E5%B0%8F%E7%B2%BE%E9%9D%88+team_id%3A447">夢幻戀櫻</a>','http://www.pokemon.co.jp/']);
sunarray.push(['https://share.dmhy.org/images/weekly/OnePiece.jpg','海賊王','%E6%B5%B7%E8%B3%8A%E7%8E%8B','<a href="/topics/list?keyword=%E6%B5%B7%E8%B3%8A%E7%8E%8B+team_id%3A34">楓雪連載</a><a href="/ topics/list?keyword=%E6%B5%B7%E8%B3%8A%E7%8E%8B+team_id%3A57">月光戀曲</a><a href="/topics/list?keyword=% E6%B5%B7%E8%B3%8A%E7%8E%8B+team_id%3A380">豬豬</a><a href="/topics/list?keyword=%E6%B5%B7%E8% B3%8A%E7%8E%8B+team_id%3A485">天空樹</a><a href="/topics/list?keyword=%E6%B5%B7%E8%B3%8A%E7%8E%8B+team_id%3A506">銀光</a>','http://www.toei-anim.co.jp/tv/onep/']);
sunarray.push(['https://share.dmhy.org/images/weekly/404.gif','Healin Good Precure光之美少女','Healin+%E5%85%89%E4%B9%8B','<a href="/topics/ list?keyword=Healin+%E5%85%89%E4%B9%8B+team_id%3A37">雪飄</a>','http://www.toei-anim.co.jp/tv/healingood_precure/']);
sunarray.push(['https://share.dmhy.org/images/weekly/404.gif','閃亮頻道','%E9%96%83%E4%BA%AE+%E9%A2%91%E9%81%93','<a href="/topics/list?keyword=%E9%96%83%E4%BA%AE+%E9%A2%91%E9%81%93+team_id%3A434">風之聖殿</a>','https://www.tv-tokyo.co.jp/anime/prichan/']);
sunarray.push(['https://share.dmhy.org/images/weekly/404.gif','甜夢貓','%E7%94%9C%E5%A4%A2%E8%B2%93','<a href="/topics/list?keyword=%E7%94%9C%E5%A4%A2%E8%B2%93+team_id%3A604">cc動漫</a><a href="/topics/list?keyword=%E7%94%9C%E5%A4%A2%E8%B2%93+team_id%3A675">AikatsuFans</a>','https://mewkledreamy.com/']);
sunarray.push(['https://share.dmhy.org/images/weekly/404.gif','Rebirth','Rebirth','<a href="/topics/list?keyword=Rebirth+team_id%3A669">喵萌奶茶屋</a>','https://rebirth-fy.com/anime/']);
sunarray.push(['https://share.dmhy.org/images/weekly/404.gif','數碼寶貝：','%E6%95%B8%E7%A2%BC+%E5%A4%A7%E5%86%92%E9%9A%AA','<a href="/topics/list?keyword =%E6%95%B8%E7%A2%BC+%E5%A4%A7%E5%86%92%E9%9A%AA+team_id%3A626">馴獸師聯盟</a><a href="/topics/list?keyword=%E6%95%B8%E7%A2%BC+%E5%A4%A7%E5%86%92%E9%9A%AA+team_id%3A454">風車</a><a href="/ topics/list?keyword=%E6%95%B8%E7%A2%BC+%E5%A4%A7%E5%86%92%E9%9A%AA+team_id%3A731">星空</a><a href="/topics/list?keyword=%E6%95%B8%E7%A2%BC+%E5%A4%A7%E5%86%92%E9%9A%AA+team_id%3A619">桜都</a>',' http://www.toei-anim.co.jp/tv/digimon/']);
sunarray.push(['https://share.dmhy.org/images/weekly/404.gif','巨人族的新娘','%E5%B7%A8%E4%BA%BA%E6%97%8F%E7%9A%84%E6%96%B0%E5%A8%98','<a href="/topics/list?keyword=%E5%B7%A8%E4%BA%BA%E6%97%8F%E7%9A%84%E6%96%B0%E5%A8%98+team_id%3A185">極影</a><a href="/topics/list?keyword=%E5%B7%A8%E4%BA%BA%E6%97%8F%E7%9A%84%E6%96%B0%E5%A8%98+team_id%3A241">幻櫻</a>','https://kyojinzoku.cf-anime.com/']);
sunarray.push(['https://share.dmhy.org/images/weekly/404.gif','魔物娘的醫生','%E9%AD%94%E7%89%A9%E5%A8%98%E7%9A%84%E5%8C%BB%E7%94%9F','<a href="/topics/list?keyword=%E9%AD%94%E7%89%A9%E5%A8%98%E7%9A%84%E9%86%AB%E7%94%9F+team_id%3A710">咪梦动漫组</a><a href="/topics/list?keyword=%E9%AD%94%E7%89%A9%E5%A8%98%E7%9A%84%E9%86%AB%E7%94%9F+team_id%3A619">桜都</a><a href="/topics/list?keyword=%E9%AD%94%E7%89%A9%E5%A8%98%E7%9A%84%E9%86%AB%E7%94%9F+team_id%3A604">c.c动漫</a>','https://mon-isha-anime.com/']);






//星期壹 (月)

monarray.push(['https://share.dmhy.org/images/weekly/404.gif','普通都市柏传超说R','%E8%B6%85%E6%99%AE%E9%80%9A%E9%83%BD%E5%B8%82%E6%9F%8F%E4%BC%A0%E8%AF%B4','<a href="/topics/list?keyword=%E8%B6%85%E6%99%AE%E9%80%9A%E9%83%BD%E5%B8%82%E6%9F%8F%E4%BC%A0%E8%AF%B4+team_id%3A619">桜都</a>','https://kashiden.dekimachi.com/']);
monarray.push(['https://share.dmhy.org/images/weekly/404.gif','the god of high school','the+god+of+high+school','<a href="/topics/list?keyword=the+god+of+high+school+team_id%3A710">咪梦</a><a href="/topics/list?keyword=the+god+of+high+school+team_id%3A731">星空</a><a href="/topics/list?keyword=the+god+of+high+school+team_id%3A669">喵萌奶茶屋</a>','http://goh-anime.com/']);
monarray.push(['https://share.dmhy.org/images/weekly/404.gif','魔法水果籃 2nd season','%E6%B0%B4%E6%9E%9C','<a href="/topics/list?keyword=%E6%B0%B4%E6%9E%9C+team_id%3A604">c.c动漫</a><a href="/topics/list?keyword=%E6%B0%B4%E6%9E%9C+team_id%3A731">星空</a><a href="/topics/list?keyword=%E6%B0%B4%E6%9E%9C+team_id%3A669">喵萌奶茶屋</a><a href="/topics/list?keyword=%E6%B0%B4%E6%9E%9C+team_id%3A407">DHR</a>','https://fruba.jp/']);






//星期二 (火)

tuearray.push(['http://share.dmhy.org/images/weekly/BlackClover.gif','黑色五葉草','%E4%BA%94%E8%91%89%E8%8D%89','<a href="/topics/list?keyword=%E4%BA%94%E8%91%89%E8%8D%89+team_id%3A117">動漫花園</a><a href="/topics/list?keyword=%E4%BA%94%E8%91%89%E8%8D%89+team_id%3A502">F宅</a><a href="/topics/list?keyword=%E4%BA%94%E8%91%89%E8%8D%89+team_id%3A663">八重櫻</a><a href="/topics/list?keyword=%E4%BA%94%E8%91%89%E8%8D%89+team_id%3A651">追新番</a><a href="/topics/list?keyword=%E4%BA%94%E8%91%89%E8%8D%89+team_id%3A185">極影</a>','http://bclover.jp/']);
tuearray.push(['https://share.dmhy.org/images/weekly/404.gif','Shadowverse','Shadowverse','<a href="/topics/list?keyword=Shadowverse+team_id%3A604">c.c动漫</a><a href="topics/list?keyword=Shadowverse+team_id%3A619">桜都</a>','https://anime.shadowverse.jp/']);
tuearray.push(['https://share.dmhy.org/images/weekly/404.gif','魔法律事務所','%E9%AD%94%E6%B3%95%E5%BE%8B%E4%BA%8B%E5%8B%99%E6%89%80','<a href="/topics/list?keyword=%E9%AD%94%E6%B3%95%E5%BE%8B%E4%BA%8B%E5%8B%99%E6%89%80+team_id%3A"></a>','http://mahouritsu.com/']);
tuearray.push(['https://share.dmhy.org/images/weekly/404.gif','馬娘四格','%E9%A6%AC%E5%A8%98%E5%9B%9B%E6%A0%BC','<a href="/topics/list?keyword=%E9%A6%AC%E5%A8%98%E5%9B%9B%E6%A0%BC+team_id%3A669">喵萌奶茶屋</a><a href="/topics/list?keyword=%E9%A6%AC%E5%A8%98%E5%9B%9B%E6%A0%BC+team_id%3A407">DHR</a>','https://umamusume.jp/umayon/']);
tuearray.push(['https://share.dmhy.org/images/weekly/404.gif','放學后海堤日記','%E9%A6%AC%E5%A8%98%E5%9B%9B%E6%A0%BC%E6%94%BE%E5%AD%A6%E5%90%8E+%E6%97%A5%E8%AE%B0','<a href="/topics/list?keyword=%E6%94%BE%E5%AD%A6%E5%90%8E+%E6%97%A5%E8%AE%B0+team_id%3A321">轻之国度</a><a href="/topics/list?keyword=%E6%94%BE%E5%AD%A6%E5%90%8E+%E6%97%A5%E8%AE%B0+team_id%3A619">桜都</a><a href="/topics/list?keyword=%E6%94%BE%E5%AD%A6%E5%90%8E+%E6%97%A5%E8%AE%B0+team_id%3A283">千夏</a>','https://teibotv.com/']);



//星期三 (水)

wedarray.push(['http://share.dmhy.org/images/weekly/boruto.gif','BORUTO 火影新世代','%E5%8D%9A%E4%BA%BA%7CBoruto','<a href="/topics/list?keyword=%E5%8D%9A%E4%BA%BA%7CBoruto+team_id%3A407">DHR</a><a href="/topics/list?keyword=%E5%8D%9A%E4%BA%BA%7CBoruto+team_id%3A370">旋風</a>','http://www.tv-tokyo.co.jp/anime/boruto/']);
wedarray.push(['https://share.dmhy.org/images/weekly/404.gif','籃球少年王','%E7%B1%83%E7%90%83%E5%B0%91%E5%B9%B4%E7%8E%8B','<a href="/topics/list?keyword=%E7%B1%83%E7%90%83%E5%B0%91%E5%B9%B4%E7%8E%8B+team_id%3A604">c.c动漫</a>','http://ahirunosora.jp/']);
wedarray.push(['https://share.dmhy.org/images/weekly/404.gif','Re: 从零开始的异世界生活 第二季','%E5%BC%82%E4%B8%96%E7%95%8C%E7%94%9F%E6%B4%BB','<a href="/topics/list?keyword=%E5%BC%82%E4%B8%96%E7%95%8C%E7%94%9F%E6%B4%BB+team_id%3A731">星空</a><a href="/topics/list?keyword=%E5%BC%82%E4%B8%96%E7%95%8C%E7%94%9F%E6%B4%BB+team_id%3A47">愛戀&漫貓</a>','http://re-zero-anime.jp/tv2/']);
wedarray.push(['https://share.dmhy.org/images/weekly/404.gif','沒落要塞 DECA-DENCE','DECA-DENCE','<a href="/topics/list?keyword=DECA-DENCE+team_id%3A669">喵萌奶茶屋</a><a href="/topics/list?keyword=DECA-DENCE+team_id%3A520">豌豆&風之聖殿</a><a href="/topics/list?keyword=DECA-DENCE+team_id%3A185">極影</a><a href="/topics/list?keyword=DECA-DENCE+team_id%3A283">千夏</a>','http://decadence-anime.com/']);
wedarray.push(['https://share.dmhy.org/images/weekly/404.gif','大欺詐師 GREAT PRETENDER','GREAT+PRETENDER','<a href="/topics/list?keyword=GREAT+PRETENDER+team_id%3A650">SweetSub</a><a href="/topics/list?keyword=GREAT+PRETENDER+team_id%3A88">dymy</a><a href="/topics/list?keyword=GREAT+PRETENDER+team_id%3A185">極影</a>','http://www.greatpretender.jp/']);
wedarray.push(['https://share.dmhy.org/images/weekly/404.gif','GIBIATE 獵魔武士','%E7%8D%B5%E9%AD%94%E6%AD%A6%E5%A3%AB','<a href="/topics/list?keyword=%E7%8D%B5%E9%AD%94%E6%AD%A6%E5%A3%AB"</a>','https://gibiate.com/anime/jp/']);
wedarray.push(['https://share.dmhy.org/images/weekly/404.gif','戀与製作人','%E6%88%80%E4%B8%8E%E8%A3%BD%E4%BD%9C%E4%BA%BA','<a href="/topics/list?keyword=%E6%88%80%E4%B8%8E%E8%A3%BD%E4%BD%9C%E4%BA%BA"</a>','https://koipro-anime.love/']);


//星期四 (木)

thuarray.push(['http://share.dmhy.org/images/weekly/404.gif','Gundam Build Divers Re-RISE','Gundam+Build+Divers+Re-RISE','<a href="/topics/list?keyword=Gundam+Build+Divers+Re-RISE+team_id%3A706">K&W-RAWS</a>','http://gundam-bd.net/']);
thuarray.push(['http://share.dmhy.org/images/weekly/404.gif','非槍人生/No Guns Life','No+Guns+Life','<a href="/topics/list?keyword=No+Guns+Life+team_id%3A669">喵萌奶茶屋</a><a href="/topics/list?keyword=No+Guns+Life+team_id%3A185">極影</a>','http://nogunslife.com/']);
thuarray.push(['http://share.dmhy.org/images/weekly/404.gif','我的青春戀愛物語果然有問題。完','%E9%9D%92%E6%98%A5%E6%88%80%E6%84%9B','<a href="/topics/list?keyword=%E9%9D%92%E6%98%A5%E6%88%80%E6%84%9B+team_id%3A731">星空</a><a href="/topics/list?keyword=%E9%9D%92%E6%98%A5%E6%88%80%E6%84%9B+team_id%3A321">轻之国度</a><a href="topics/list?keyword=%E9%9D%92%E6%98%A5%E6%88%80%E6%84%9B+team_id%3A185">極影</a><a href="/topics/list?keyword=%E9%9D%92%E6%98%A5%E6%88%80%E6%84%9B+team_id%3A303">動漫國</a><a href="/topics/list?keyword=%E9%9D%92%E6%98%A5%E6%88%80%E6%84%9B+team_id%3A0">漫游&澄空</a><a href="/topics/list?keyword=%E9%9D%92%E6%98%A5%E6%81%8B%E7%88%B1+team_id%3A669">喵萌奶茶屋</a><a href="/topics/list?keyword=%E9%9D%92%E6%98%A5%E6%88%80%E6%84%9B+team_id%3A619">桜都</a>','http://www.tbs.co.jp/anime/oregairu/']);
thuarray.push(['http://share.dmhy.org/images/weekly/404.gif','富豪刑事','%E5%AF%8C%E8%B1%AA%E5%88%91%E4%BA%8B','<a href="/topics/list?keyword=%E5%AF%8C%E8%B1%AA%E5%88%91%E4%BA%8B+team_id%3A731">星空</a><a href="/topics/list?keyword=%E5%AF%8C%E8%B1%AA%E5%88%91%E4%BA%8B+team_id%3A185">極影</a><a href="/topics/list?keyword=%E5%AF%8C%E8%B1%AA%E5%88%91%E4%BA%8B+team_id%3A241">幻櫻</a>','https://www.fugoukeiji-bul.com/']);



//星期五 (金)
friarray.push(['http://share.dmhy.org/images/weekly/404.gif','科學超電磁炮T','%E8%B6%85%E7%94%B5%E7%A3%81%E7%82%AE','<a href="/topics/list?keyword=%E8%B6%85%E7%94%B5%E7%A3%81%E7%82%AET+team_id%3A303">澄空學園&動漫國</a><a href="/topics/list?keyword=%E8%B6%85%E7%94%B5%E7%A3%81%E7%82%AET+team_id%3A283">千夏</a><a href="/topics/list?keyword=%E8%B6%85%E7%94%B5%E7%A3%81%E7%82%AET+team_id%3A185">極影</a><a href="/topics/list?keyword=%E8%B6%85%E7%94%B5%E7%A3%81%E7%82%AET+team_id%3A619">桜都</a>','https://toaru-project.com/railgun_t/']);
friarray.push(['http://share.dmhy.org/images/weekly/404.gif','天晴爛漫','%E5%A4%A9%E6%99%B4%E7%88%9B%E6%BC%AB','<a href="/topics/list?keyword=%E5%A4%A9%E6%99%B4%E7%88%9B%E6%BC%AB+team_id%3A669">喵萌奶茶屋</a><a href="/topics/list?keyword=%E5%A4%A9%E6%99%B4%E7%88%9B%E6%BC%AB+team_id%3A785">野比家</a>','http://appareranman.com/']);
friarray.push(['http://share.dmhy.org/images/weekly/404.gif','弩級戰隊HXEROS','HXEROS','<a href="/topics/list?keyword=HXEROS+%E7%99%BD%E6%81%8B">白恋</a><a href="/topics/list?keyword=HXEROS+team_id%3A619">桜都</a>','https://hxeros.com/']);
friarray.push(['http://share.dmhy.org/images/weekly/404.gif','文豪與鍊金術師～審判的齒輪～','%E5%AF%A9%E5%88%A4%E7%9A%84%E9%BD%92%E8%BC%AA','<a href="/topics/list?keyword=%E5%AF%A9%E5%88%A4%E7%9A%84%E9%BD%92%E8%BC%AA+team_id%3A604">c.c动漫</a>','https://anime-bungo-alchemist.com/']);
friarray.push(['http://share.dmhy.org/images/weekly/404.gif','炎炎消防隊 二之章','%E7%82%8E%E7%82%8E%E6%B6%88%E9%98%B2%E9%9A%8A','<a href="/topics/list?keyword=%E7%82%8E%E7%82%8E%E6%B6%88%E9%98%B2%E9%9A%8A+team_id%3A731">星空</a><a href="/topics/list?keyword=%E7%82%8E%E7%82%8E%E6%B6%88%E9%98%B2%E9%9A%8A+team_id%3A8">dymy</a>','http://www.fireforce-anime.jp/']);
friarray.push(['http://share.dmhy.org/images/weekly/404.gif','食戟之靈 豪之皿','%E9%A3%9F%E6%88%9F%E4%B9%8B%E9%9D%88+%E8%B1%AA%E4%B9%8B%E7%9A%BF','<a href="/topics/list?keyword=%E9%A3%9F%E6%88%9F%E4%B9%8B%E9%9D%88+%E8%B1%AA%E4%B9%8B%E7%9A%BF+team_id%3A657">LoliHouse</a><a href="/topics/list?keyword=%E9%A3%9F%E6%88%9F%E4%B9%8B%E9%9D%88+%E8%B1%AA%E4%B9%8B%E7%9A%BF+team_id%3A604">c.c动漫</a><a href="/topics/list?keyword=%E9%A3%9F%E6%88%9F%E4%B9%8B%E9%9D%88+%E8%B1%AA%E4%B9%8B%E7%9A%BF+team_id%3A185">極影</a><a href="/topics/list?keyword=%E9%A3%9F%E6%88%9F%E4%B9%8B%E9%9D%88+%E8%B1%AA%E4%B9%8B%E7%9A%BF+team_id%3A49">雪飘&澄空&华盟</a>','http://shokugekinosoma.com/5thplate/']);
friarray.push(['http://share.dmhy.org/images/weekly/404.gif','宇崎學妹想要玩','%E5%AE%87%E5%B4%8E','<a href="/topics/list?keyword=%E5%AE%87%E5%B4%8E+team_id%3A669">喵萌奶茶屋</a><a href="/topics/list?keyword=%E5%AE%87%E5%B4%8E+team_id%3A241">幻櫻</a><a href="/topics/list?keyword=%E5%AE%87%E5%B4%8E+team_id%3A321">轻之国度</a><a href="/topics/list?keyword=%E5%AE%87%E5%B4%8E+team_id%3A47">愛戀&漫貓</a><a href="/topics/list?keyword=%E5%AE%87%E5%B4%8E+team_id%3A185">極影</a>','https://uzakichan.com/']);
friarray.push(['http://share.dmhy.org/images/weekly/404.gif','租借女友','%E7%A7%9F+%E5%A5%B3%E5%8F%8B','<a href="/topics/list?keyword=%E7%A7%9F+%E5%A5%B3%E5%8F%8B+team_id%3A669">喵萌奶茶屋</a><a href="/topics/list?keyword=%E7%A7%9F+%E5%A5%B3%E5%8F%8B+team_id%3A185">極影</a><a href="/topics/list?keyword=%E7%A7%9F+%E5%A5%B3%E5%8F%8B+team_id%3A283">千夏</a><a href="/topics/list?keyword=%E7%A7%9F+%E5%A5%B3%E5%8F%8B+team_id%3A731">星空</a><a href="/topics/list?keyword=%E7%A7%9F+%E5%A5%B3%E5%8F%8B+team_id%3A604">c.c动漫</a><a href="/topics/list?keyword=%E7%A7%9F+%E5%A5%B3%E5%8F%8B+team_id%3A241">幻櫻</a><a href="/topics/list?keyword=%E7%A7%9F+%E5%A5%B3%E5%8F%8B+team_id%3A407">DHR</a>','http://kanokari-official.com/']);
friarray.push(['http://share.dmhy.org/images/weekly/404.gif','彼得·格里爾的賢者時間','%E8%B3%A2%E8%80%85%E6%99%82%E9%96%93','<a href="/topics/list?keyword=%E8%B3%A2%E8%80%85%E6%99%82%E9%96%93+team_id%3A321">轻之国度</a><a href="topics/list?keyword=%E8%B3%A2%E8%80%85%E6%99%82%E9%96%93+team_id%3A669">喵萌奶茶屋</a>','http://petergrill-anime.jp/']);




//星期六 (土)
satarray.push(['http://share.dmhy.org/images/weekly/Doraemon.jpg','Doraemon','doraemon%7C%E5%95%A6a%E5%A4%A2','<a href="/topics/list?keyword=Doraemon+team_id%3A453">天空</a><a href="/topics/list?keyword=Doraemon+team_id%3A506">銀光</a><a href="/topics/list?keyword=Doraemon+team_id%3A561">釘鐺</a><a href="/topics/list?keyword=Doraemon+team_id%3A574">夢藍</a><a href="/topics/list?keyword=Doraemon+team_id%3A531">清藍</a><a href="/topics/list?keyword=Doraemon+team_id%3A564">風翼</a>','http://www.tv-asahi.co.jp/doraemon/']);
satarray.push(['http://share.dmhy.org/images/weekly/717087432ab9c0e226db9e2d1d81b8c2.jpg','蠟筆小新','%E8%A0%9F%E7%AD%86%E5%B0%8F%E6%96%B0','<a href="/topics/list?keyword=Crayon_Shin-chan+team_id%3A453">天空</a><a href="/topics/list?keyword=%E8%9C%A1%E7%AC%94%E5%B0%8F%E6%96%B0+team_id%3A63">琵琶行</a><a href="/topics/list?keyword=%E8%9C%A1%E7%AC%94%E5%B0%8F%E6%96%B0+team_id%3A574">夢藍</a><a href="/topics/list?keyword=%E8%9C%A1%E7%AC%94%E5%B0%8F%E6%96%B0+team_id%3A506">銀光</a>','http://www.tv-asahi.co.jp/shinchan/']);
satarray.push(['http://share.dmhy.org/images/weekly/conan.gif','名偵探柯南','%E6%9F%AF%E5%8D%97','<a href="/topics/list?keyword=%E6%9F%AF%E5%8D%97+team_id%3A241">幻櫻</a><a href="/topics/list?keyword=%E6%9F%AF%E5%8D%97+team_id%3A75">柯南事務所</a><a href="/topics/list?keyword=%E6%9F%AF%E5%8D%97+team_id%3A380">豬豬</a><a href="/topics/list?keyword=%E6%9F%AF%E5%8D%97+team_id%3A299">星光</a><a href="/topics/list?keyword=%E6%9F%AF%E5%8D%97+team_id%3A359">極速</a><a href="/topics/list?keyword=%E6%9F%AF%E5%8D%97+team_id%3A412">ByCanon</a><a href="/topics/list?keyword=%E6%9F%AF%E5%8D%97+team_id%3A506">銀光</a><a href="/topics/list?keyword=%E6%9F%AF%E5%8D%97+team_id%3A454">風車</a><a href="/topics/list?keyword=%E6%9F%AF%E5%8D%97+team_id%3A576">銀色子彈</a>','http://www.ytv.co.jp/conan/']);
satarray.push(['http://share.dmhy.org/images/weekly/404.gif','棒球大聯盟2nd 第二季','%E6%A3%92%E7%90%83%E5%A4%A7%E8%81%AF%E7%9B%9F2nd+%E7%AC%AC%E4%BA%8C%E5%AD%A3','<a href="/topics/list?keyword=%E6%A3%92%E7%90%83%E5%A4%A7%E8%81%AF%E7%9B%9F2nd+%E7%AC%AC%E4%BA%8C%E5%AD%A3+team_id%3A88">dymy</a><a href="/topics/list?keyword=%E6%A3%92%E7%90%83%E5%A4%A7%E8%81%AF%E7%9B%9F2nd+%E7%AC%AC%E4%BA%8C%E5%AD%A3+%E6%B0%B8%E6%81%92">永恒译制</a><a href="/https://share.dmhy.org/topics/list?keyword=%E6%A3%92%E7%90%83%E5%A4%A7%E8%81%AF%E7%9B%9F2nd+%E7%AC%AC%E4%BA%8C%E5%AD%A3+team_id%3A786">棒聯貼吧</a>','http://www.nhk.or.jp/anime/major2nd/']);
satarray.push(['http://share.dmhy.org/images/weekly/404.gif','遊戲王SEVENS','%E6%B8%B8%E6%88%8F%E7%8E%8B+SEVENS','<a href="/topics/list?keyword=%E6%B8%B8%E6%88%8F%E7%8E%8B+SEVENS+team_id%3A673">VRAINSTORM</a>','https://www.tv-tokyo.co.jp/anime/yugioh-sevens/']);
satarray.push(['http://share.dmhy.org/images/weekly/404.gif','寶石幻想','Lapis+Re%3ALiGHTs','<a href="/topics/list?keyword=Lapis+Re%3ALiGHTs+team_id%3A669">喵萌奶茶屋</a>','https://www.lapisrelights.com/anime/']);
satarray.push(['http://share.dmhy.org/images/weekly/404.gif','魔王學院的不適任者','%E9%AD%94%E7%8E%8B%E5%AD%A6%E9%99%A2','<a href="/topics/list?keyword=%E9%AD%94%E7%8E%8B%E5%AD%A6%E9%99%A2+team_id%3A669">喵萌奶茶屋</a><a href="/topics/list?keyword=%E9%AD%94%E7%8E%8B%E5%AD%A6%E9%99%A2+team_id%3A321">轻之国度</a><a href="/topics/list?keyword=%E9%AD%94%E7%8E%8B%E5%AD%A6%E9%99%A2+team_id%3A731">星空</a>','https://maohgakuin.com/']);
satarray.push(['http://share.dmhy.org/images/weekly/404.gif','刀劍神域Alicization War of Underworld','Alicization','<a href="/topics/list?keyword=Alicization+team_id%3A151">悠哈璃羽＆拉斯观测组</a><a href="/topics/list?keyword=Alicization+team_id%3A185">極影</a><a href="/topics/list?keyword=Alicization+team_id%3A37">澄空&華盟&雪飄</a><a href="/topics/list?keyword=Alicization+team_id%3A619">桜都</a><a href="/topics/list?keyword=Alicization+team_id%3A604">c.c动漫</a><a href="/topics/list?keyword=Alicization+team_id%3A432">傲嬌零&自由</a>','https://maohgakuin.com/']);








</script>

<div>
<style type="text/css" media="screen">
table.jmd{
width:100%;
table-layout:fixed;
}
table.jmd tr th{
padding: 1px;
width: 6%;
background-color: #7E99BE;
color: white;
}
table.jmd tr td{
padding: 1px;
word-wrap: break-word;
word-break: break-all;
}
table.jmd .today{
background-color: #ff6;
}
table.jmd a{
border: 1px solid #ffa500;
padding: 1px;
margin: 1px;
line-height: 20px;
}
table.jmd a:hover { 
color: orange;
}
</style>


<div id="mini_jmd" class="table fl" style="width:99%">
<div class="nav_title"><span class="fr"> 
<a href="/cms/page/name/programme.html">詳細每週番組</a></span>新番資源索引 - 
<script language="JavaScript">
today=new Date();
year=today.getYear()
if (year < 1000)
year+=1900
function initArray(){
this.length=initArray.arguments.length
for(var i=0;i<this.length;i++)
this[i+1]=initArray.arguments[i]}
var d=new initArray("星期日","星期一","星期二","星期三","星期四","星期五","星期六");
document.write(" 公元 ",year," 年 ",today.getMonth()+1," 月 ",today.getDate()," 日 ",d[today.getDay()+1]);
</script>


 | 便捷搜索: <a href="/topics/list?keyword=BDMV%7CBDISO%7CBDBOX%7CBDRIP%7CBD">BD</a>,  <a href="/topics/list?keyword=DVDISO%7CDVDRIP%7CDVD">DVD</a>,  <a href="/topics/list?keyword=OVA%7COAD">OVA/OAD</a>,  </a><a href="/topics/list?keyword=%E5%8A%87%E5%A0%B4%7C%E9%9B%BB%E5%BD%B1%7CMOVIE">劇場版</a>


</div>
<table class="jmd"></table>
</div>


<script type="text/javascript">

table = $('<table class="jmd_base" style="width:100%;"></table>');


td = $('<td></td>');
for(f=0;f<sunarray.length;f++){
if(sunarray[f][1]!='&nbsp;'){td.append('<a href="/topics/list?keyword='+sunarray[f][2]+'">'+sunarray[f][1]+'</a>')};
}
$('<tr></tr>').append('<th>週日 （日）</th>').append(td).appendTo(table);

td = $('<td></td>');
for(f=0;f<monarray.length;f++){
if(monarray[f][1]!='&nbsp;'){td.append('<a href="/topics/list?keyword='+monarray[f][2]+'">'+monarray[f][1]+'</a>')};
}
$('<tr></tr>').append('<th>週一（月）</th>').append(td).appendTo(table);


td = $('<td></td>');
for(f=0;f<tuearray.length;f++){
if(tuearray[f][1]!='&nbsp;'){td.append('<a href="/topics/list?keyword='+tuearray[f][2]+'">'+tuearray[f][1]+'</a>')};
}
$('<tr></tr>').append('<th>週二（火）</th>').append(td).appendTo(table);

td = $('<td></td>');
for(f=0;f<wedarray.length;f++){
if(wedarray[f][1]!='&nbsp;'){td.append('<a href="/topics/list?keyword='+wedarray[f][2]+'">'+wedarray[f][1]+'</a>')};
}
$('<tr></tr>').append('<th>週三（水）</th>').append(td).appendTo(table);

td = $('<td></td>');
for(f=0;f<thuarray.length;f++){
if(thuarray[f][1]!='&nbsp;'){td.append('<a href="/topics/list?keyword='+thuarray[f][2]+'">'+thuarray[f][1]+'</a>')};
}
$('<tr></tr>').append('<th>週四（木）</th>').append(td).appendTo(table);

td = $('<td></td>');
for(f=0;f<friarray.length;f++){
if(friarray[f][1]!='&nbsp;'){td.append('<a href="/topics/list?keyword='+friarray[f][2]+'">'+friarray[f][1]+'</a>')};
}
$('<tr></tr>').append('<th>週五（金）</th>').append(td).appendTo(table);

td = $('<td></td>');
for(f=0;f<satarray.length;f++){
if(satarray[f][1]!='&nbsp;'){td.append('<a href="/topics/list?keyword='+satarray[f][2]+'">'+satarray[f][1]+'</a>')};
}
$('<tr></tr>').append('<th>週六（土）</th>').append(td).appendTo(table);


table.appendTo('#mini_jmd').hide();

d = new Date();
day = d.getDay();

$(".jmd")
.append($($(".jmd_base tr")[(d.getDay()+5) % 7]).clone())
.append($($(".jmd_base tr")[(d.getDay()+6) % 7]).clone())
.append($($(".jmd_base tr")[(d.getDay()) % 7]).clone().addClass('today'))
.append($($(".jmd_base tr")[(d.getDay()+1) % 7]).clone())

$(".jmd tr:even").addClass("even");
$(".jmd tr:odd").addClass("odd");
</script>
</div>


<div  class="clear" ></div><div class="quick_search">
<form action="/topics/list" method="get">
	<input id="keyword" type="text" class="quick_input" maxlength="50" x-webkit-speech lang="zh-tw" x-webkit-grammar="builtin:translate" 
		name="keyword" value="" />
	<input type="submit" class="formButton" value=" 搜尋 " />&nbsp;<a href="javascript:void(0)" onclick="showHideAdvSearch()">高級&nbsp;<span id="searchMode">▼</span></a>
	<div id="AdvSearch">
	</div>
</form>
</div>
<script language="javascript">
var AdvSearchLoaded = false;
jQuery("#keyword").autocomplete("/topics/search-suggest/",{
    max:10,
    width:366,
    multipleSeparator: ' ',  
	dataType: 'json',
	selectFirst: false,
	parse: function(data) {
		var rows = [];
		for(var i=0; i<data.length; i++){
			rows[rows.length] = { 
				data:data[i], 
				value:data[i].search_keyword, 
				result:data[i].search_keyword 
			}; 
        }
		return rows;
	},
	formatItem: function(item) {
		return format(item);
	}
}).result(function(e, item) {
	// fill value into hide field
});
function format(row) {
	return "<span class=\"count\" >" + row.search_count + "個結果</span>"+row.search_keyword;
}
function showHideAdvSearch(){
	var advsearch = jQuery("#AdvSearch");
	if(!AdvSearchLoaded){
		advsearch.load("/topics/advanced-search?team_id=0&sort_id=0&orderby=");
		AdvSearchLoaded = true;
	}
	advsearch.slideToggle(function(){$("#searchMode").text(advsearch.is(":visible") ? "▲" : "▼");});
}
</script><div id="1280_ad" align="center">
<a href="https://7segu.taobao.com?mm_sycmid=1_116976_add3d8643cf663a5da153928c0e70775"><img src="/sheng1.gif?7" ></a>
<!---
<script> 
var pic_id = "/MY";
var random = Math.floor(Math.random()*1+1);
pic_id = pic_id + random + '.';
var extension = random == 1 ? 'gif?7' : 'gif?7';
var pic_html = pic_id + extension;
if(random == 1){
var href = "https://ant.antss028.com/aff/a5Cg";
}else if(random == 2){
var href = "https://ant.antss028.com/aff/a5Cg";
}else if(random == 3){
var href = "https://ant.antss028.com/aff/a5Cg";
}else{
var href = "https://ant.antss028.com/aff/a5Cg";
}
document.getElementById('1280_ad').innerHTML='<a onclick="_hmt.push([\'_trackEvent\', \'ad\', \'click\', \'1280\'])" href="'+href+'" target="_blank"><img src="' + pic_html + '" width="1280" height="120" /></a>';
</script>
</div>
<!---
<div class="kiwi-ad-wrapper-1280x120"></div>
---><div class="clear" >
<div style="width: 120px;float: left;padding-bottom: 3px;">
<a href="/topics/add" style="display: block;"><img src="/images/btn_postbt.gif" alt="發佈BT資源" ></a>
</div>
<div class="announce">
<div class="inner">
<marquee id="announce_marquee" direction="left" scrolldelay="1" scrollamount="2" 
onmouseout="{this.start();}" onmouseover="{this.stop();}">
<a target="_blank" href="/announce#ann140">
<strong>番組表開放編輯</strong>
( by otakuzero, 2020/06/15 2:13 )</a>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
<a target="_blank" href="/announce#ann117">
<strong>動畫資源標題格式說明</strong>
( by otakuzero, 2012/07/06 20:27 )</a>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
<a target="_blank" href="/announce#ann108">
<strong>搜索小提示</strong>
( by otakuzero, 2010/10/23 15:37 )</a>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
</marquee>
</div>
</div>
</div>
<div id="overlay" style="display: none;">请稍候... </div>

<div class="table clear">
<div class="nav_title">
	<div class="fl">
						上一頁		第1頁 
				<a href="/topics/list/page/2">下一頁</a>
			</div>
	<div class="fr">
		<a href="/topics/list/recent/1+day">最近一天</a>&nbsp;|&nbsp;
		<a href="/topics/list/recent/1+week">最近一週</a>&nbsp;|&nbsp;
		<a href="/topics/rss/rss.xml">
		<img style="vertical-align:bottom;" src="/images/rss.gif" alt="rss"> 订阅这个分类</a>&nbsp;|&nbsp;
		<a href="https://t.me/dmhy_org">订阅花园TG频道</a>
		
	</div>
</div>

<div class="clear">
<table  style="width:100%;"  class="tablesorter" id="topic_list" cellpadding="0" cellspacing="1" border="0" width="" frame="void">
	<thead>
		<tr>
			<th class="{sorter: 'magicDate'}" width="96"><span class="title">張貼日期</span></th>
			<th class="{sorter: 'text'}" width="5%"><span class="title">分類</span></th>
			<th class="{sorter: 'text'}" ><span class="title">標題</span></th>
			<th class="{sorter: false}" width="6%" nowrap="nowrap"><span class="title">磁鏈</span></th>
			<th class="{sorter: 'filesize', sortInitialOrder: 'desc'}"  width="7%"><span class="title">大小</span></th>
			<th class="{sorter: 'digit', sortInitialOrder: 'desc'}" width="5%"><span class="title">種子</span></th>
			<th class="{sorter: 'digit', sortInitialOrder: 'desc'}" width="5%"><span class="title">下載</span></th>
			<th class="{sorter: 'digit', sortInitialOrder: 'desc'}" width="5%"><span class="title">完成</span></th>
			<th class="{sorter: 'text'}" width="9%"><span class="title">發佈人</span></th>
		</tr>
	</thead>
	<tbody>
			<tr class="">
			<td width="98">
			今天 22:31			<span style="display: none;">2022/02/21 22:31</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/454" >
				风车字幕组</a></span>				
				<a href="/topics/view/593422_10_43_1080P_MP4.html"  target="_blank" >
				[風車字幕組十一周年][10月新番][半妖的夜叉姬][第43集][暗轉的舞台][1080P][繁體][MP4]</a>
				<span style="color: gray;">約1條評論</span>			</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:57D3BF47GWWDIYE6SSZMMYXS6TA4A6PE&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Fopen.nyaatorrents.info%3A6544%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=http%3A%2F%2Ftracker.kisssub.org%2Fannounce&tr=http%3A%2F%2Ftracker.ex.ua%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.yify-torrents.com%2Fannounce&tr=http%3A%2F%2Fannounce.torrentsmd.com%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.leechers-paradise.org%3A6969%2Fannounce&tr=http%3A%2F%2Fretracker.telecom.kz%2Fannounce&tr=http%3A%2F%2Fbt.careland.com.cn%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.ktxp.com%3A6868%2Fannounce&tr=http%3A%2F%2Ftracker.dmhy.org%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.xfsub.com%3A6868%2Fannounce&tr=http%3A%2F%2Fmgtracker.org%3A2710%2Fannounce&tr=http%3A%2F%2Fmgtracker.org%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.grepler.com%3A6969%2Fannounce&tr=http%3A%2F%2Fbt.evrl.to%2Fannounce&tr=http%3A%2F%2Ftracker.mg64.net%3A6881%2Fannounce&tr=http%3A%2F%2Ftracker.torrentyorg.pl%2Fannounce&tr=http%3A%2F%2Ftracker.baravik.org%3A6970%2Fannounce&tr=http%3A%2F%2Ftracker.filetracker.pl%3A8089%2Fannounce&tr=http%3A%2F%2Ftracker1.wasabii.com.tw%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker2.wasabii.com.tw%3A6969%2Fannounce&tr=http%3A%2F%2Fretracker.local%2Fannounce&tr=http%3A%2F%2F37.19.5.139%3A6969%2Fannounce&tr=http%3A%2F%2F37.19.5.155%3A6881%2Fannounce&tr=http%3A%2F%2F46.4.109.148%3A6969%2Fannounce&tr=http%3A%2F%2F210.244.71.26%3A6969%2Fannounce&tr=http%3A%2F%2F210.244.71.25%3A6969%2Fannounce&tr=http%3A%2F%2F125.227.35.196%3A6969%2Fannounce&tr=http%3A%2F%2F213.159.215.198%3A6970%2Fannounce&tr=http%3A%2F%2Ftorrentsmd.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&tr=http%3A%2F%2Ftracker.internetwarriors.net%3A1337%2Fannounce&tr=http%3A%2F%2Fexplodie.org%3A6969%2Fannounce&tr=http%3A%2F%2Fopentracker.i2p.rocks%3A6969%2Fannounce&tr=http%3A%2F%2F5rt.tace.ru%3A60889%2Fannounce&tr=http%3A%2F%2Fopen.acgnxtracker.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.tiny-vps.com%3A6969%2Fannounce&tr=http%3A%2F%2Frt.tace.ru%3A80%2Fannounce&tr=https%3A%2F%2Ftrakx.herokuapp.com%3A443%2Fannounce&tr=http%3A%2F%2Fwww.loushao.net%3A8080%2Fannounce&tr=http%3A%2F%2Fh4.trakx.nibba.trade%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.zerobytes.xyz%3A1337%2Fannounce&tr=http%3A%2F%2Fvps02.net.orel.ru%3A80%2Fannounce&tr=https%3A%2F%2Fw.wwwww.wtf%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.tamersunion.org%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.sloppyta.co%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.renfei.net%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.nitrix.me%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.nanoha.org%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.lelux.fi%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.imgoingto.icu%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.hama3.net%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.foreverpirates.co%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.coalition.space%3A443%2Fannounce&tr=https%3A%2F%2Faaa.army%3A8866%2Fannounce&tr=https%3A%2F%2F1337.abcvg.info%3A443%2Fannounce&tr=http%3A%2F%2Fvpn.flying-datacenter.de%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker2.dler.org%3A80%2Fannounce&tr=http%3A%2F%2Ftracker1.bt.moack.co.kr%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.vraphim.com%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.sloppyta.co%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.skyts.net%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.noobsubs.net%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.loadbt.com%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.lelux.fi%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.gbitt.info%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.dler.org%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.anonwebz.xyz%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker-cdn.moeking.me%3A2095%2Fannounce&tr=http%3A%2F%2Ftorrenttracker.nwc.acsalaska.net%3A6969%2Fannounce&tr=http%3A%2F%2Ftorrentclub.online%3A54123%2Fannounce&tr=http%3A%2F%2Ft.overflow.biz%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%3A80%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce&tr=http%3A%2F%2Fretracker.sevstar.net%3A2710%2Fannounce&tr=http%3A%2F%2Fpow7.com%3A80%2Fannounce&tr=http%3A%2F%2Fns3107607.ip-54-36-126.eu%3A6969%2Fannounce&tr=http%3A%2F%2Fmail2.zelenaya.net%3A80%2Fannounce&tr=http%3A%2F%2Fbt.okmp3.ru%3A2710%2Fannounce&tr=http%3A%2F%2Fbt.100.pet%3A2710%2Fannounce&tr=http%3A%2F%2Fbobbialbano.com%3A6969%2Fannounce&tr=http%3A%2F%2Faaa.army%3A8866%2Fannounce&tr=https%3A%2F%2Ftr.bangumi.moe%3A9696%2Fannounce&tr=https%3A%2F%2Ftr.ready4.icu%3A443%2Fannounce&tr=http%3A%2F%2Ftracker.bt4g.com%3A2095%2Fannounce&tr=http%3A%2F%2Fderpyradio.net%3A6969%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:efc7b0979f35ac34609e94b2c662f2f4c1c079e4&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">627MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">-</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">-</span></td>
			<td nowrap="nowrap" align="center">-</td>
			<td align="center"><a href="/topics/list/user_id/661609">工藤新一</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 22:30			<span style="display: none;">2022/02/21 22:30</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/454" >
				风车字幕组</a></span>				
				<a href="/topics/view/593421_10_43_1080P_MP4.html"  target="_blank" >
				[风车字幕组十一周年][10月新番][半妖的夜叉姬][第43集][暗转的舞台][1080P][简体][MP4]</a>
							</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:CG7NKEQFKLJFWTEZ4ZFSRRW4LAKOZYTK&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Fopen.nyaatorrents.info%3A6544%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=http%3A%2F%2Ftracker.kisssub.org%2Fannounce&tr=http%3A%2F%2Ftracker.ex.ua%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.yify-torrents.com%2Fannounce&tr=http%3A%2F%2Fannounce.torrentsmd.com%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.leechers-paradise.org%3A6969%2Fannounce&tr=http%3A%2F%2Fretracker.telecom.kz%2Fannounce&tr=http%3A%2F%2Fbt.careland.com.cn%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.ktxp.com%3A6868%2Fannounce&tr=http%3A%2F%2Ftracker.dmhy.org%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.xfsub.com%3A6868%2Fannounce&tr=http%3A%2F%2Fmgtracker.org%3A2710%2Fannounce&tr=http%3A%2F%2Fmgtracker.org%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.grepler.com%3A6969%2Fannounce&tr=http%3A%2F%2Fbt.evrl.to%2Fannounce&tr=http%3A%2F%2Ftracker.mg64.net%3A6881%2Fannounce&tr=http%3A%2F%2Ftracker.torrentyorg.pl%2Fannounce&tr=http%3A%2F%2Ftracker.baravik.org%3A6970%2Fannounce&tr=http%3A%2F%2Ftracker.filetracker.pl%3A8089%2Fannounce&tr=http%3A%2F%2Ftracker1.wasabii.com.tw%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker2.wasabii.com.tw%3A6969%2Fannounce&tr=http%3A%2F%2Fretracker.local%2Fannounce&tr=http%3A%2F%2F37.19.5.139%3A6969%2Fannounce&tr=http%3A%2F%2F37.19.5.155%3A6881%2Fannounce&tr=http%3A%2F%2F46.4.109.148%3A6969%2Fannounce&tr=http%3A%2F%2F210.244.71.26%3A6969%2Fannounce&tr=http%3A%2F%2F210.244.71.25%3A6969%2Fannounce&tr=http%3A%2F%2F125.227.35.196%3A6969%2Fannounce&tr=http%3A%2F%2F213.159.215.198%3A6970%2Fannounce&tr=http%3A%2F%2Ftorrentsmd.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&tr=http%3A%2F%2Ftracker.internetwarriors.net%3A1337%2Fannounce&tr=http%3A%2F%2Fexplodie.org%3A6969%2Fannounce&tr=http%3A%2F%2Fopentracker.i2p.rocks%3A6969%2Fannounce&tr=http%3A%2F%2F5rt.tace.ru%3A60889%2Fannounce&tr=http%3A%2F%2Fopen.acgnxtracker.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.tiny-vps.com%3A6969%2Fannounce&tr=http%3A%2F%2Frt.tace.ru%3A80%2Fannounce&tr=https%3A%2F%2Ftrakx.herokuapp.com%3A443%2Fannounce&tr=http%3A%2F%2Fwww.loushao.net%3A8080%2Fannounce&tr=http%3A%2F%2Fh4.trakx.nibba.trade%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.zerobytes.xyz%3A1337%2Fannounce&tr=http%3A%2F%2Fvps02.net.orel.ru%3A80%2Fannounce&tr=https%3A%2F%2Fw.wwwww.wtf%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.tamersunion.org%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.sloppyta.co%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.renfei.net%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.nitrix.me%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.nanoha.org%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.lelux.fi%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.imgoingto.icu%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.hama3.net%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.foreverpirates.co%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.coalition.space%3A443%2Fannounce&tr=https%3A%2F%2Faaa.army%3A8866%2Fannounce&tr=https%3A%2F%2F1337.abcvg.info%3A443%2Fannounce&tr=http%3A%2F%2Fvpn.flying-datacenter.de%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker2.dler.org%3A80%2Fannounce&tr=http%3A%2F%2Ftracker1.bt.moack.co.kr%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.vraphim.com%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.sloppyta.co%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.skyts.net%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.noobsubs.net%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.loadbt.com%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.lelux.fi%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.gbitt.info%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.dler.org%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.anonwebz.xyz%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker-cdn.moeking.me%3A2095%2Fannounce&tr=http%3A%2F%2Ftorrenttracker.nwc.acsalaska.net%3A6969%2Fannounce&tr=http%3A%2F%2Ftorrentclub.online%3A54123%2Fannounce&tr=http%3A%2F%2Ft.overflow.biz%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%3A80%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce&tr=http%3A%2F%2Fretracker.sevstar.net%3A2710%2Fannounce&tr=http%3A%2F%2Fpow7.com%3A80%2Fannounce&tr=http%3A%2F%2Fns3107607.ip-54-36-126.eu%3A6969%2Fannounce&tr=http%3A%2F%2Fmail2.zelenaya.net%3A80%2Fannounce&tr=http%3A%2F%2Fbt.okmp3.ru%3A2710%2Fannounce&tr=http%3A%2F%2Fbt.100.pet%3A2710%2Fannounce&tr=http%3A%2F%2Fbobbialbano.com%3A6969%2Fannounce&tr=http%3A%2F%2Faaa.army%3A8866%2Fannounce&tr=https%3A%2F%2Ftr.bangumi.moe%3A9696%2Fannounce&tr=https%3A%2F%2Ftr.ready4.icu%3A443%2Fannounce&tr=http%3A%2F%2Ftracker.bt4g.com%3A2095%2Fannounce&tr=http%3A%2F%2Fderpyradio.net%3A6969%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:11bed5120552d25b4c99e64b28c6dc5814ece26a&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">626.9MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">-</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">-</span></td>
			<td nowrap="nowrap" align="center">-</td>
			<td align="center"><a href="/topics/list/user_id/661609">工藤新一</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 22:28			<span style="display: none;">2022/02/21 22:28</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/454" >
				风车字幕组</a></span>				
				<a href="/topics/view/593420_10_42_1080P_MP4.html"  target="_blank" >
				[風車字幕組十一周年][10月新番][半妖的夜叉姬][第42集][崩壞的時之風車][1080P][繁體][MP4]</a>
				<span style="color: gray;">約1條評論</span>			</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:QFKRDPWFVKLZ6F5OAE63ISOSPOFFKKQZ&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Fopen.nyaatorrents.info%3A6544%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=http%3A%2F%2Ftracker.kisssub.org%2Fannounce&tr=http%3A%2F%2Ftracker.ex.ua%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.yify-torrents.com%2Fannounce&tr=http%3A%2F%2Fannounce.torrentsmd.com%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.leechers-paradise.org%3A6969%2Fannounce&tr=http%3A%2F%2Fretracker.telecom.kz%2Fannounce&tr=http%3A%2F%2Fbt.careland.com.cn%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.ktxp.com%3A6868%2Fannounce&tr=http%3A%2F%2Ftracker.dmhy.org%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.xfsub.com%3A6868%2Fannounce&tr=http%3A%2F%2Fmgtracker.org%3A2710%2Fannounce&tr=http%3A%2F%2Fmgtracker.org%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.grepler.com%3A6969%2Fannounce&tr=http%3A%2F%2Fbt.evrl.to%2Fannounce&tr=http%3A%2F%2Ftracker.mg64.net%3A6881%2Fannounce&tr=http%3A%2F%2Ftracker.torrentyorg.pl%2Fannounce&tr=http%3A%2F%2Ftracker.baravik.org%3A6970%2Fannounce&tr=http%3A%2F%2Ftracker.filetracker.pl%3A8089%2Fannounce&tr=http%3A%2F%2Ftracker1.wasabii.com.tw%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker2.wasabii.com.tw%3A6969%2Fannounce&tr=http%3A%2F%2Fretracker.local%2Fannounce&tr=http%3A%2F%2F37.19.5.139%3A6969%2Fannounce&tr=http%3A%2F%2F37.19.5.155%3A6881%2Fannounce&tr=http%3A%2F%2F46.4.109.148%3A6969%2Fannounce&tr=http%3A%2F%2F210.244.71.26%3A6969%2Fannounce&tr=http%3A%2F%2F210.244.71.25%3A6969%2Fannounce&tr=http%3A%2F%2F125.227.35.196%3A6969%2Fannounce&tr=http%3A%2F%2F213.159.215.198%3A6970%2Fannounce&tr=http%3A%2F%2Ftorrentsmd.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&tr=http%3A%2F%2Ftracker.internetwarriors.net%3A1337%2Fannounce&tr=http%3A%2F%2Fexplodie.org%3A6969%2Fannounce&tr=http%3A%2F%2Fopentracker.i2p.rocks%3A6969%2Fannounce&tr=http%3A%2F%2F5rt.tace.ru%3A60889%2Fannounce&tr=http%3A%2F%2Fopen.acgnxtracker.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.tiny-vps.com%3A6969%2Fannounce&tr=http%3A%2F%2Frt.tace.ru%3A80%2Fannounce&tr=https%3A%2F%2Ftrakx.herokuapp.com%3A443%2Fannounce&tr=http%3A%2F%2Fwww.loushao.net%3A8080%2Fannounce&tr=http%3A%2F%2Fh4.trakx.nibba.trade%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.zerobytes.xyz%3A1337%2Fannounce&tr=http%3A%2F%2Fvps02.net.orel.ru%3A80%2Fannounce&tr=https%3A%2F%2Fw.wwwww.wtf%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.tamersunion.org%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.sloppyta.co%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.renfei.net%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.nitrix.me%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.nanoha.org%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.lelux.fi%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.imgoingto.icu%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.hama3.net%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.foreverpirates.co%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.coalition.space%3A443%2Fannounce&tr=https%3A%2F%2Faaa.army%3A8866%2Fannounce&tr=https%3A%2F%2F1337.abcvg.info%3A443%2Fannounce&tr=http%3A%2F%2Fvpn.flying-datacenter.de%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker2.dler.org%3A80%2Fannounce&tr=http%3A%2F%2Ftracker1.bt.moack.co.kr%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.vraphim.com%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.sloppyta.co%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.skyts.net%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.noobsubs.net%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.loadbt.com%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.lelux.fi%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.gbitt.info%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.dler.org%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.anonwebz.xyz%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker-cdn.moeking.me%3A2095%2Fannounce&tr=http%3A%2F%2Ftorrenttracker.nwc.acsalaska.net%3A6969%2Fannounce&tr=http%3A%2F%2Ftorrentclub.online%3A54123%2Fannounce&tr=http%3A%2F%2Ft.overflow.biz%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%3A80%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce&tr=http%3A%2F%2Fretracker.sevstar.net%3A2710%2Fannounce&tr=http%3A%2F%2Fpow7.com%3A80%2Fannounce&tr=http%3A%2F%2Fns3107607.ip-54-36-126.eu%3A6969%2Fannounce&tr=http%3A%2F%2Fmail2.zelenaya.net%3A80%2Fannounce&tr=http%3A%2F%2Fbt.okmp3.ru%3A2710%2Fannounce&tr=http%3A%2F%2Fbt.100.pet%3A2710%2Fannounce&tr=http%3A%2F%2Fbobbialbano.com%3A6969%2Fannounce&tr=http%3A%2F%2Faaa.army%3A8866%2Fannounce&tr=https%3A%2F%2Ftr.bangumi.moe%3A9696%2Fannounce&tr=https%3A%2F%2Ftr.ready4.icu%3A443%2Fannounce&tr=http%3A%2F%2Ftracker.bt4g.com%3A2095%2Fannounce&tr=http%3A%2F%2Fderpyradio.net%3A6969%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:815511bec5aa979f17ae013db449d27b8a552a19&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">700.3MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">-</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">-</span></td>
			<td nowrap="nowrap" align="center">-</td>
			<td align="center"><a href="/topics/list/user_id/661609">工藤新一</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 22:26			<span style="display: none;">2022/02/21 22:26</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/454" >
				风车字幕组</a></span>				
				<a href="/topics/view/593419_10_42_1080P_MP4.html"  target="_blank" >
				[风车字幕组十一周年][10月新番][半妖的夜叉姬][第42集][崩坏的时之风车][1080P][简体][MP4]</a>
							</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:PWC7F6OZMUHKAALUA5P3KZDV7QNXB2A3&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Fopen.nyaatorrents.info%3A6544%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=http%3A%2F%2Ftracker.kisssub.org%2Fannounce&tr=http%3A%2F%2Ftracker.ex.ua%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.yify-torrents.com%2Fannounce&tr=http%3A%2F%2Fannounce.torrentsmd.com%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.leechers-paradise.org%3A6969%2Fannounce&tr=http%3A%2F%2Fretracker.telecom.kz%2Fannounce&tr=http%3A%2F%2Fbt.careland.com.cn%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.ktxp.com%3A6868%2Fannounce&tr=http%3A%2F%2Ftracker.dmhy.org%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.xfsub.com%3A6868%2Fannounce&tr=http%3A%2F%2Fmgtracker.org%3A2710%2Fannounce&tr=http%3A%2F%2Fmgtracker.org%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.grepler.com%3A6969%2Fannounce&tr=http%3A%2F%2Fbt.evrl.to%2Fannounce&tr=http%3A%2F%2Ftracker.mg64.net%3A6881%2Fannounce&tr=http%3A%2F%2Ftracker.torrentyorg.pl%2Fannounce&tr=http%3A%2F%2Ftracker.baravik.org%3A6970%2Fannounce&tr=http%3A%2F%2Ftracker.filetracker.pl%3A8089%2Fannounce&tr=http%3A%2F%2Ftracker1.wasabii.com.tw%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker2.wasabii.com.tw%3A6969%2Fannounce&tr=http%3A%2F%2Fretracker.local%2Fannounce&tr=http%3A%2F%2F37.19.5.139%3A6969%2Fannounce&tr=http%3A%2F%2F37.19.5.155%3A6881%2Fannounce&tr=http%3A%2F%2F46.4.109.148%3A6969%2Fannounce&tr=http%3A%2F%2F210.244.71.26%3A6969%2Fannounce&tr=http%3A%2F%2F210.244.71.25%3A6969%2Fannounce&tr=http%3A%2F%2F125.227.35.196%3A6969%2Fannounce&tr=http%3A%2F%2F213.159.215.198%3A6970%2Fannounce&tr=http%3A%2F%2Ftorrentsmd.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&tr=http%3A%2F%2Ftracker.internetwarriors.net%3A1337%2Fannounce&tr=http%3A%2F%2Fexplodie.org%3A6969%2Fannounce&tr=http%3A%2F%2Fopentracker.i2p.rocks%3A6969%2Fannounce&tr=http%3A%2F%2F5rt.tace.ru%3A60889%2Fannounce&tr=http%3A%2F%2Fopen.acgnxtracker.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.tiny-vps.com%3A6969%2Fannounce&tr=http%3A%2F%2Frt.tace.ru%3A80%2Fannounce&tr=https%3A%2F%2Ftrakx.herokuapp.com%3A443%2Fannounce&tr=http%3A%2F%2Fwww.loushao.net%3A8080%2Fannounce&tr=http%3A%2F%2Fh4.trakx.nibba.trade%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.zerobytes.xyz%3A1337%2Fannounce&tr=http%3A%2F%2Fvps02.net.orel.ru%3A80%2Fannounce&tr=https%3A%2F%2Fw.wwwww.wtf%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.tamersunion.org%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.sloppyta.co%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.renfei.net%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.nitrix.me%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.nanoha.org%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.lelux.fi%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.imgoingto.icu%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.hama3.net%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.foreverpirates.co%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.coalition.space%3A443%2Fannounce&tr=https%3A%2F%2Faaa.army%3A8866%2Fannounce&tr=https%3A%2F%2F1337.abcvg.info%3A443%2Fannounce&tr=http%3A%2F%2Fvpn.flying-datacenter.de%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker2.dler.org%3A80%2Fannounce&tr=http%3A%2F%2Ftracker1.bt.moack.co.kr%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.vraphim.com%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.sloppyta.co%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.skyts.net%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.noobsubs.net%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.loadbt.com%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.lelux.fi%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.gbitt.info%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.dler.org%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.anonwebz.xyz%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker-cdn.moeking.me%3A2095%2Fannounce&tr=http%3A%2F%2Ftorrenttracker.nwc.acsalaska.net%3A6969%2Fannounce&tr=http%3A%2F%2Ftorrentclub.online%3A54123%2Fannounce&tr=http%3A%2F%2Ft.overflow.biz%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%3A80%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce&tr=http%3A%2F%2Fretracker.sevstar.net%3A2710%2Fannounce&tr=http%3A%2F%2Fpow7.com%3A80%2Fannounce&tr=http%3A%2F%2Fns3107607.ip-54-36-126.eu%3A6969%2Fannounce&tr=http%3A%2F%2Fmail2.zelenaya.net%3A80%2Fannounce&tr=http%3A%2F%2Fbt.okmp3.ru%3A2710%2Fannounce&tr=http%3A%2F%2Fbt.100.pet%3A2710%2Fannounce&tr=http%3A%2F%2Fbobbialbano.com%3A6969%2Fannounce&tr=http%3A%2F%2Faaa.army%3A8866%2Fannounce&tr=https%3A%2F%2Ftr.bangumi.moe%3A9696%2Fannounce&tr=https%3A%2F%2Ftr.ready4.icu%3A443%2Fannounce&tr=http%3A%2F%2Ftracker.bt4g.com%3A2095%2Fannounce&tr=http%3A%2F%2Fderpyradio.net%3A6969%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:7d85f2f9d9650ea00174075fb56475fc1b70e81b&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">700.3MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">-</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">-</span></td>
			<td nowrap="nowrap" align="center">-</td>
			<td align="center"><a href="/topics/list/user_id/661609">工藤新一</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 22:23			<span style="display: none;">2022/02/21 22:23</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/454" >
				风车字幕组</a></span>				
				<a href="/topics/view/593418_17_1080P_MP4.html"  target="_blank" >
				[風車字幕組十一周年][數碼寶貝幽靈遊戲][第17集][極寒地獄][1080P][繁體][MP4]</a>
							</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:RTZ3KQIPDQLTWN7NUVLQXXF5A6XTWIRC&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Fopen.nyaatorrents.info%3A6544%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=http%3A%2F%2Ftracker.kisssub.org%2Fannounce&tr=http%3A%2F%2Ftracker.ex.ua%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.yify-torrents.com%2Fannounce&tr=http%3A%2F%2Fannounce.torrentsmd.com%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.leechers-paradise.org%3A6969%2Fannounce&tr=http%3A%2F%2Fretracker.telecom.kz%2Fannounce&tr=http%3A%2F%2Fbt.careland.com.cn%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.ktxp.com%3A6868%2Fannounce&tr=http%3A%2F%2Ftracker.dmhy.org%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.xfsub.com%3A6868%2Fannounce&tr=http%3A%2F%2Fmgtracker.org%3A2710%2Fannounce&tr=http%3A%2F%2Fmgtracker.org%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.grepler.com%3A6969%2Fannounce&tr=http%3A%2F%2Fbt.evrl.to%2Fannounce&tr=http%3A%2F%2Ftracker.mg64.net%3A6881%2Fannounce&tr=http%3A%2F%2Ftracker.torrentyorg.pl%2Fannounce&tr=http%3A%2F%2Ftracker.baravik.org%3A6970%2Fannounce&tr=http%3A%2F%2Ftracker.filetracker.pl%3A8089%2Fannounce&tr=http%3A%2F%2Ftracker1.wasabii.com.tw%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker2.wasabii.com.tw%3A6969%2Fannounce&tr=http%3A%2F%2Fretracker.local%2Fannounce&tr=http%3A%2F%2F37.19.5.139%3A6969%2Fannounce&tr=http%3A%2F%2F37.19.5.155%3A6881%2Fannounce&tr=http%3A%2F%2F46.4.109.148%3A6969%2Fannounce&tr=http%3A%2F%2F210.244.71.26%3A6969%2Fannounce&tr=http%3A%2F%2F210.244.71.25%3A6969%2Fannounce&tr=http%3A%2F%2F125.227.35.196%3A6969%2Fannounce&tr=http%3A%2F%2F213.159.215.198%3A6970%2Fannounce&tr=http%3A%2F%2Ftorrentsmd.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&tr=http%3A%2F%2Ftracker.internetwarriors.net%3A1337%2Fannounce&tr=http%3A%2F%2Fexplodie.org%3A6969%2Fannounce&tr=http%3A%2F%2Fopentracker.i2p.rocks%3A6969%2Fannounce&tr=http%3A%2F%2F5rt.tace.ru%3A60889%2Fannounce&tr=http%3A%2F%2Fopen.acgnxtracker.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.tiny-vps.com%3A6969%2Fannounce&tr=http%3A%2F%2Frt.tace.ru%3A80%2Fannounce&tr=https%3A%2F%2Ftrakx.herokuapp.com%3A443%2Fannounce&tr=http%3A%2F%2Fwww.loushao.net%3A8080%2Fannounce&tr=http%3A%2F%2Fh4.trakx.nibba.trade%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.zerobytes.xyz%3A1337%2Fannounce&tr=http%3A%2F%2Fvps02.net.orel.ru%3A80%2Fannounce&tr=https%3A%2F%2Fw.wwwww.wtf%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.tamersunion.org%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.sloppyta.co%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.renfei.net%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.nitrix.me%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.nanoha.org%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.lelux.fi%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.imgoingto.icu%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.hama3.net%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.foreverpirates.co%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.coalition.space%3A443%2Fannounce&tr=https%3A%2F%2Faaa.army%3A8866%2Fannounce&tr=https%3A%2F%2F1337.abcvg.info%3A443%2Fannounce&tr=http%3A%2F%2Fvpn.flying-datacenter.de%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker2.dler.org%3A80%2Fannounce&tr=http%3A%2F%2Ftracker1.bt.moack.co.kr%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.vraphim.com%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.sloppyta.co%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.skyts.net%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.noobsubs.net%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.loadbt.com%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.lelux.fi%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.gbitt.info%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.dler.org%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.anonwebz.xyz%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker-cdn.moeking.me%3A2095%2Fannounce&tr=http%3A%2F%2Ftorrenttracker.nwc.acsalaska.net%3A6969%2Fannounce&tr=http%3A%2F%2Ftorrentclub.online%3A54123%2Fannounce&tr=http%3A%2F%2Ft.overflow.biz%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%3A80%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce&tr=http%3A%2F%2Fretracker.sevstar.net%3A2710%2Fannounce&tr=http%3A%2F%2Fpow7.com%3A80%2Fannounce&tr=http%3A%2F%2Fns3107607.ip-54-36-126.eu%3A6969%2Fannounce&tr=http%3A%2F%2Fmail2.zelenaya.net%3A80%2Fannounce&tr=http%3A%2F%2Fbt.okmp3.ru%3A2710%2Fannounce&tr=http%3A%2F%2Fbt.100.pet%3A2710%2Fannounce&tr=http%3A%2F%2Fbobbialbano.com%3A6969%2Fannounce&tr=http%3A%2F%2Faaa.army%3A8866%2Fannounce&tr=https%3A%2F%2Ftr.bangumi.moe%3A9696%2Fannounce&tr=https%3A%2F%2Ftr.ready4.icu%3A443%2Fannounce&tr=http%3A%2F%2Ftracker.bt4g.com%3A2095%2Fannounce&tr=http%3A%2F%2Fderpyradio.net%3A6969%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:8cf3b5410f1c173b37eda5570bdcbd07af3b2222&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">816.2MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">-</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">-</span></td>
			<td nowrap="nowrap" align="center">-</td>
			<td align="center"><a href="/topics/list/user_id/661609">工藤新一</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 22:22			<span style="display: none;">2022/02/21 22:22</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/454" >
				风车字幕组</a></span>				
				<a href="/topics/view/593417_17_1080P_MP4.html"  target="_blank" >
				[风车字幕组十一周年][数码宝贝幽灵游戏][第17集][极寒地狱][1080P][简体][MP4]</a>
							</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:H7TKTGXHGQITMAOAZI4353EVY2IOU5S4&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Fopen.nyaatorrents.info%3A6544%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=http%3A%2F%2Ftracker.kisssub.org%2Fannounce&tr=http%3A%2F%2Ftracker.ex.ua%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.yify-torrents.com%2Fannounce&tr=http%3A%2F%2Fannounce.torrentsmd.com%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.leechers-paradise.org%3A6969%2Fannounce&tr=http%3A%2F%2Fretracker.telecom.kz%2Fannounce&tr=http%3A%2F%2Fbt.careland.com.cn%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.ktxp.com%3A6868%2Fannounce&tr=http%3A%2F%2Ftracker.dmhy.org%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.xfsub.com%3A6868%2Fannounce&tr=http%3A%2F%2Fmgtracker.org%3A2710%2Fannounce&tr=http%3A%2F%2Fmgtracker.org%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.grepler.com%3A6969%2Fannounce&tr=http%3A%2F%2Fbt.evrl.to%2Fannounce&tr=http%3A%2F%2Ftracker.mg64.net%3A6881%2Fannounce&tr=http%3A%2F%2Ftracker.torrentyorg.pl%2Fannounce&tr=http%3A%2F%2Ftracker.baravik.org%3A6970%2Fannounce&tr=http%3A%2F%2Ftracker.filetracker.pl%3A8089%2Fannounce&tr=http%3A%2F%2Ftracker1.wasabii.com.tw%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker2.wasabii.com.tw%3A6969%2Fannounce&tr=http%3A%2F%2Fretracker.local%2Fannounce&tr=http%3A%2F%2F37.19.5.139%3A6969%2Fannounce&tr=http%3A%2F%2F37.19.5.155%3A6881%2Fannounce&tr=http%3A%2F%2F46.4.109.148%3A6969%2Fannounce&tr=http%3A%2F%2F210.244.71.26%3A6969%2Fannounce&tr=http%3A%2F%2F210.244.71.25%3A6969%2Fannounce&tr=http%3A%2F%2F125.227.35.196%3A6969%2Fannounce&tr=http%3A%2F%2F213.159.215.198%3A6970%2Fannounce&tr=http%3A%2F%2Ftorrentsmd.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&tr=http%3A%2F%2Ftracker.internetwarriors.net%3A1337%2Fannounce&tr=http%3A%2F%2Fexplodie.org%3A6969%2Fannounce&tr=http%3A%2F%2Fopentracker.i2p.rocks%3A6969%2Fannounce&tr=http%3A%2F%2F5rt.tace.ru%3A60889%2Fannounce&tr=http%3A%2F%2Fopen.acgnxtracker.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.tiny-vps.com%3A6969%2Fannounce&tr=http%3A%2F%2Frt.tace.ru%3A80%2Fannounce&tr=https%3A%2F%2Ftrakx.herokuapp.com%3A443%2Fannounce&tr=http%3A%2F%2Fwww.loushao.net%3A8080%2Fannounce&tr=http%3A%2F%2Fh4.trakx.nibba.trade%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.zerobytes.xyz%3A1337%2Fannounce&tr=http%3A%2F%2Fvps02.net.orel.ru%3A80%2Fannounce&tr=https%3A%2F%2Fw.wwwww.wtf%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.tamersunion.org%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.sloppyta.co%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.renfei.net%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.nitrix.me%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.nanoha.org%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.lelux.fi%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.imgoingto.icu%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.hama3.net%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.foreverpirates.co%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.coalition.space%3A443%2Fannounce&tr=https%3A%2F%2Faaa.army%3A8866%2Fannounce&tr=https%3A%2F%2F1337.abcvg.info%3A443%2Fannounce&tr=http%3A%2F%2Fvpn.flying-datacenter.de%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker2.dler.org%3A80%2Fannounce&tr=http%3A%2F%2Ftracker1.bt.moack.co.kr%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.vraphim.com%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.sloppyta.co%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.skyts.net%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.noobsubs.net%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.loadbt.com%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.lelux.fi%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.gbitt.info%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.dler.org%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.anonwebz.xyz%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker-cdn.moeking.me%3A2095%2Fannounce&tr=http%3A%2F%2Ftorrenttracker.nwc.acsalaska.net%3A6969%2Fannounce&tr=http%3A%2F%2Ftorrentclub.online%3A54123%2Fannounce&tr=http%3A%2F%2Ft.overflow.biz%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%3A80%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce&tr=http%3A%2F%2Fretracker.sevstar.net%3A2710%2Fannounce&tr=http%3A%2F%2Fpow7.com%3A80%2Fannounce&tr=http%3A%2F%2Fns3107607.ip-54-36-126.eu%3A6969%2Fannounce&tr=http%3A%2F%2Fmail2.zelenaya.net%3A80%2Fannounce&tr=http%3A%2F%2Fbt.okmp3.ru%3A2710%2Fannounce&tr=http%3A%2F%2Fbt.100.pet%3A2710%2Fannounce&tr=http%3A%2F%2Fbobbialbano.com%3A6969%2Fannounce&tr=http%3A%2F%2Faaa.army%3A8866%2Fannounce&tr=https%3A%2F%2Ftr.bangumi.moe%3A9696%2Fannounce&tr=https%3A%2F%2Ftr.ready4.icu%3A443%2Fannounce&tr=http%3A%2F%2Ftracker.bt4g.com%3A2095%2Fannounce&tr=http%3A%2F%2Fderpyradio.net%3A6969%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:3fe6a99ae734113601c0ca39beec95c690ea765c&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">816.6MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">-</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">-</span></td>
			<td nowrap="nowrap" align="center">-</td>
			<td align="center"><a href="/topics/list/user_id/661609">工藤新一</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 21:53			<span style="display: none;">2022/02/21 21:53</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/666" >
				中肯字幕組</a></span>				
				<a href="/topics/view/593416_1_06_BIG5_MP4_1920X1080.html"  target="_blank" >
				【中肯字幕組】【1月新番】【川尻小玉的懒散生活】【06】【BIG5_MP4】【1920X1080】</a>
				<span style="color: gray;">約1條評論</span>			</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:UMVRBJGR7GYRKPCS55KBK4DP3Q7LDZPB&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:a32b10a4d1f9b1153c52ef5415706fdc3eb1e5e1&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">24.7MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">-</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">-</span></td>
			<td nowrap="nowrap" align="center">-</td>
			<td align="center"><a href="/topics/list/user_id/665734">yellow9395</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 21:49			<span style="display: none;">2022/02/21 21:49</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/185" >
				极影字幕社</a></span>				
				<a href="/topics/view/593415_10_Shingeki_no_Kyojin_The_Final_Season_23_BIG5_MP4_720.html"  target="_blank" >
				【極影字幕社】 ★10月新番 【進擊的巨人 最終季】【Shingeki no Kyojin The Final Season】【23】BIG5 MP4_720</a>
							</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:VV6AROJKEBW6JV3TPWNLFCJP66LDUUOH&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.ktxp.com%3A6868%2Fannounce&tr=http%3A%2F%2Ftracker.ktxp.com%3A7070%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%2Fannounce&tr=http%3A%2F%2Ftracker.dmhy.org%3A8000%2Fannounce&tr=http%3A%2F%2Fbt.popgo.net%3A7456%2Fannounce&tr=http%3A%2F%2Fnyaatorrents.info%3A7266%2Fannounce&tr=http%3A%2F%2Fnyaatorrents.info%3A3277%2Fannounce&tr=http%3A%2F%2Fbt.sc-ol.com%3A2710%2Fannounce&tr=http%3A%2F%2Fanisource.spb.ru%2Fannounce&tr=http%3A%2F%2Ftracker.anirena.com%3A81%2Fannounce&tr=http%3A%2F%2Ftracker.anirena.com%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.frozen-layer.net%3A6969%2Fannounce&tr=http%3A%2F%2Fmkfs.ru%2Fannounce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker.levelup.cn%3A8080%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fbt.wiiyi.com%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.dmguo.com%3A2710%2Fannounce&tr=http%3A%2F%2Fbt1.the9.com%3A6969%2Fannounce&tr=http%3A%2F%2Ftorrent-download.to%3A5869%2Fannounce&tr=http%3A%2F%2Ftorrent-download.to%3A5869%2Fannounce.php&tr=http%3A%2F%2Ftorrent-downloads.to%3A5869%2Fannounce&tr=http%3A%2F%2Ftorrent-downloads.to%3A5869%2Fannounce.php&tr=http%3A%2F%2Ftracker.bittorrent.am%2Fannounce&tr=http%3A%2F%2Ftracker.torrent.to%3A2710%2Fannounce&tr=http%3A%2F%2Ftv.tracker.prq.to%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce.php&tr=http%3A%2F%2Ftracker9.bol.bg%2Fannounce.php&tr=http%3A%2F%2Fdenis.stalker.h3q.com%2Fannounce&tr=http%3A%2F%2Fdenis.stalker.h3q.com%2Fannounce.php&tr=http%3A%2F%2Fdenis.stalker.h3q.com%3A6969%2Fannounce&tr=http%3A%2F%2Fdenis.stalker.h3q.com%3A6969%2Fannounce.php&tr=http%3A%2F%2Fbttrack.9you.com%3A8080%2Fannounce&tr=http%3A%2F%2F9.rarbg.com%3A2710%2Fannounce&tr=http%3A%2F%2Fgenesis.1337x.org%3A1337%2Fannounce&tr=http%3A%2F%2Ftpb.tracker.prq.to%2Fannounce&tr=http%3A%2F%2Fbt.edwardk.info%3A6969%2Fannounce&tr=http%3A%2F%2Fbt.sumisora.com%3A2710%2Fannounce&tr=http%3A%2F%2Fshare.hkg-fansub.info%2Fannounce.php">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:ad7c08b92a206de4d7737d9ab2892ff7963a51c7&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">181.7MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">-</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">-</span></td>
			<td nowrap="nowrap" align="center">-</td>
			<td align="center"><a href="/topics/list/user_id/403163">rossina</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 21:32			<span style="display: none;">2022/02/21 21:32</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/801" >
				NC-Raws</a></span>				
				<a href="/topics/view/593414_NC-Raws_TRIBE_NINE_-_07_B-Global_1920x1080_HEVC_AAC_MKV.html"  target="_blank" >
				[NC-Raws] 夜街酷斗 / TRIBE NINE - 07 (B-Global 1920x1080 HEVC AAC MKV)</a>
							</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:KH65TZM4QGCYKGPAN43565ZCAKU6L27T&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Fopen.nyaatorrents.info%3A6544%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=https%3A%2F%2Ftr.bangumi.moe%3A9696%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:51fdd9e59c81858519e06f37df772202a9e5ebf3&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">641.3MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">-</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">-</span></td>
			<td nowrap="nowrap" align="center">-</td>
			<td align="center"><a href="/topics/list/user_id/637871">九十九朔夜</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 21:31			<span style="display: none;">2022/02/21 21:31</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/801" >
				NC-Raws</a></span>				
				<a href="/topics/view/593413_NC-Raws_TRIBE_NINE_-_07_Baha_1920x1080_AVC_AAC_MP4.html"  target="_blank" >
				[NC-Raws] TRIBE NINE - 07 (Baha 1920x1080 AVC AAC MP4)</a>
				<span style="color: gray;">約1條評論</span>			</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:CNXJG2WME7UAUIAOJOSFG2TE37RFOFGL&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Fopen.nyaatorrents.info%3A6544%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=https%3A%2F%2Ftr.bangumi.moe%3A9696%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:136e936acc27e80a200e4ba4536a64dfe25714cb&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">301.1MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">-</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">-</span></td>
			<td nowrap="nowrap" align="center">-</td>
			<td align="center"><a href="/topics/list/user_id/637871">九十九朔夜</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 21:29			<span style="display: none;">2022/02/21 21:29</span></td>
			<td width="6%" align="center">
			<a class="sort-6" 
				href="/topics/list/sort_id/6">
				<font color=blue>日劇</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/648" >
				魔星字幕团</a></span>				
				<a href="/topics/view/593412_MagicStar_EP07_WEBDL_1080p.html"  target="_blank" >
				[MagicStar] 勿言推理 / ミステリという勿れ EP07 [WEBDL] [1080p]【生】【附日字】</a>
							</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:FSTASNTY4J5KQ7HLFU7NTEDHSSLKTJCO&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:2ca6093678e27aa87ceb2d3ed990679496a9a44e&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">951.8MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">-</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">-</span></td>
			<td nowrap="nowrap" align="center">-</td>
			<td align="center"><a href="/topics/list/user_id/105217">jackieliiy</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 20:58			<span style="display: none;">2022/02/21 20:58</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/805" >
				DBD制作组</a></span>				
				<a href="/topics/view/593411_DBD-Raws_Sakugan_01-12TV_1080P_BDRip_HEVC-10bit_FLAC_MKV.html"  target="_blank" >
				[DBD-Raws][迷宫标记者/Sakugan/サクガン][01-12TV全集][1080P][BDRip][HEVC-10bit][简繁日外挂][FLAC][MKV]</a>
				<span style="color: gray;">約1條評論</span>			</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:V3KRY35SMRJIOV77PNFKMMEQLBAHTH2K&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2F1337.abcvg.info%2Fannounce&tr=http%3A%2F%2F178.175.143.27%2Fannounce&tr=http%3A%2F%2F5rt.tace.ru%3A60889%2Fannounce&tr=http%3A%2F%2F78.30.254.12%3A2710%2Fannounce&tr=http%3A%2F%2F91.217.91.21%3A3218%2Fannounce&tr=http%3A%2F%2F%5B2001%3A1b10%3A1000%3A8101%3A0%3A242%3Aac11%3A2%5D%3A6969%2Fannounce&tr=http%3A%2F%2F%5B2001%3A470%3A1%3A189%3A0%3A1%3A2%3A3%5D%3A6969%2Fannounce&tr=http%3A%2F%2F%5B2a04%3Aac00%3A1%3A3dd8%3A%3A1%3A2710%5D%3A2710%2Fannounce&tr=http%3A%2F%2Faaa.army%3A8866%2Fannounce&tr=http%3A%2F%2Fatrack.pow7.com%2Fannounce&tr=http%3A%2F%2Fbobbialbano.com%3A6969%2Fannounce&tr=http%3A%2F%2Fbt.3kb.xyz%2Fannounce&tr=http%3A%2F%2Fbt.pusacg.org%3A8080%2Fannounce&tr=http%3A%2F%2Fderpyradio.net%3A6969%2Fannounce&tr=http%3A%2F%2Fexplodie.org%3A6969%2Fannounce&tr=http%3A%2F%2Fgrifon.info%2Fannounce&tr=http%3A%2F%2Fh4.trakx.nibba.trade%2Fannounce&tr=http%3A%2F%2Fjsb.by%3A8000%2Fannounce&tr=http%3A%2F%2Fnovaopcj.icu%3A10325%2Fannounce&tr=http%3A%2F%2Fns349743.ip-91-121-106.eu%2Fannounce&tr=http%3A%2F%2Fopen.acgnxtracker.com%2Fannounce&tr=http%3A%2F%2Fopen.touki.ru%2Fannounce.php&tr=http%3A%2F%2Fopentracker.i2p.rocks%3A6969%2Fannounce&tr=http%3A%2F%2Fopentracker.xyz%2Fannounce&tr=http%3A%2F%2Fpow7.com%2Fannounce&tr=http%3A%2F%2Fpublictracker.ch%3A6969%2Fannounce&tr=http%3A%2F%2Fretracker.hotplug.ru%3A2710%2Fannounce&tr=http%3A%2F%2Fretracker.sevstar.net%3A2710%2Fannounce&tr=http%3A%2F%2Frt.tace.ru%2Fannounce&tr=http%3A%2F%2Fsecure.pow7.com%2Fannounce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Ft.overflow.biz%3A6969%2Fannounce&tr=http%3A%2F%2Ft1.pow7.com%2Fannounce&tr=http%3A%2F%2Ft2.pow7.com%2Fannounce&tr=http%3A%2F%2Fthetracker.org%2Fannounce&tr=http%3A%2F%2Ftorrentsmd.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftorrenttracker.nwc.acsalaska.net%3A6969%2Fannounce&tr=http%3A%2F%2Ftr.cili001.com%3A8070%2Fannounce&tr=http%3A%2F%2Ftracker.anonwebz.xyz%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker.birkenwald.de%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.bittor.pw%3A1337%2Fannounce&tr=http%3A%2F%2Ftracker.bittorrent.am%2Fannounce&tr=http%3A%2F%2Ftracker.bt4g.com%3A2095%2Fannounce&tr=http%3A%2F%2Ftracker.bz%2Fannounce&tr=http%3A%2F%2Ftracker.dler.org%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.dutchtracking.nl%2Fannounce&tr=http%3A%2F%2Ftracker.files.fm%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.gbitt.info%2Fannounce&tr=http%3A%2F%2Ftracker.internetwarriors.net%3A1337%2Fannounce&tr=http%3A%2F%2Ftracker.ipv6tracker.ru%2Fannounce&tr=http%3A%2F%2Ftracker.kamigami.org%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.kuroy.me%3A5944%2Fannounce&tr=http%3A%2F%2Ftracker.lelux.fi%2Fannounce&tr=http%3A%2F%2Ftracker.moeking.me%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.noobsubs.net%2Fannounce&tr=http%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&tr=http%3A%2F%2Ftracker.sakurato.xyz%3A23333%2Fannounce&tr=http%3A%2F%2Ftracker.skyts.net%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.sloppyta.co%2Fannounce&tr=http%3A%2F%2Ftracker.trackerfix.com%2Fannounce&tr=http%3A%2F%2Ftracker.uw0.xyz%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.ygsub.com%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.zerobytes.xyz%3A1337%2Fannounce&tr=http%3A%2F%2Ftracker2.dler.org%2Fannounce&tr=http%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftrun.tom.ru%2Fannounce&tr=http%3A%2F%2Fvpn.flying-datacenter.de%3A6969%2Fannounce&tr=http%3A%2F%2Fvps02.net.orel.ru%2Fannounce&tr=http%3A%2F%2Fwww.wareztorrent.com%2Fannounce&tr=https%3A%2F%2F1337.abcvg.info%2Fannounce&tr=https%3A%2F%2F2.tracker.eu.org%2Fannounce&tr=https%3A%2F%2F3.tracker.eu.org%2Fannounce&tr=https%3A%2F%2Faaa.army%3A8866%2Fannounce&tr=https%3A%2F%2Fopen.kickasstracker.com%2Fannounce&tr=https%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=https%3A%2F%2Ftr.ready4.icu%2Fannounce&tr=https%3A%2F%2Ftracker.bt-hash.com%2Fannounce&tr=https%3A%2F%2Ftracker.gbitt.info%2Fannounce&tr=https%3A%2F%2Ftracker.hama3.net%2Fannounce&tr=https%3A%2F%2Ftracker.imgoingto.icu%2Fannounce&tr=https%3A%2F%2Ftracker.lelux.fi%2Fannounce&tr=https%3A%2F%2Ftracker.lilithraws.cf%2Fannounce&tr=https%3A%2F%2Ftracker.nitrix.me%2Fannounce&tr=https%3A%2F%2Ftracker.parrotsec.org%2Fannounce&tr=https%3A%2F%2Ftracker.sloppyta.co%2Fannounce&tr=https%3A%2F%2Ftracker.tamersunion.org%2Fannounce&tr=https%3A%2F%2Ftrakx.herokuapp.com%2Fannounce&tr=https%3A%2F%2Fw.wwwww.wtf%2Fannounce&tr=https%3A%2F%2Fwww.wareztorrent.com%2Fannounce&tr=http%3A%2F%2F1337.abcvg.info%2Fannounce%2Chttp%3A%2F%2F178.175.143.27%3A80%2Fannounce%2Chttp%3A%2F%2F5rt.tace.ru%3A60889%2Fannounce%2Chttp%3A%2F%2F78.30.254.12%3A2710%2Fannounce%2Chttp%3A%2F%2F91.217.91.21%3A3218%2Fannounce%2Chttp%3A%2F%2F%5B2001%3A1b10%3A1000%3A8101%3A0%3A242%3Aac11%3A2%5D%3A6969%2Fannounce%2Chttp%3A%2F%2F%5B2001%3A470%3A1%3A189%3A0%3A1%3A2%3A3%5D%3A6969%2Fannounce%2Chttp%3A%2F%2F%5B2a04%3Aac00%3A1%3A3dd8%3A%3A1%3A2710%5D%3A2710%2Fannounce%2Chttp%3A%2F%2Faaa.army%3A8866%2Fannounce%2Chttp%3A%2F%2Fatrack.pow7.com%3A80%2Fannounce%2Chttp%3A%2F%2Fbobbialbano.com%3A6969%2Fannounce%2Chttp%3A%2F%2Fbt.3kb.xyz%3A80%2Fannounce%2Chttp%3A%2F%2Fbt.pusacg.org%3A8080%2Fannounce%2Chttp%3A%2F%2Fderpyradio.net%3A6969%2Fannounce%2Chttp%3A%2F%2Fexplodie.org%3A6969%2Fannounce%2Chttp%3A%2F%2Fgrifon.info%3A80%2Fannounce%2Chttp%3A%2F%2Fh4.trakx.nibba.trade%3A80%2Fannounce%2Chttp%3A%2F%2Fjsb.by%3A8000%2Fannounce%2Chttp%3A%2F%2Fnovaopcj.icu%3A10325%2Fannounce%2Chttp%3A%2F%2Fns349743.ip-91-121-106.eu%3A80%2Fannounce%2Chttp%3A%2F%2Fopen.acgnxtracker.com%3A80%2Fannounce%2Chttp%3A%2F%2Fopen.touki.ru%3A80%2Fannounce.php%2Chttp%3A%2F%2Fopentracker.i2p.rocks%3A6969%2Fannounce%2Chttp%3A%2F%2Fopentracker.xyz%3A80%2Fannounce%2Chttp%3A%2F%2Fpow7.com%3A80%2Fannounce%2Chttp%3A%2F%2Fpublictracker.ch%3A6969%2Fannounce%2Chttp%3A%2F%2Fretracker.hotplug.ru%3A2710%2Fannounce%2Chttp%3A%2F%2Fretracker.sevstar.net%3A2710%2Fannounce%2Chttp%3A%2F%2Frt.tace.ru%3A80%2Fannounce%2Chttp%3A%2F%2Fsecure.pow7.com%3A80%2Fannounce%2Chttp%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce%2Chttp%3A%2F%2Ft.acg.rip%3A6699%2Fannounce%2Chttp%3A%2F%2Ft.nyaatracker.com%3A80%2Fannounce%2Chttp%3A%2F%2Ft.overflow.biz%3A6969%2Fannounce%2Chttp%3A%2F%2Ft1.pow7.com%3A80%2Fannounce%2Chttp%3A%2F%2Ft2.pow7.com%3A80%2Fannounce%2Chttp%3A%2F%2Fthetracker.org%3A80%2Fannounce%2Chttp%3A%2F%2Ftorrentsmd.com%3A8080%2Fannounce%2Chttp%3A%2F%2Ftorrenttracker.nwc.acsalaska.net%3A6969%2Fannounce%2Chttp%3A%2F%2Ftr.cili001.com%3A8070%2Fannounce%2Chttp%3A%2F%2Ftracker.anonwebz.xyz%3A8080%2Fannounce%2Chttp%3A%2F%2Ftracker.birkenwald.de%3A6969%2Fannounce%2Chttp%3A%2F%2Ftracker.bittor.pw%3A1337%2Fannounce%2Chttp%3A%2F%2Ftracker.bittorrent.am%3A80%2Fannounce%2Chttp%3A%2F%2Ftracker.bt4g.com%3A2095%2Fannounce%2Chttp%3A%2F%2Ftracker.bz%3A80%2Fannounce%2Chttp%3A%2F%2Ftracker.dler.org%3A6969%2Fannounce%2Chttp%3A%2F%2Ftracker.dutchtracking.nl%3A80%2Fannounce%2Chttp%3A%2F%2Ftracker.files.fm%3A6969%2Fannounce%2Chttp%3A%2F%2Ftracker.gbitt.info%3A80%2Fannounce%2Chttp%3A%2F%2Ftracker.internetwarriors.net%3A1337%2Fannounce%2Chttp%3A%2F%2Ftracker.ipv6tracker.ru%3A80%2Fannounce%2Chttp%3A%2F%2Ftracker.kamigami.org%3A2710%2Fannounce%2Chttp%3A%2F%2Ftracker.kuroy.me%3A5944%2Fannounce%2Chttp%3A%2F%2Ftracker.lelux.fi%3A80%2Fannounce%2Chttp%3A%2F%2Ftracker.moeking.me%3A6969%2Fannounce%2Chttp%3A%2F%2Ftracker.noobsubs.net%3A80%2Fannounce%2Chttp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce%2Chttp%3A%2F%2Ftracker.sakurato.xyz%3A23333%2Fannounce%2Chttp%3A%2F%2Ftracker.skyts.net%3A6969%2Fannounce%2Chttp%3A%2F%2Ftracker.sloppyta.co%3A80%2Fannounce%2Chttp%3A%2F%2Ftracker.trackerfix.com%3A80%2Fannounce%2Chttp%3A%2F%2Ftracker.uw0.xyz%3A6969%2Fannounce%2Chttp%3A%2F%2Ftracker.ygsub.com%3A6969%2Fannounce%2Chttp%3A%2F%2Ftracker.zerobytes.xyz%3A1337%2Fannounce%2Chttp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce%2Chttp%3A%2F%2Ftracker2.dler.org%3A80%2Fannounce%2Chttp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce%2Chttp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce%2Chttp%3A%2F%2Ftrun.tom.ru%3A80%2Fannounce%2Chttp%3A%2F%2Fvpn.flying-datacenter.de%3A6969%2Fannounce%2Chttp%3A%2F%2Fvps02.net.orel.ru%3A80%2Fannounce%2Chttp%3A%2F%2Fwww.wareztorrent.com%3A80%2Fannounce%2Chttps%3A%2F%2F1337.abcvg.info%3A443%2Fannounce%2Chttps%3A%2F%2F2.tracker.eu.org%3A443%2Fannounce%2Chttps%3A%2F%2F3.tracker.eu.org%3A443%2Fannounce%2Chttps%3A%2F%2Faaa.army%3A8866%2Fannounce%2Chttps%3A%2F%2Fopen.kickasstracker.com%3A443%2Fannounce%2Chttps%3A%2F%2Fopentracker.acgnx.se%3A443%2Fannounce%2Chttps%3A%2F%2Ftr.ready4.icu%3A443%2Fannounce%2Chttps%3A%2F%2Ftracker.bt-hash.com%3A443%2Fannounce%2Chttps%3A%2F%2Ftracker.gbitt.info%3A443%2Fannounce%2Chttps%3A%2F%2Ftracker.hama3.net%3A443%2Fannounce%2Chttps%3A%2F%2Ftracker.imgoingto.icu%3A443%2Fannounce%2Chttps%3A%2F%2Ftracker.lelux.fi%3A443%2Fannounce%2Chttps%3A%2F%2Ftracker.lilithraws.cf%3A443%2Fannounce%2Chttps%3A%2F%2Ftracker.nitrix.me%3A443%2Fannounce%2Chttps%3A%2F%2Ftracker.parrotsec.org%3A443%2Fannounce%2Chttps%3A%2F%2Ftracker.sloppyta.co%3A443%2Fannounce%2Chttps%3A%2F%2Ftracker.tamersunion.org%3A443%2Fannounce%2Chttps%3A%2F%2Ftrakx.herokuapp.com%3A443%2Fannounce%2Chttps%3A%2F%2Fw.wwwww.wtf%3A443%2Fannounce%2Chttps%3A%2F%2Fwww.wareztorrent.com%3A443%2Fannounce%2Cudp%3A%2F%2F104.238.198.186%3A8000%2Fannounce%2Cudp%3A%2F%2F107.150.14.110%3A6969%2Fannounce%2Cudp%3A%2F%2F109.121.134.121%3A1337%2Fannounce%2Cudp%3A%2F%2F114.55.113.60%3A6969%2Fannounce%2Cudp%3A%2F%2F128.199.70.66%3A5944%2Fannounce%2Cudp%3A%2F%2F151.80.120.114%3A2710%2Fannounce%2Cudp%3A%2F%2F168.235.67.63%3A6969%2Fannounce%2Cudp%3A%2F%2F178.33.73.26%3A2710%2Fannounce%2Cudp%3A%2F%2F182.176.139.129%3A6969%2Fannounce%2Cudp%3A%2F%2F185.5.97.139%3A8089%2Fannounce%2Cudp%3A%2F%2F185.83.215.123%3A6969%2Fannounce%2Cudp%3A%2F%2F185.86.149.205%3A1337%2Fannounce%2Cudp%3A%2F%2F188.165.253.109%3A1337%2Fannounce%2Cudp%3A%2F%2F191.101.229.236%3A1337%2Fannounce%2Cudp%3A%2F%2F194.106.216.222%3A80%2Fannounce%2Cudp%3A%2F%2F195.123.209.37%3A1337%2Fannounce%2Cudp%3A%2F%2F195.123.209.40%3A80%2Fannounce%2Cudp%3A%2F%2F208.67.16.113%3A8000%2Fannounce%2Cudp%3A%2F%2F212.1.226.176%3A2710%2Fannounce%2Cudp%3A%2F%2F212.47.227.58%3A6969%2Fannounce%2Cudp%3A%2F%2F213.163.67.56%3A1337%2Fannounce%2Cudp%3A%2F%2F37.19.5.155%3A2710%2Fannounce%2Cudp%3A%2F%2F3rt.tace.ru%3A60889%2Fannounce%2Cudp%3A%2F%2F46.4.109.148%3A6969%2Fannounce%2Cudp%3A%2F%2F47.ip-51-68-199.eu%3A6969%2Fannounce%2Cudp%3A%2F%2F5.79.249.77%3A6969%2Fannounce%2Cudp%3A%2F%2F5.79.83.193%3A6969%2Fannounce%2Cudp%3A%2F%2F51.254.244.161%3A6969%2Fannounce%2Cudp%3A%2F%2F52.58.128.163%3A6969%2Fannounce%2Cudp%3A%2F%2F61626c.net%3A6969%2Fannounce%2Cudp%3A%2F%2F62.138.0.158%3A6969%2Fannounce%2Cudp%3A%2F%2F62.212.85.66%3A2710%2Fannounce%2Cudp%3A%2F%2F74.82.52.209%3A6969%2Fannounce%2Cudp%3A%2F%2F78.30.254.12%3A2710%2Fannounce%2Cudp%3A%2F%2F85.17.19.180%3A80%2Fannounce%2Cudp%3A%2F%2F89.234.156.205%3A80%2Fannounce%2Cudp%3A%2F%2F9.rarbg.com%3A2710%2Fannounce%2Cudp%3A%2F%2F9.rarbg.me%3A2710%2Fannounce%2Cudp%3A%2F%2F9.rarbg.me%3A2780%2Fannounce%2Cudp%3A%2F%2F9.rarbg.to%3A2710%2Fannounce%2Cudp%3A%2F%2F9.rarbg.to%3A2730%2Fannounce%2Cudp%3A%2F%2F91.216.110.52%3A451%2Fannounce%2Cudp%3A%2F%2F91.218.230.81%3A6969%2Fannounce%2Cudp%3A%2F%2F94.23.183.33%3A6969%2Fannounce%2Cudp%3A%2F%2F95.211.168.204%3A2710%2Fannounce%2Cudp%3A%2F%2F%5B2001%3A1b10%3A1000%3A8101%3A0%3A242%3Aac11%3A2%5D%3A6969%2Fannounce%2Cudp%3A%2F%2F%5B2001%3A470%3A1%3A189%3A0%3A1%3A2%3A3%5D%3A6969%2Fannounce%2Cudp%3A%2F%2F%5B2a03%3A7220%3A8083%3Acd00%3A%3A1%5D%3A451%2Fannounce%2Cudp%3A%2F%2F%5B2a04%3Aac00%3A1%3A3dd8%3A%3A1%3A2710%5D%3A2710%2Fannounce%2Cudp%3A%2F%2F%5B2a04%3Ac44%3Ae00%3A32e0%3A4cf%3A6aff%3Afe00%3Aaa1%5D%3A6969%2Fannounce%2Cudp%3A%2F%2Faaa.army%3A8866%2Fannounce%2Cudp%3A%2F%2Fadm.category5.tv%3A6969%2Fannounce%2Cudp%3A%2F%2Fadmin.videoenpoche.info%3A6969%2Fannounce%2Cudp%3A%2F%2Fadminion.n-blade.ru%3A6969%2Fannounce%2Cudp%3A%2F%2Fanidex.moe%3A6969%2Fannounce%2Cudp%3A%2F%2Fapi.bitumconference.ru%3A6969%2Fannounce%2Cudp%3A%2F%2Faruacfilmes.com.br%3A6969%2Fannounce%2Cudp%3A%2F%2Fasger.xyz%3A6969%2Fannounce%2Cudp%3A%2F%2Fbclearning.top%3A6969%2Fannounce%2Cudp%3A%2F%2Fbenouworldtrip.fr%3A6969%2Fannounce%2Cudp%3A%2F%2Fbioquantum.co.za%3A6969%2Fannounce%2Cudp%3A%2F%2Fbitsparadise.info%3A6969%2Fannounce%2Cudp%3A%2F%2Fblokas.io%3A6969%2Fannounce%2Cudp%3A%2F%2Fbms-hosxp.com%3A6969%2Fannounce%2Cudp%3A%2F%2Fbt.firebit.org%3A2710%2Fannounce%2Cudp%3A%2F%2Fbt.okmp3.ru%3A2710%2Fannounce%2Cudp%3A%2F%2Fbt1.archive.org%3A6969%2Fannounce%2Cudp%3A%2F%2Fbt2.54new.com%3A8080%2Fannounce%2Cudp%3A%2F%2Fbt2.archive.org%3A6969%2Fannounce%2Cudp%3A%2F%2Fbubu.mapfactor.com%3A6969%2Fannounce%2Cudp%3A%2F%2Fcamera.lei001.com%3A6969%2Fannounce%2Cudp%3A%2F%2Fcdn-1.gamecoast.org%3A6969%2Fannounce%2Cudp%3A%2F%2Fcdn-2.gamecoast.org%3A6969%2Fannounce%2Cudp%3A%2F%2Fchanchan.uchuu.co.uk%3A6969%2Fannounce%2Cudp%3A%2F%2Fchihaya.toss.li%3A9696%2Fannounce%2Cudp%3A%2F%2Fcode2chicken.nl%3A6969%2Fannounce%2Cudp%3A%2F%2Fconcen.org%3A6969%2Fannounce%2Cudp%3A%2F%2Fcutiegirl.ru%3A6969%2Fannounce%2Cudp%3A%2F%2Fdaveking.com%3A6969%2Fannounce%2Cudp%3A%2F%2Fdenis.stalker.upeer.me%3A6969%2Fannounce%2Cudp%3A%2F%2Fdiscord.heihachi.pw%3A6969%2Fannounce%2Cudp%3A%2F%2Fdpiui.reedlan.com%3A6969%2Fannounce%2Cudp%3A%2F%2Fdrumkitx.com%3A6969%2Fannounce%2Cudp%3A%2F%2Feddie4.nl%3A6969%2Fannounce%2Cudp%3A%2F%2Fedu.uifr.ru%3A6969%2Fannounce%2Cudp%3A%2F%2Fengplus.ru%3A6969%2Fannounce%2Cudp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce%2Cudp%3A%2F%2Ffe.dealclub.de%3A6969%2Fannounce%2Cudp%3A%2F%2Fforever-tracker.zooki.xyz%3A6969%2Fannounce%2Cudp%3A%2F%2Ffree-tracker.zooki.xyz%3A6969%2Fannounce%2Cudp%3A%2F%2Fgit.vulnix.sh%3A6969%2Fannounce%2Cudp%3A%2F%2Finferno.demonoid.is%3A3391%2Fannounce%2Cudp%3A%2F%2Fipv4.tracker.harry.lu%3A80%2Fannounce%2Cudp%3A%2F%2Fipv6.tracker.zerobytes.xyz%3A16661%2Fannounce%2Cudp%3A%2F%2Fjsb.by%3A8000%2Fannounce%2Cudp%3A%2F%2Fkanal-4.de%3A6969%2Fannounce%2Cudp%3A%2F%2Fkoli.services%3A6969%2Fannounce%2Cudp%3A%2F%2Fline-net.ru%3A6969%2Fannounce%2Cudp%3A%2F%2Fln.mtahost.co%3A6969%2Fannounce%2Cudp%3A%2F%2Fmail.realliferpg.de%3A6969%2Fannounce%2Cudp%3A%2F%2Fmgtracker.org%3A2710%2Fannounce%2Cudp%3A%2F%2Fmovies.zsw.ca%3A6969%2Fannounce%2Cudp%3A%2F%2Fmts.tvbit.co%3A6969%2Fannounce%2Cudp%3A%2F%2Fnagios.tks.sumy.ua%3A80%2Fannounce%2Cudp%3A%2F%2Fns-1.x-fins.com%3A6969%2Fannounce%2Cudp%3A%2F%2Fns389251.ovh.net%3A6969%2Fannounce%2Cudp%3A%2F%2Fopen.demonii.com%3A1337%2Fannounce%2Cudp%3A%2F%2Fopen.lolicon.eu%3A7777%2Fannounce%2Cudp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce%2Cudp%3A%2F%2Fopentor.org%3A2710%2Fannounce%2Cudp%3A%2F%2Fopentracker.i2p.rocks%3A6969%2Fannounce%2Cudp%3A%2F%2Fpeerfect.org%3A6969%2Fannounce%2Cudp%3A%2F%2Fpublic-tracker.zooki.xyz%3A6969%2Fannounce%2Cudp%3A%2F%2Fpublic.popcorn-tracker.org%3A6969%2Fannounce%2Cudp%3A%2F%2Fpublic.publictracker.xyz%3A6969%2Fannounce%2Cudp%3A%2F%2Fpublictracker.xyz%3A6969%2Fannounce%2Cudp%3A%2F%2Fretracker.akado-ural.ru%3A80%2Fannounce%2Cudp%3A%2F%2Fretracker.hotplug.ru%3A2710%2Fannounce%2Cudp%3A%2F%2Fretracker.lanta-net.ru%3A2710%2Fannounce%2Cudp%3A%2F%2Fretracker.netbynet.ru%3A2710%2Fannounce%2Cudp%3A%2F%2Fretracker.nts.su%3A2710%2Fannounce%2Cudp%3A%2F%2Fretracker.sevstar.net%3A2710%2Fannounce%2Cudp%3A%2F%2Frutorrent.frontline-mod.com%3A6969%2Fannounce%2Cudp%3A%2F%2Fsd-161673.dedibox.fr%3A6969%2Fannounce%2Cudp%3A%2F%2Fshadowshq.eddie4.nl%3A6969%2Fannounce%2Cudp%3A%2F%2Fshadowshq.yi.org%3A6969%2Fannounce%2Cudp%3A%2F%2Fstorage.groupees.com%3A6969%2Fannounce%2Cudp%3A%2F%2Ft1.leech.ie%3A1337%2Fannounce%2Cudp%3A%2F%2Ft2.leech.ie%3A1337%2Fannounce%2Cudp%3A%2F%2Ft3.leech.ie%3A1337%2Fannounce%2Cudp%3A%2F%2Fteamspeak.value-wolf.org%3A6969%2Fannounce%2Cudp%3A%2F%2Fthetracker.org%3A80%2Fannounce%2Cudp%3A%2F%2Ftorrent.tdjs.tech%3A6969%2Fannounce%2Cudp%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce%2Cudp%3A%2F%2Ftr.cili001.com%3A8070%2Fannounce%2Cudp%3A%2F%2Ftr2.ysagin.top%3A2710%2Fannounce%2Cudp%3A%2F%2Ftracker-udp.gbitt.info%3A80%2Fannounce%2Cudp%3A%2F%2Ftracker-v6.zooki.xyz%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.0x.tf%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.aletorrenty.pl%3A2710%2Fannounce%2Cudp%3A%2F%2Ftracker.altrosky.nl%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.army%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.beeimg.com%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.birkenwald.de%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.bittor.pw%3A1337%2Fannounce%2Cudp%3A%2F%2Ftracker.blacksparrowmedia.net%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.coppersurfer.tk%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.cyberia.is%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.dler.org%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.ds.is%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.dyne.org%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.e-utp.net%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.eddie4.nl%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.ex.ua%3A80%2Fannounce%2Cudp%3A%2F%2Ftracker.filemail.com%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.flashtorrents.org%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.fortu.io%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.grepler.com%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.internetwarriors.net%3A1337%2Fannounce%2Cudp%3A%2F%2Ftracker.kali.org%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.kamigami.org%3A2710%2Fannounce%2Cudp%3A%2F%2Ftracker.kuroy.me%3A5944%2Fannounce%2Cudp%3A%2F%2Ftracker.leechers-paradise.org%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.lelux.fi%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.moeking.me%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.open-internet.nl%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce%2Cudp%3A%2F%2Ftracker.piratepublic.com%3A1337%2Fannounce%2Cudp%3A%2F%2Ftracker.publictracker.xyz%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.shkinev.me%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.skynetcloud.site%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.skyts.net%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.swateam.org.uk%3A2710%2Fannounce%2Cudp%3A%2F%2Ftracker.tiny-vps.com%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce%2Cudp%3A%2F%2Ftracker.tvunderground.org.ru%3A3218%2Fannounce%2Cudp%3A%2F%2Ftracker.uw0.xyz%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.v6speed.org%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.vulnix.sh%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.yoshi210.com%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.zemoj.com%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.zerobytes.xyz%3A1337%2Fannounce%2Cudp%3A%2F%2Ftracker.zum.bi%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker0.ufibox.com%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce%2Cudp%3A%2F%2Ftracker2.christianbro.pw%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker2.dler.org%3A80%2Fannounce%2Cudp%3A%2F%2Ftracker2.indowebster.com%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce%2Cudp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce%2Cudp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce%2Cudp%3A%2F%2Ftracker6.dler.org%3A2710%2Fannounce%2Cudp%3A%2F%2Ftsundere.pw%3A6969%2Fannounce%2Cudp%3A%2F%2Fu.wwwww.wtf%3A1%2Fannounce%2Cudp%3A%2F%2Fultra.zt.ua%3A6969%2Fannounce%2Cudp%3A%2F%2Fus-tracker.publictracker.xyz%3A6969%2Fannounce%2Cudp%3A%2F%2Fvalakas.rollo.dnsabr.com%3A2710%2Fannounce%2Cudp%3A%2F%2Fvibe.community%3A6969%2Fannounce%2Cudp%3A%2F%2Fwassermann.online%3A6969%2Fannounce%2Cudp%3A%2F%2Fwww.loushao.net%3A8080%2Fannounce%2Cudp%3A%2F%2Fwww.midea123.z-media.com.cn%3A6969%2Fannounce%2Cudp%3A%2F%2Fwww.torrent.eu.org%3A451%2Fannounce%2Cudp%3A%2F%2Fzephir.monocul.us%3A6969%2Fannounce%2Cudp%3A%2F%2Fzer0day.ch%3A1337%2Fannounce%2Cudp%3A%2F%2Fzer0day.to%3A1337%2Fannounce&tr=http%3A%2F%2F104.238.198.186%3A8000%2Fannounce&tr=http%3A%2F%2F185.83.215.123%3A6969%2Fannounce&tr=http%3A%2F%2F62.210.202.61%2Fannounce&tr=http%3A%2F%2F87.110.238.140%3A6969%2Fannounce&tr=http%3A%2F%2F95.211.168.204%3A2710%2Fannounce&tr=http%3A%2F%2Fmail2.zelenaya.net%2Fannounce&tr=http%3A%2F%2Ftracker.electro-torrent.pl%2Fannounce&tr=http%3A%2F%2Ftracker.nyaa.uk%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.torrentyorg.pl%2Fannounce&tr=http%3A%2F%2Ftracker01.loveapp.com%3A6789%2Fannounce&tr=https%3A%2F%2Ftracker.nanoha.org%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.com%3A6869%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:aed51c6fb264528757ff7b4aa630905840799f4a&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">9.5GB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">-</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">-</span></td>
			<td nowrap="nowrap" align="center">-</td>
			<td align="center"><a href="/topics/list/user_id/743372">黑暗路基艾尔</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 20:57			<span style="display: none;">2022/02/21 20:57</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/619" >
				桜都字幕组</a></span>				
				<a href="/topics/view/593410_Ouji_no_Honmei_wa_Akuyaku_Reijou_06_1080p.html"  target="_blank" >
				[桜都字幕组] 王子的本命是反派大小姐 / Ouji no Honmei wa Akuyaku Reijou [06][1080p][简繁内封]</a>
				<span style="color: gray;">約2條評論</span>			</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:YGGTPWAW6AP4Z7OZG54Q3U3XYLXTI5GI&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Fopen.nyaatorrents.info%3A6544%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=https%3A%2F%2Ftracker.sakurato.art%3A23334%2Fannounce&tr=http%3A%2F%2Ftracker.sakurato.art%3A23333%2Fannounce&tr=https%3A%2F%2Ftr.bangumi.moe%3A9696%2Fannounce&tr=http%3A%2F%2Fsukebei.tracker.wf%3A8888%2Fannounce&tr=http%3A%2F%2Fopen.acgnxtracker.com%2Fannounce&tr=https%3A%2F%2Fopen.acgnxtracker.com%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:c18d37d816f01fccfdd937790dd377c2ef3474c8&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">96.5MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">-</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">-</span></td>
			<td nowrap="nowrap" align="center">-</td>
			<td align="center"><a href="/topics/list/user_id/676357">sakurato</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 20:57			<span style="display: none;">2022/02/21 20:57</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/619" >
				桜都字幕组</a></span>				
				<a href="/topics/view/593409_Ouji_no_Honmei_wa_Akuyaku_Reijou_06_1080p.html"  target="_blank" >
				[桜都字幕組] 王子的本命是反派大小姐 / Ouji no Honmei wa Akuyaku Reijou [06][1080p][繁體內嵌]</a>
							</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:6MMIQ3NPZGKLEJT7EOZLYNMKCGYQ6T2N&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Fopen.nyaatorrents.info%3A6544%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=https%3A%2F%2Ftracker.sakurato.art%3A23334%2Fannounce&tr=http%3A%2F%2Ftracker.sakurato.art%3A23333%2Fannounce&tr=https%3A%2F%2Ftr.bangumi.moe%3A9696%2Fannounce&tr=http%3A%2F%2Fsukebei.tracker.wf%3A8888%2Fannounce&tr=http%3A%2F%2Fopen.acgnxtracker.com%2Fannounce&tr=https%3A%2F%2Fopen.acgnxtracker.com%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:f318886dafc994b2267f23b2bc358a11b10f4f4d&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">90.9MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">-</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">-</span></td>
			<td nowrap="nowrap" align="center">-</td>
			<td align="center"><a href="/topics/list/user_id/676357">sakurato</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 20:56			<span style="display: none;">2022/02/21 20:56</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/619" >
				桜都字幕组</a></span>				
				<a href="/topics/view/593408_Ouji_no_Honmei_wa_Akuyaku_Reijou_06_1080p.html"  target="_blank" >
				[桜都字幕组] 王子的本命是反派大小姐 / Ouji no Honmei wa Akuyaku Reijou [06][1080p][简体内嵌]</a>
							</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:QEPMDRQQ243HAH3YPM74IMEERCTDUETU&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Fopen.nyaatorrents.info%3A6544%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=https%3A%2F%2Ftracker.sakurato.art%3A23334%2Fannounce&tr=http%3A%2F%2Ftracker.sakurato.art%3A23333%2Fannounce&tr=https%3A%2F%2Ftr.bangumi.moe%3A9696%2Fannounce&tr=http%3A%2F%2Fsukebei.tracker.wf%3A8888%2Fannounce&tr=http%3A%2F%2Fopen.acgnxtracker.com%2Fannounce&tr=https%3A%2F%2Fopen.acgnxtracker.com%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:811ec1c610d736701f787b3fc4308488a63a1274&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">90.7MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">-</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">-</span></td>
			<td nowrap="nowrap" align="center">-</td>
			<td align="center"><a href="/topics/list/user_id/676357">sakurato</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 20:24			<span style="display: none;">2022/02/21 20:24</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/37" >
				雪飄工作室(FLsnow)</a></span>				
				<a href="/topics/view/593407_Delicious_Party_Precure_WEBDL_1080p_03_Q.html"  target="_blank" >
				 [雪飘工作室][Delicious Party Precure/デリシャスパーティ プリキュア][WEBDL][1080p][03][简繁外挂](检索:光之美少女/Q娃)</a>
				<span style="color: gray;">約1條評論</span>			</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:CT5M6H6EXFIAOK45Z6CQCRKGOBVELC2U&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker.ktxp.com%3A6868%2Fannounce&tr=http%3A%2F%2Ftracker.ktxp.com%3A7070%2Fannounce&tr=http%3A%2F%2F61.154.116.205%3A8000%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2F208.67.16.113%3A8000%2Fannonuce&tr=http%3A%2F%2F121.14.98.151%3A9090%2Fannounce&tr=http%3A%2F%2F94.228.192.98%2Fannounce&tr=http%3A%2F%2Fbt.dmhy.net%2Fannonuce&tr=http%3A%2F%2Ftracker.btcake.com%2Fannounce&tr=http%3A%2F%2Ftracker.ipv6tracker.org%3A80%2Fannounce&tr=http%3A%2F%2Fbt.sc-ol.com%3A2710%2Fannounce&tr=http%3A%2F%2Fshare.dmhy.me%2Fannonuce&tr=http%3A%2F%2Fbtfile.sdo.com%3A6961%2Fannounce&tr=https%3A%2F%2Ftr.bangumi.moe%3A9696%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker3.torrentino.com%2Fannounce&tr=http%3A%2F%2Ftracker2.torrentino.com%2Fannounce&tr=http%3A%2F%2Fpubt.net%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.tfile.me%2Fannounce&tr=http%3A%2F%2Fbigfoot1942.sektori.org%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2F104.238.198.186%3A8000%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:14facf1fc4b950072b9dcf85014546706a458b54&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">1.4GB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">-</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">-</span></td>
			<td nowrap="nowrap" align="center">-</td>
			<td align="center"><a href="/topics/list/user_id/8721">MingHyuk</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 20:02			<span style="display: none;">2022/02/21 20:02</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/801" >
				NC-Raws</a></span>				
				<a href="/topics/view/593406_NC-Raws_IMMORTALITY_-_05_B-Global_Donghua_1920x1080_HEVC_AAC_MKV.html"  target="_blank" >
				[NC-Raws] 永生 / IMMORTALITY - 05 (B-Global Donghua 1920x1080 HEVC AAC MKV)</a>
							</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:KQXZREIXT2KEA7VAEYVOYDCWHP5CEGJZ&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Fopen.nyaatorrents.info%3A6544%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=https%3A%2F%2Ftr.bangumi.moe%3A9696%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:542f9891179e94407ea0262aec0c563bfa221939&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">661.7MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">-</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">-</span></td>
			<td nowrap="nowrap" align="center">-</td>
			<td align="center"><a href="/topics/list/user_id/637871">九十九朔夜</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 20:01			<span style="display: none;">2022/02/21 20:01</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/801" >
				NC-Raws</a></span>				
				<a href="/topics/view/593405_NC-Raws_BUSTED%21_DARKLORD_-_07_B-Global_Donghua_1920x1080_HEVC_AAC_MKV.html"  target="_blank" >
				[NC-Raws] 小魔头暴露啦 / BUSTED! DARKLORD - 07 (B-Global Donghua 1920x1080 HEVC AAC MKV)</a>
							</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:23H3AGTHNOQGU3V2JTMIBWIBUWXFA3XR&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Fopen.nyaatorrents.info%3A6544%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=https%3A%2F%2Ftr.bangumi.moe%3A9696%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:d6cfb01a676ba06a6eba4cd880d901a5ae506ef1&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">216.6MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">-</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">-</span></td>
			<td nowrap="nowrap" align="center">-</td>
			<td align="center"><a href="/topics/list/user_id/637871">九十九朔夜</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 19:47			<span style="display: none;">2022/02/21 19:47</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/47" >
				爱恋字幕社</a></span>				
				<a href="/topics/view/593404_1_05_720p_MP4_GB.html"  target="_blank" >
				[爱恋&漫猫字幕组][1月新番][终末的后宫][05][720p][MP4][GB][简中]</a>
							</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:J57YFK2CI4XC5V6LNSTI2LVFJM7MM22K&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Fopen.nyaatorrents.info%3A6544%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=http%3A%2F%2F104.238.198.186%3A8000%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:4f7f82ab42472e2ed7cb6ca68d2ea54b3ec66b4a&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">188.4MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">-</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">-</span></td>
			<td nowrap="nowrap" align="center">-</td>
			<td align="center"><a href="/topics/list/user_id/488934">KissSub</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 19:47			<span style="display: none;">2022/02/21 19:47</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/47" >
				爱恋字幕社</a></span>				
				<a href="/topics/view/593403_1_05_720p_MP4_BIG5.html"  target="_blank" >
				[愛戀&漫貓字幕組][1月新番][終末的後宮][05][720p][MP4][BIG5][繁中]</a>
							</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:CLKUAKQGVEOYNEE74JFCXV2FD3QT74GI&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Fopen.nyaatorrents.info%3A6544%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=http%3A%2F%2F104.238.198.186%3A8000%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:12d5402a06a91d86909fe24a2bd7451ee13ff0c8&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">188.3MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">-</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">-</span></td>
			<td nowrap="nowrap" align="center">-</td>
			<td align="center"><a href="/topics/list/user_id/488934">KissSub</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 19:47			<span style="display: none;">2022/02/21 19:47</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/47" >
				爱恋字幕社</a></span>				
				<a href="/topics/view/593402_1_04_720p_MP4_GB.html"  target="_blank" >
				[爱恋&漫猫字幕组][1月新番][终末的后宫][04][720p][MP4][GB][简中]</a>
							</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:MOJAEUVSR7CFEWKBR4X43IN2YYRTKD6K&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Fopen.nyaatorrents.info%3A6544%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=http%3A%2F%2F104.238.198.186%3A8000%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:63920252b28fc45259418f2fcda1bac623350fca&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">208.4MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">-</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">-</span></td>
			<td nowrap="nowrap" align="center">-</td>
			<td align="center"><a href="/topics/list/user_id/488934">KissSub</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 19:47			<span style="display: none;">2022/02/21 19:47</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/47" >
				爱恋字幕社</a></span>				
				<a href="/topics/view/593401_1_04_720p_MP4_BIG5.html"  target="_blank" >
				[愛戀&漫貓字幕組][1月新番][終末的後宮][04][720p][MP4][BIG5][繁中]</a>
							</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:MMY4Z6R4MPUQE6ME6DY3ZBYVNCWQPVX3&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Fopen.nyaatorrents.info%3A6544%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=http%3A%2F%2F104.238.198.186%3A8000%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:6331ccfa3c63e9027984f0f1bc871568ad07d6fb&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">208.4MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">-</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">-</span></td>
			<td nowrap="nowrap" align="center">-</td>
			<td align="center"><a href="/topics/list/user_id/488934">KissSub</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 19:11			<span style="display: none;">2022/02/21 19:11</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/47" >
				爱恋字幕社</a></span>				
				<a href="/topics/view/593400_1_03_720p_MP4_GB.html"  target="_blank" >
				[爱恋&漫猫字幕组][1月新番][终末的后宫][03][720p][MP4][GB][简中]</a>
							</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:BVCRPU23EESS2PPOFXQBXVKRR5PROB2R&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Fopen.nyaatorrents.info%3A6544%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=http%3A%2F%2F104.238.198.186%3A8000%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:0d4517d35b21252d3dee2de01bd5518f5f170751&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">196.8MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">-</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">-</span></td>
			<td nowrap="nowrap" align="center">-</td>
			<td align="center"><a href="/topics/list/user_id/488934">KissSub</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 19:11			<span style="display: none;">2022/02/21 19:11</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/47" >
				爱恋字幕社</a></span>				
				<a href="/topics/view/593399_1_03_720p_MP4_BIG5.html"  target="_blank" >
				[愛戀&漫貓字幕組][1月新番][終末的後宮][03][720p][MP4][BIG5][繁中]</a>
							</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:74LRRG24KFOFOPVW3LVLI74GLR3DSJIU&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Fopen.nyaatorrents.info%3A6544%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=http%3A%2F%2F104.238.198.186%3A8000%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:ff17189b5c515c573eb6daeab47f865c76392514&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">196.7MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">-</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">-</span></td>
			<td nowrap="nowrap" align="center">-</td>
			<td align="center"><a href="/topics/list/user_id/488934">KissSub</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 19:11			<span style="display: none;">2022/02/21 19:11</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/47" >
				爱恋字幕社</a></span>				
				<a href="/topics/view/593398_1_02_720p_MP4_GB.html"  target="_blank" >
				[爱恋&漫猫字幕组][1月新番][终末的后宫][02][720p][MP4][GB][简中]</a>
							</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:A6BA23JWKPXUW2IZI7B3A563AKXCZ4KD&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Fopen.nyaatorrents.info%3A6544%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=http%3A%2F%2F104.238.198.186%3A8000%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:07820d6d3653ef4b691947c3b077db02ae2cf143&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">187.3MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">-</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">-</span></td>
			<td nowrap="nowrap" align="center">-</td>
			<td align="center"><a href="/topics/list/user_id/488934">KissSub</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 19:10			<span style="display: none;">2022/02/21 19:10</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/47" >
				爱恋字幕社</a></span>				
				<a href="/topics/view/593397_1_02_720p_MP4_BIG5.html"  target="_blank" >
				[愛戀&漫貓字幕組][1月新番][終末的後宮][02][720p][MP4][BIG5][繁中]</a>
							</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:27T2CKZFFJMR52ISBAHRBKPQXTXSERER&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Fopen.nyaatorrents.info%3A6544%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=http%3A%2F%2F104.238.198.186%3A8000%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:d7e7a12b252a591ee912080f10a9f0bcef224491&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">187.3MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">-</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">-</span></td>
			<td nowrap="nowrap" align="center">-</td>
			<td align="center"><a href="/topics/list/user_id/488934">KissSub</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 18:45			<span style="display: none;">2022/02/21 18:45</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/619" >
				桜都字幕组</a></span>				
				<a href="/topics/view/593396_The_Case_Study_of_Vanitas_18_1080p.html"  target="_blank" >
				[漫貓&桜都字幕組] 瓦尼塔斯的手札 / The Case Study of Vanitas [18][1080p][繁體內嵌]</a>
				<span style="color: gray;">約1條評論</span>			</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:4IX3IPH2LYSXQZIMV35CVMTZXXH6FULV&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Fopen.nyaatorrents.info%3A6544%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=https%3A%2F%2Ftracker.sakurato.art%3A23334%2Fannounce&tr=http%3A%2F%2Ftracker.sakurato.art%3A23333%2Fannounce&tr=https%3A%2F%2Ftr.bangumi.moe%3A9696%2Fannounce&tr=http%3A%2F%2Fsukebei.tracker.wf%3A8888%2Fannounce&tr=http%3A%2F%2Fopen.acgnxtracker.com%2Fannounce&tr=https%3A%2F%2Fopen.acgnxtracker.com%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:e22fb43cfa5e2578650caefa2ab279bdcfe2d175&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">363.5MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">-</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">-</span></td>
			<td nowrap="nowrap" align="center">-</td>
			<td align="center"><a href="/topics/list/user_id/676357">sakurato</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 18:42			<span style="display: none;">2022/02/21 18:42</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/619" >
				桜都字幕组</a></span>				
				<a href="/topics/view/593395_The_Case_Study_of_Vanitas_18_1080p.html"  target="_blank" >
				[漫猫&桜都字幕组] 瓦尼塔斯的手札 / The Case Study of Vanitas [18][1080p][简体内嵌]</a>
				<span style="color: gray;">約1條評論</span>			</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:QF7CXDCOTEOJPCLIJ3BGKOLBIJUCBH32&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Fopen.nyaatorrents.info%3A6544%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=https%3A%2F%2Ftracker.sakurato.art%3A23334%2Fannounce&tr=http%3A%2F%2Ftracker.sakurato.art%3A23333%2Fannounce&tr=https%3A%2F%2Ftr.bangumi.moe%3A9696%2Fannounce&tr=http%3A%2F%2Fsukebei.tracker.wf%3A8888%2Fannounce&tr=http%3A%2F%2Fopen.acgnxtracker.com%2Fannounce&tr=https%3A%2F%2Fopen.acgnxtracker.com%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:817e2b8c4e991c9789684ec26539614268209f7a&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">363.6MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">-</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">-</span></td>
			<td nowrap="nowrap" align="center">-</td>
			<td align="center"><a href="/topics/list/user_id/676357">sakurato</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 18:30			<span style="display: none;">2022/02/21 18:30</span></td>
			<td width="6%" align="center">
			<a class="sort-31" 
				href="/topics/list/sort_id/31">
				<b><font color=red>季度全集</font></b></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/790" >
				WBX-SUB</a></span>				
				<a href="/topics/view/593394_WBX-Raws_MAHO_GIRLS_PRECURE%21_Precure_TV_EP01-50_BDrip_HEVC_1080P_FLAC.html"  target="_blank" >
				[WBX-Raws] 魔法つかいプリキュア！/MAHO GIRLS PRECURE!/魔法使Precure！TV EP01-50 全 [BDrip][HEVC 1080P FLAC]</a>
				<span style="color: gray;">約2條評論</span>			</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:NE3RSTKZL75C63I7BWBT4SD242JOCKYN&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:6937194d595ffa2f6d1f0d833e487ae692e12b0d&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">56.5GB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">-</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">-</span></td>
			<td nowrap="nowrap" align="center">-</td>
			<td align="center"><a href="/topics/list/user_id/728142">Seconder</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 18:21			<span style="display: none;">2022/02/21 18:21</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/767" >
				天月動漫&發佈組</a></span>				
				<a href="/topics/view/593393_Part6_Lupin_III_-_Part_6_-_19_1080P.html"  target="_blank" >
				[天月搬運組] 魯邦三世 Part6 / Lupin III - Part 6 - 19 [1080P][簡繁日外掛]</a>
				<span style="color: gray;">約1條評論</span>			</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:2UZBNFSN63MP7DFUAEETU3PT35VJSR77&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=http%3A%2F%2Fanidex.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.anirena.com%3A80%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:d53216964df6d8ff8cb401093a6df3df6a9947ff&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">207MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">-</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">-</span></td>
			<td nowrap="nowrap" align="center">-</td>
			<td align="center"><a href="/topics/list/user_id/730809">Laputa</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 18:19			<span style="display: none;">2022/02/21 18:19</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/767" >
				天月動漫&發佈組</a></span>				
				<a href="/topics/view/593392_Baraou_no_Souretsu_-_07_1080P.html"  target="_blank" >
				[天月搬運組] 薔薇王的葬列 / Baraou no Souretsu - 07 [1080P][簡繁日外掛]</a>
				<span style="color: gray;">約1條評論</span>			</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:EIVE6C2YG7DN3UE2XWBGBBRGCH2OTROR&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=http%3A%2F%2Fanidex.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.anirena.com%3A80%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:222a4f0b5837c6ddd09abd8260862611f4e9c5d1&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">182MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">-</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">-</span></td>
			<td nowrap="nowrap" align="center">-</td>
			<td align="center"><a href="/topics/list/user_id/730809">Laputa</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 17:36			<span style="display: none;">2022/02/21 17:36</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/669" >
				喵萌奶茶屋</a></span>				
				<a href="/topics/view/593391_01_Sono_Bisque_Doll_wa_Koi_wo_Suru_07_1080p.html"  target="_blank" >
				【喵萌奶茶屋】★01月新番★[更衣人偶墜入愛河/戀上換裝娃娃/Sono Bisque Doll wa Koi wo Suru][07][1080p][繁體][招募翻譯校對]</a>
							</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:CGSKPKDS3GG7K6K4OLKABP34C2KHKPKI&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Fopen.nyaatorrents.info%3A6544%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%3A80%2Fannounce&tr=https%3A%2F%2Fopen.acgnxtracker.com%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce&tr=http%3A%2F%2Fsukebei.tracker.wf%3A8888%2Fannounce&tr=http%3A%2F%2Fp4p.arenabg.com%3A1337%2Fannounce&tr=http%3A%2F%2Fopenbittorrent.com%3A80%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:11a4a7a872d98df5795c72d400bf7c1694753d48&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">326.5MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">10</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">7</span></td>
			<td nowrap="nowrap" align="center">15</td>
			<td align="center"><a href="/topics/list/user_id/693094">nekomoekissaten</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 17:36			<span style="display: none;">2022/02/21 17:36</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/669" >
				喵萌奶茶屋</a></span>				
				<a href="/topics/view/593390_01_Sono_Bisque_Doll_wa_Koi_wo_Suru_07_720p.html"  target="_blank" >
				【喵萌奶茶屋】★01月新番★[更衣人偶墜入愛河/戀上換裝娃娃/Sono Bisque Doll wa Koi wo Suru][07][720p][繁體][招募翻譯校對]</a>
							</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:NMV5P7DVRNMHS6D6IDKZE6F677KLH5RH&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Fopen.nyaatorrents.info%3A6544%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%3A80%2Fannounce&tr=https%3A%2F%2Fopen.acgnxtracker.com%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce&tr=http%3A%2F%2Fsukebei.tracker.wf%3A8888%2Fannounce&tr=http%3A%2F%2Fp4p.arenabg.com%3A1337%2Fannounce&tr=http%3A%2F%2Fopenbittorrent.com%3A80%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:6b2bd7fc758b5879787e40d59278beffd4b3f627&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">147MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">4</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">2</span></td>
			<td nowrap="nowrap" align="center">5</td>
			<td align="center"><a href="/topics/list/user_id/693094">nekomoekissaten</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 14:31			<span style="display: none;">2022/02/21 14:31</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/488" >
				丸子家族</a></span>				
				<a href="/topics/view/593389_Chibi_Maruko-chan_II_1326_2022_02_20_BIG5_1080P_MP4.html"  target="_blank" >
				[丸子家族][櫻桃小丸子第二期(Chibi Maruko-chan II)][1326]前田的紐扣紛爭&媽媽的神秘備忘錄[2022.02.20][BIG5][1080P][MP4]</a>
				<span style="color: gray;">約1條評論</span>			</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:OACRMMDMHYABXL4T7ZL22J52VICZYP4P&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Fopen.nyaatorrents.info%3A6544%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=http%3A%2F%2Fp4p.arenabg.com%3A1337%2Fannounce&tr=http%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&tr=http%3A%2F%2Fopenbittorrent.com%3A80%2Fannounce&tr=https%3A%2F%2Ftracker.lilithraws.cf%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.iriseden.fr%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.iriseden.eu%3A443%2Fannounce&tr=http%3A%2F%2Ftracker1.bt.moack.co.kr%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.moeking.me%3A6969%2Fannounce&tr=http%3A%2F%2Ftr.cili001.com%3A8070%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%3A80%2Fannounce&tr=http%3A%2F%2Fretracker.sevstar.net%3A2710%2Fannounce&tr=http%3A%2F%2Fretracker.joxnet.ru%3A80%2Fannounce&tr=http%3A%2F%2Ffxtt.ru%3A80%2Fannounce&tr=http%3A%2F%2Fbt.okmp3.ru%3A2710%2Fannounce&tr=https%3A%2F%2Ftrackme.theom.nz%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.tamersunion.org%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.nitrix.me%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.kuroy.me%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.foreverpirates.co%3A443%2Fannounce&tr=https%3A%2F%2Ftr.torland.ga%3A443%2Fannounce&tr=https%3A%2F%2F1337.abcvg.info%3A443%2Fannounce&tr=http%3A%2F%2Fvps02.net.orel.ru%3A80%2Fannounce&tr=http%3A%2F%2Ftrackme.theom.nz%3A80%2Fannounce&tr=http%3A%2F%2Ftracker2.dler.org%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.zerobytes.xyz%3A1337%2Fannounce&tr=http%3A%2F%2Ftracker.noobsubs.net%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.loadbt.com%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.files.fm%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.dler.org%3A6969%2Fannounce&tr=http%3A%2F%2Ftorrenttracker.nwc.acsalaska.net%3A6969%2Fannounce&tr=http%3A%2F%2Ft.overflow.biz%3A6969%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce&tr=http%3A%2F%2Fpow7.com%3A80%2Fannounce&tr=http%3A%2F%2Fopen.acgnxtracker.com%3A80%2Fannounce&tr=http%3A%2F%2F1337.abcvg.info%3A80%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:700516306c3e001baf93fe57ad27baaa059c3f8f&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">193.2MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">3</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">0</span></td>
			<td nowrap="nowrap" align="center">15</td>
			<td align="center"><a href="/topics/list/user_id/482886">丸子家族</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 14:31			<span style="display: none;">2022/02/21 14:31</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/488" >
				丸子家族</a></span>				
				<a href="/topics/view/593388_Chibi_Maruko-chan_II_1326_2022_02_20_GB_1080P_MP4.html"  target="_blank" >
				[丸子家族][樱桃小丸子第二期(Chibi Maruko-chan II)][1326]前田的纽扣纷争&妈妈的神秘备忘录[2022.02.20][GB][1080P][MP4]</a>
				<span style="color: gray;">約2條評論</span>			</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:OCCX3EQYOKEKMBIYX5C4ERLRMPANZUAZ&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Fopen.nyaatorrents.info%3A6544%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=http%3A%2F%2Fp4p.arenabg.com%3A1337%2Fannounce&tr=http%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&tr=http%3A%2F%2Fopenbittorrent.com%3A80%2Fannounce&tr=https%3A%2F%2Ftracker.lilithraws.cf%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.iriseden.fr%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.iriseden.eu%3A443%2Fannounce&tr=http%3A%2F%2Ftracker1.bt.moack.co.kr%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.moeking.me%3A6969%2Fannounce&tr=http%3A%2F%2Ftr.cili001.com%3A8070%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%3A80%2Fannounce&tr=http%3A%2F%2Fretracker.sevstar.net%3A2710%2Fannounce&tr=http%3A%2F%2Fretracker.joxnet.ru%3A80%2Fannounce&tr=http%3A%2F%2Ffxtt.ru%3A80%2Fannounce&tr=http%3A%2F%2Fbt.okmp3.ru%3A2710%2Fannounce&tr=https%3A%2F%2Ftrackme.theom.nz%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.tamersunion.org%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.nitrix.me%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.kuroy.me%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.foreverpirates.co%3A443%2Fannounce&tr=https%3A%2F%2Ftr.torland.ga%3A443%2Fannounce&tr=https%3A%2F%2F1337.abcvg.info%3A443%2Fannounce&tr=http%3A%2F%2Fvps02.net.orel.ru%3A80%2Fannounce&tr=http%3A%2F%2Ftrackme.theom.nz%3A80%2Fannounce&tr=http%3A%2F%2Ftracker2.dler.org%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.zerobytes.xyz%3A1337%2Fannounce&tr=http%3A%2F%2Ftracker.noobsubs.net%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.loadbt.com%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.files.fm%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.dler.org%3A6969%2Fannounce&tr=http%3A%2F%2Ftorrenttracker.nwc.acsalaska.net%3A6969%2Fannounce&tr=http%3A%2F%2Ft.overflow.biz%3A6969%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce&tr=http%3A%2F%2Fpow7.com%3A80%2Fannounce&tr=http%3A%2F%2Fopen.acgnxtracker.com%3A80%2Fannounce&tr=http%3A%2F%2F1337.abcvg.info%3A80%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:70857d92187288a60518bf45c2457163c0dcd019&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">193.1MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">0</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">1</span></td>
			<td nowrap="nowrap" align="center">7</td>
			<td align="center"><a href="/topics/list/user_id/482886">丸子家族</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 14:13			<span style="display: none;">2022/02/21 14:13</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
								
				<a href="/topics/view/593387_jibaketa_TVB_Z_Z_Shinkansen_Henkei_Robo_Shinkalion_Z_-_15_WEB_1920x1080_AVC_AACx2_SRT_TVB_CHT.html"  target="_blank" >
				[jibaketa合成&音頻壓制][TVB粵語]新幹線戰士Z / 新幹線變形機器人Z / Shinkansen Henkei Robo Shinkalion Z - 15 [粵日雙語+內封繁體中文字幕][WEB 1920x1080 AVC AACx2 SRT TVB CHT]</a>
				<span style="color: gray;">約1條評論</span>			</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:BAB5TKDQ7T6VFLLQ5WPTSH5XCEZNP2JO&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce&tr=https%3A%2F%2Ftr.bangumi.moe%3A9696%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:0803d9a870fcfd52ad70ed9f391fb71132d7e92e&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">934.3MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">6</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">1</span></td>
			<td nowrap="nowrap" align="center">16</td>
			<td align="center"><a href="/topics/list/user_id/672257">jibaketa</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 13:51			<span style="display: none;">2022/02/21 13:51</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/185" >
				极影字幕社</a></span>				
				<a href="/topics/view/593386_10_Shingeki_no_Kyojin_The_Final_Season_22_GB_MP4_720.html"  target="_blank" >
				【极影字幕社】 ★10月新番 【进击的巨人 最终季】【Shingeki no Kyojin The Final Season】【22】GB MP4_720</a>
				<span style="color: gray;">約2條評論</span>			</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:6XJQ542JEN7T7AUPGEKGQ5PUX7RXTZZD&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.ktxp.com%3A6868%2Fannounce&tr=http%3A%2F%2Ftracker.ktxp.com%3A7070%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%2Fannounce&tr=http%3A%2F%2Ftracker.dmhy.org%3A8000%2Fannounce&tr=http%3A%2F%2Fbt.popgo.net%3A7456%2Fannounce&tr=http%3A%2F%2Fnyaatorrents.info%3A7266%2Fannounce&tr=http%3A%2F%2Fnyaatorrents.info%3A3277%2Fannounce&tr=http%3A%2F%2Fbt.sc-ol.com%3A2710%2Fannounce&tr=http%3A%2F%2Fanisource.spb.ru%2Fannounce&tr=http%3A%2F%2Ftracker.anirena.com%3A81%2Fannounce&tr=http%3A%2F%2Ftracker.anirena.com%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.frozen-layer.net%3A6969%2Fannounce&tr=http%3A%2F%2Fmkfs.ru%2Fannounce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker.levelup.cn%3A8080%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fbt.wiiyi.com%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.dmguo.com%3A2710%2Fannounce&tr=http%3A%2F%2Fbt1.the9.com%3A6969%2Fannounce&tr=http%3A%2F%2Ftorrent-download.to%3A5869%2Fannounce&tr=http%3A%2F%2Ftorrent-download.to%3A5869%2Fannounce.php&tr=http%3A%2F%2Ftorrent-downloads.to%3A5869%2Fannounce&tr=http%3A%2F%2Ftorrent-downloads.to%3A5869%2Fannounce.php&tr=http%3A%2F%2Ftracker.bittorrent.am%2Fannounce&tr=http%3A%2F%2Ftracker.torrent.to%3A2710%2Fannounce&tr=http%3A%2F%2Ftv.tracker.prq.to%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce.php&tr=http%3A%2F%2Ftracker9.bol.bg%2Fannounce.php&tr=http%3A%2F%2Fdenis.stalker.h3q.com%2Fannounce&tr=http%3A%2F%2Fdenis.stalker.h3q.com%2Fannounce.php&tr=http%3A%2F%2Fdenis.stalker.h3q.com%3A6969%2Fannounce&tr=http%3A%2F%2Fdenis.stalker.h3q.com%3A6969%2Fannounce.php&tr=http%3A%2F%2Fbttrack.9you.com%3A8080%2Fannounce&tr=http%3A%2F%2F9.rarbg.com%3A2710%2Fannounce&tr=http%3A%2F%2Fgenesis.1337x.org%3A1337%2Fannounce&tr=http%3A%2F%2Ftpb.tracker.prq.to%2Fannounce&tr=http%3A%2F%2Fbt.edwardk.info%3A6969%2Fannounce&tr=http%3A%2F%2Fbt.sumisora.com%3A2710%2Fannounce&tr=http%3A%2F%2Fshare.hkg-fansub.info%2Fannounce.php">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:f5d30ef349237f3f828f31146875f4bfe379e723&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">181.9MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">14</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">2</span></td>
			<td nowrap="nowrap" align="center">102</td>
			<td align="center"><a href="/topics/list/user_id/403163">rossina</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 13:36			<span style="display: none;">2022/02/21 13:36</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/805" >
				DBD制作组</a></span>				
				<a href="/topics/view/593385_DBD-Raws_Dragon_Ball_Majinjou_no_Nemurihime_1080P_BDRip_HEVC-10bit_FLAC_MKV.html"  target="_blank" >
				[DBD-Raws][龙珠剧场版：魔神城内的睡美人/Dragon Ball: Majinjou no Nemurihime/ドラゴンボール 魔神城のねむり姫][1080P][BDRip][HEVC-10bit][FLAC][MKV]</a>
				<span style="color: gray;">約3條評論</span>			</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:53SBSMOONNJW7ZZPRL27UH75R6GFNZMI&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2F1337.abcvg.info%2Fannounce&tr=http%3A%2F%2F178.175.143.27%2Fannounce&tr=http%3A%2F%2F5rt.tace.ru%3A60889%2Fannounce&tr=http%3A%2F%2F78.30.254.12%3A2710%2Fannounce&tr=http%3A%2F%2F91.217.91.21%3A3218%2Fannounce&tr=http%3A%2F%2F%5B2001%3A1b10%3A1000%3A8101%3A0%3A242%3Aac11%3A2%5D%3A6969%2Fannounce&tr=http%3A%2F%2F%5B2001%3A470%3A1%3A189%3A0%3A1%3A2%3A3%5D%3A6969%2Fannounce&tr=http%3A%2F%2F%5B2a04%3Aac00%3A1%3A3dd8%3A%3A1%3A2710%5D%3A2710%2Fannounce&tr=http%3A%2F%2Faaa.army%3A8866%2Fannounce&tr=http%3A%2F%2Fatrack.pow7.com%2Fannounce&tr=http%3A%2F%2Fbobbialbano.com%3A6969%2Fannounce&tr=http%3A%2F%2Fbt.3kb.xyz%2Fannounce&tr=http%3A%2F%2Fbt.pusacg.org%3A8080%2Fannounce&tr=http%3A%2F%2Fderpyradio.net%3A6969%2Fannounce&tr=http%3A%2F%2Fexplodie.org%3A6969%2Fannounce&tr=http%3A%2F%2Fgrifon.info%2Fannounce&tr=http%3A%2F%2Fh4.trakx.nibba.trade%2Fannounce&tr=http%3A%2F%2Fjsb.by%3A8000%2Fannounce&tr=http%3A%2F%2Fnovaopcj.icu%3A10325%2Fannounce&tr=http%3A%2F%2Fns349743.ip-91-121-106.eu%2Fannounce&tr=http%3A%2F%2Fopen.acgnxtracker.com%2Fannounce&tr=http%3A%2F%2Fopen.touki.ru%2Fannounce.php&tr=http%3A%2F%2Fopentracker.i2p.rocks%3A6969%2Fannounce&tr=http%3A%2F%2Fopentracker.xyz%2Fannounce&tr=http%3A%2F%2Fpow7.com%2Fannounce&tr=http%3A%2F%2Fpublictracker.ch%3A6969%2Fannounce&tr=http%3A%2F%2Fretracker.hotplug.ru%3A2710%2Fannounce&tr=http%3A%2F%2Fretracker.sevstar.net%3A2710%2Fannounce&tr=http%3A%2F%2Frt.tace.ru%2Fannounce&tr=http%3A%2F%2Fsecure.pow7.com%2Fannounce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Ft.overflow.biz%3A6969%2Fannounce&tr=http%3A%2F%2Ft1.pow7.com%2Fannounce&tr=http%3A%2F%2Ft2.pow7.com%2Fannounce&tr=http%3A%2F%2Fthetracker.org%2Fannounce&tr=http%3A%2F%2Ftorrentsmd.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftorrenttracker.nwc.acsalaska.net%3A6969%2Fannounce&tr=http%3A%2F%2Ftr.cili001.com%3A8070%2Fannounce&tr=http%3A%2F%2Ftracker.anonwebz.xyz%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker.birkenwald.de%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.bittor.pw%3A1337%2Fannounce&tr=http%3A%2F%2Ftracker.bittorrent.am%2Fannounce&tr=http%3A%2F%2Ftracker.bt4g.com%3A2095%2Fannounce&tr=http%3A%2F%2Ftracker.bz%2Fannounce&tr=http%3A%2F%2Ftracker.dler.org%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.dutchtracking.nl%2Fannounce&tr=http%3A%2F%2Ftracker.files.fm%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.gbitt.info%2Fannounce&tr=http%3A%2F%2Ftracker.internetwarriors.net%3A1337%2Fannounce&tr=http%3A%2F%2Ftracker.ipv6tracker.ru%2Fannounce&tr=http%3A%2F%2Ftracker.kamigami.org%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.kuroy.me%3A5944%2Fannounce&tr=http%3A%2F%2Ftracker.lelux.fi%2Fannounce&tr=http%3A%2F%2Ftracker.moeking.me%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.noobsubs.net%2Fannounce&tr=http%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&tr=http%3A%2F%2Ftracker.sakurato.xyz%3A23333%2Fannounce&tr=http%3A%2F%2Ftracker.skyts.net%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.sloppyta.co%2Fannounce&tr=http%3A%2F%2Ftracker.trackerfix.com%2Fannounce&tr=http%3A%2F%2Ftracker.uw0.xyz%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.ygsub.com%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.zerobytes.xyz%3A1337%2Fannounce&tr=http%3A%2F%2Ftracker2.dler.org%2Fannounce&tr=http%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftrun.tom.ru%2Fannounce&tr=http%3A%2F%2Fvpn.flying-datacenter.de%3A6969%2Fannounce&tr=http%3A%2F%2Fvps02.net.orel.ru%2Fannounce&tr=http%3A%2F%2Fwww.wareztorrent.com%2Fannounce&tr=https%3A%2F%2F1337.abcvg.info%2Fannounce&tr=https%3A%2F%2F2.tracker.eu.org%2Fannounce&tr=https%3A%2F%2F3.tracker.eu.org%2Fannounce&tr=https%3A%2F%2Faaa.army%3A8866%2Fannounce&tr=https%3A%2F%2Fopen.kickasstracker.com%2Fannounce&tr=https%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=https%3A%2F%2Ftr.ready4.icu%2Fannounce&tr=https%3A%2F%2Ftracker.bt-hash.com%2Fannounce&tr=https%3A%2F%2Ftracker.gbitt.info%2Fannounce&tr=https%3A%2F%2Ftracker.hama3.net%2Fannounce&tr=https%3A%2F%2Ftracker.imgoingto.icu%2Fannounce&tr=https%3A%2F%2Ftracker.lelux.fi%2Fannounce&tr=https%3A%2F%2Ftracker.lilithraws.cf%2Fannounce&tr=https%3A%2F%2Ftracker.nitrix.me%2Fannounce&tr=https%3A%2F%2Ftracker.parrotsec.org%2Fannounce&tr=https%3A%2F%2Ftracker.sloppyta.co%2Fannounce&tr=https%3A%2F%2Ftracker.tamersunion.org%2Fannounce&tr=https%3A%2F%2Ftrakx.herokuapp.com%2Fannounce&tr=https%3A%2F%2Fw.wwwww.wtf%2Fannounce&tr=https%3A%2F%2Fwww.wareztorrent.com%2Fannounce&tr=http%3A%2F%2F1337.abcvg.info%2Fannounce%2Chttp%3A%2F%2F178.175.143.27%3A80%2Fannounce%2Chttp%3A%2F%2F5rt.tace.ru%3A60889%2Fannounce%2Chttp%3A%2F%2F78.30.254.12%3A2710%2Fannounce%2Chttp%3A%2F%2F91.217.91.21%3A3218%2Fannounce%2Chttp%3A%2F%2F%5B2001%3A1b10%3A1000%3A8101%3A0%3A242%3Aac11%3A2%5D%3A6969%2Fannounce%2Chttp%3A%2F%2F%5B2001%3A470%3A1%3A189%3A0%3A1%3A2%3A3%5D%3A6969%2Fannounce%2Chttp%3A%2F%2F%5B2a04%3Aac00%3A1%3A3dd8%3A%3A1%3A2710%5D%3A2710%2Fannounce%2Chttp%3A%2F%2Faaa.army%3A8866%2Fannounce%2Chttp%3A%2F%2Fatrack.pow7.com%3A80%2Fannounce%2Chttp%3A%2F%2Fbobbialbano.com%3A6969%2Fannounce%2Chttp%3A%2F%2Fbt.3kb.xyz%3A80%2Fannounce%2Chttp%3A%2F%2Fbt.pusacg.org%3A8080%2Fannounce%2Chttp%3A%2F%2Fderpyradio.net%3A6969%2Fannounce%2Chttp%3A%2F%2Fexplodie.org%3A6969%2Fannounce%2Chttp%3A%2F%2Fgrifon.info%3A80%2Fannounce%2Chttp%3A%2F%2Fh4.trakx.nibba.trade%3A80%2Fannounce%2Chttp%3A%2F%2Fjsb.by%3A8000%2Fannounce%2Chttp%3A%2F%2Fnovaopcj.icu%3A10325%2Fannounce%2Chttp%3A%2F%2Fns349743.ip-91-121-106.eu%3A80%2Fannounce%2Chttp%3A%2F%2Fopen.acgnxtracker.com%3A80%2Fannounce%2Chttp%3A%2F%2Fopen.touki.ru%3A80%2Fannounce.php%2Chttp%3A%2F%2Fopentracker.i2p.rocks%3A6969%2Fannounce%2Chttp%3A%2F%2Fopentracker.xyz%3A80%2Fannounce%2Chttp%3A%2F%2Fpow7.com%3A80%2Fannounce%2Chttp%3A%2F%2Fpublictracker.ch%3A6969%2Fannounce%2Chttp%3A%2F%2Fretracker.hotplug.ru%3A2710%2Fannounce%2Chttp%3A%2F%2Fretracker.sevstar.net%3A2710%2Fannounce%2Chttp%3A%2F%2Frt.tace.ru%3A80%2Fannounce%2Chttp%3A%2F%2Fsecure.pow7.com%3A80%2Fannounce%2Chttp%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce%2Chttp%3A%2F%2Ft.acg.rip%3A6699%2Fannounce%2Chttp%3A%2F%2Ft.nyaatracker.com%3A80%2Fannounce%2Chttp%3A%2F%2Ft.overflow.biz%3A6969%2Fannounce%2Chttp%3A%2F%2Ft1.pow7.com%3A80%2Fannounce%2Chttp%3A%2F%2Ft2.pow7.com%3A80%2Fannounce%2Chttp%3A%2F%2Fthetracker.org%3A80%2Fannounce%2Chttp%3A%2F%2Ftorrentsmd.com%3A8080%2Fannounce%2Chttp%3A%2F%2Ftorrenttracker.nwc.acsalaska.net%3A6969%2Fannounce%2Chttp%3A%2F%2Ftr.cili001.com%3A8070%2Fannounce%2Chttp%3A%2F%2Ftracker.anonwebz.xyz%3A8080%2Fannounce%2Chttp%3A%2F%2Ftracker.birkenwald.de%3A6969%2Fannounce%2Chttp%3A%2F%2Ftracker.bittor.pw%3A1337%2Fannounce%2Chttp%3A%2F%2Ftracker.bittorrent.am%3A80%2Fannounce%2Chttp%3A%2F%2Ftracker.bt4g.com%3A2095%2Fannounce%2Chttp%3A%2F%2Ftracker.bz%3A80%2Fannounce%2Chttp%3A%2F%2Ftracker.dler.org%3A6969%2Fannounce%2Chttp%3A%2F%2Ftracker.dutchtracking.nl%3A80%2Fannounce%2Chttp%3A%2F%2Ftracker.files.fm%3A6969%2Fannounce%2Chttp%3A%2F%2Ftracker.gbitt.info%3A80%2Fannounce%2Chttp%3A%2F%2Ftracker.internetwarriors.net%3A1337%2Fannounce%2Chttp%3A%2F%2Ftracker.ipv6tracker.ru%3A80%2Fannounce%2Chttp%3A%2F%2Ftracker.kamigami.org%3A2710%2Fannounce%2Chttp%3A%2F%2Ftracker.kuroy.me%3A5944%2Fannounce%2Chttp%3A%2F%2Ftracker.lelux.fi%3A80%2Fannounce%2Chttp%3A%2F%2Ftracker.moeking.me%3A6969%2Fannounce%2Chttp%3A%2F%2Ftracker.noobsubs.net%3A80%2Fannounce%2Chttp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce%2Chttp%3A%2F%2Ftracker.sakurato.xyz%3A23333%2Fannounce%2Chttp%3A%2F%2Ftracker.skyts.net%3A6969%2Fannounce%2Chttp%3A%2F%2Ftracker.sloppyta.co%3A80%2Fannounce%2Chttp%3A%2F%2Ftracker.trackerfix.com%3A80%2Fannounce%2Chttp%3A%2F%2Ftracker.uw0.xyz%3A6969%2Fannounce%2Chttp%3A%2F%2Ftracker.ygsub.com%3A6969%2Fannounce%2Chttp%3A%2F%2Ftracker.zerobytes.xyz%3A1337%2Fannounce%2Chttp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce%2Chttp%3A%2F%2Ftracker2.dler.org%3A80%2Fannounce%2Chttp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce%2Chttp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce%2Chttp%3A%2F%2Ftrun.tom.ru%3A80%2Fannounce%2Chttp%3A%2F%2Fvpn.flying-datacenter.de%3A6969%2Fannounce%2Chttp%3A%2F%2Fvps02.net.orel.ru%3A80%2Fannounce%2Chttp%3A%2F%2Fwww.wareztorrent.com%3A80%2Fannounce%2Chttps%3A%2F%2F1337.abcvg.info%3A443%2Fannounce%2Chttps%3A%2F%2F2.tracker.eu.org%3A443%2Fannounce%2Chttps%3A%2F%2F3.tracker.eu.org%3A443%2Fannounce%2Chttps%3A%2F%2Faaa.army%3A8866%2Fannounce%2Chttps%3A%2F%2Fopen.kickasstracker.com%3A443%2Fannounce%2Chttps%3A%2F%2Fopentracker.acgnx.se%3A443%2Fannounce%2Chttps%3A%2F%2Ftr.ready4.icu%3A443%2Fannounce%2Chttps%3A%2F%2Ftracker.bt-hash.com%3A443%2Fannounce%2Chttps%3A%2F%2Ftracker.gbitt.info%3A443%2Fannounce%2Chttps%3A%2F%2Ftracker.hama3.net%3A443%2Fannounce%2Chttps%3A%2F%2Ftracker.imgoingto.icu%3A443%2Fannounce%2Chttps%3A%2F%2Ftracker.lelux.fi%3A443%2Fannounce%2Chttps%3A%2F%2Ftracker.lilithraws.cf%3A443%2Fannounce%2Chttps%3A%2F%2Ftracker.nitrix.me%3A443%2Fannounce%2Chttps%3A%2F%2Ftracker.parrotsec.org%3A443%2Fannounce%2Chttps%3A%2F%2Ftracker.sloppyta.co%3A443%2Fannounce%2Chttps%3A%2F%2Ftracker.tamersunion.org%3A443%2Fannounce%2Chttps%3A%2F%2Ftrakx.herokuapp.com%3A443%2Fannounce%2Chttps%3A%2F%2Fw.wwwww.wtf%3A443%2Fannounce%2Chttps%3A%2F%2Fwww.wareztorrent.com%3A443%2Fannounce%2Cudp%3A%2F%2F104.238.198.186%3A8000%2Fannounce%2Cudp%3A%2F%2F107.150.14.110%3A6969%2Fannounce%2Cudp%3A%2F%2F109.121.134.121%3A1337%2Fannounce%2Cudp%3A%2F%2F114.55.113.60%3A6969%2Fannounce%2Cudp%3A%2F%2F128.199.70.66%3A5944%2Fannounce%2Cudp%3A%2F%2F151.80.120.114%3A2710%2Fannounce%2Cudp%3A%2F%2F168.235.67.63%3A6969%2Fannounce%2Cudp%3A%2F%2F178.33.73.26%3A2710%2Fannounce%2Cudp%3A%2F%2F182.176.139.129%3A6969%2Fannounce%2Cudp%3A%2F%2F185.5.97.139%3A8089%2Fannounce%2Cudp%3A%2F%2F185.83.215.123%3A6969%2Fannounce%2Cudp%3A%2F%2F185.86.149.205%3A1337%2Fannounce%2Cudp%3A%2F%2F188.165.253.109%3A1337%2Fannounce%2Cudp%3A%2F%2F191.101.229.236%3A1337%2Fannounce%2Cudp%3A%2F%2F194.106.216.222%3A80%2Fannounce%2Cudp%3A%2F%2F195.123.209.37%3A1337%2Fannounce%2Cudp%3A%2F%2F195.123.209.40%3A80%2Fannounce%2Cudp%3A%2F%2F208.67.16.113%3A8000%2Fannounce%2Cudp%3A%2F%2F212.1.226.176%3A2710%2Fannounce%2Cudp%3A%2F%2F212.47.227.58%3A6969%2Fannounce%2Cudp%3A%2F%2F213.163.67.56%3A1337%2Fannounce%2Cudp%3A%2F%2F37.19.5.155%3A2710%2Fannounce%2Cudp%3A%2F%2F3rt.tace.ru%3A60889%2Fannounce%2Cudp%3A%2F%2F46.4.109.148%3A6969%2Fannounce%2Cudp%3A%2F%2F47.ip-51-68-199.eu%3A6969%2Fannounce%2Cudp%3A%2F%2F5.79.249.77%3A6969%2Fannounce%2Cudp%3A%2F%2F5.79.83.193%3A6969%2Fannounce%2Cudp%3A%2F%2F51.254.244.161%3A6969%2Fannounce%2Cudp%3A%2F%2F52.58.128.163%3A6969%2Fannounce%2Cudp%3A%2F%2F61626c.net%3A6969%2Fannounce%2Cudp%3A%2F%2F62.138.0.158%3A6969%2Fannounce%2Cudp%3A%2F%2F62.212.85.66%3A2710%2Fannounce%2Cudp%3A%2F%2F74.82.52.209%3A6969%2Fannounce%2Cudp%3A%2F%2F78.30.254.12%3A2710%2Fannounce%2Cudp%3A%2F%2F85.17.19.180%3A80%2Fannounce%2Cudp%3A%2F%2F89.234.156.205%3A80%2Fannounce%2Cudp%3A%2F%2F9.rarbg.com%3A2710%2Fannounce%2Cudp%3A%2F%2F9.rarbg.me%3A2710%2Fannounce%2Cudp%3A%2F%2F9.rarbg.me%3A2780%2Fannounce%2Cudp%3A%2F%2F9.rarbg.to%3A2710%2Fannounce%2Cudp%3A%2F%2F9.rarbg.to%3A2730%2Fannounce%2Cudp%3A%2F%2F91.216.110.52%3A451%2Fannounce%2Cudp%3A%2F%2F91.218.230.81%3A6969%2Fannounce%2Cudp%3A%2F%2F94.23.183.33%3A6969%2Fannounce%2Cudp%3A%2F%2F95.211.168.204%3A2710%2Fannounce%2Cudp%3A%2F%2F%5B2001%3A1b10%3A1000%3A8101%3A0%3A242%3Aac11%3A2%5D%3A6969%2Fannounce%2Cudp%3A%2F%2F%5B2001%3A470%3A1%3A189%3A0%3A1%3A2%3A3%5D%3A6969%2Fannounce%2Cudp%3A%2F%2F%5B2a03%3A7220%3A8083%3Acd00%3A%3A1%5D%3A451%2Fannounce%2Cudp%3A%2F%2F%5B2a04%3Aac00%3A1%3A3dd8%3A%3A1%3A2710%5D%3A2710%2Fannounce%2Cudp%3A%2F%2F%5B2a04%3Ac44%3Ae00%3A32e0%3A4cf%3A6aff%3Afe00%3Aaa1%5D%3A6969%2Fannounce%2Cudp%3A%2F%2Faaa.army%3A8866%2Fannounce%2Cudp%3A%2F%2Fadm.category5.tv%3A6969%2Fannounce%2Cudp%3A%2F%2Fadmin.videoenpoche.info%3A6969%2Fannounce%2Cudp%3A%2F%2Fadminion.n-blade.ru%3A6969%2Fannounce%2Cudp%3A%2F%2Fanidex.moe%3A6969%2Fannounce%2Cudp%3A%2F%2Fapi.bitumconference.ru%3A6969%2Fannounce%2Cudp%3A%2F%2Faruacfilmes.com.br%3A6969%2Fannounce%2Cudp%3A%2F%2Fasger.xyz%3A6969%2Fannounce%2Cudp%3A%2F%2Fbclearning.top%3A6969%2Fannounce%2Cudp%3A%2F%2Fbenouworldtrip.fr%3A6969%2Fannounce%2Cudp%3A%2F%2Fbioquantum.co.za%3A6969%2Fannounce%2Cudp%3A%2F%2Fbitsparadise.info%3A6969%2Fannounce%2Cudp%3A%2F%2Fblokas.io%3A6969%2Fannounce%2Cudp%3A%2F%2Fbms-hosxp.com%3A6969%2Fannounce%2Cudp%3A%2F%2Fbt.firebit.org%3A2710%2Fannounce%2Cudp%3A%2F%2Fbt.okmp3.ru%3A2710%2Fannounce%2Cudp%3A%2F%2Fbt1.archive.org%3A6969%2Fannounce%2Cudp%3A%2F%2Fbt2.54new.com%3A8080%2Fannounce%2Cudp%3A%2F%2Fbt2.archive.org%3A6969%2Fannounce%2Cudp%3A%2F%2Fbubu.mapfactor.com%3A6969%2Fannounce%2Cudp%3A%2F%2Fcamera.lei001.com%3A6969%2Fannounce%2Cudp%3A%2F%2Fcdn-1.gamecoast.org%3A6969%2Fannounce%2Cudp%3A%2F%2Fcdn-2.gamecoast.org%3A6969%2Fannounce%2Cudp%3A%2F%2Fchanchan.uchuu.co.uk%3A6969%2Fannounce%2Cudp%3A%2F%2Fchihaya.toss.li%3A9696%2Fannounce%2Cudp%3A%2F%2Fcode2chicken.nl%3A6969%2Fannounce%2Cudp%3A%2F%2Fconcen.org%3A6969%2Fannounce%2Cudp%3A%2F%2Fcutiegirl.ru%3A6969%2Fannounce%2Cudp%3A%2F%2Fdaveking.com%3A6969%2Fannounce%2Cudp%3A%2F%2Fdenis.stalker.upeer.me%3A6969%2Fannounce%2Cudp%3A%2F%2Fdiscord.heihachi.pw%3A6969%2Fannounce%2Cudp%3A%2F%2Fdpiui.reedlan.com%3A6969%2Fannounce%2Cudp%3A%2F%2Fdrumkitx.com%3A6969%2Fannounce%2Cudp%3A%2F%2Feddie4.nl%3A6969%2Fannounce%2Cudp%3A%2F%2Fedu.uifr.ru%3A6969%2Fannounce%2Cudp%3A%2F%2Fengplus.ru%3A6969%2Fannounce%2Cudp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce%2Cudp%3A%2F%2Ffe.dealclub.de%3A6969%2Fannounce%2Cudp%3A%2F%2Fforever-tracker.zooki.xyz%3A6969%2Fannounce%2Cudp%3A%2F%2Ffree-tracker.zooki.xyz%3A6969%2Fannounce%2Cudp%3A%2F%2Fgit.vulnix.sh%3A6969%2Fannounce%2Cudp%3A%2F%2Finferno.demonoid.is%3A3391%2Fannounce%2Cudp%3A%2F%2Fipv4.tracker.harry.lu%3A80%2Fannounce%2Cudp%3A%2F%2Fipv6.tracker.zerobytes.xyz%3A16661%2Fannounce%2Cudp%3A%2F%2Fjsb.by%3A8000%2Fannounce%2Cudp%3A%2F%2Fkanal-4.de%3A6969%2Fannounce%2Cudp%3A%2F%2Fkoli.services%3A6969%2Fannounce%2Cudp%3A%2F%2Fline-net.ru%3A6969%2Fannounce%2Cudp%3A%2F%2Fln.mtahost.co%3A6969%2Fannounce%2Cudp%3A%2F%2Fmail.realliferpg.de%3A6969%2Fannounce%2Cudp%3A%2F%2Fmgtracker.org%3A2710%2Fannounce%2Cudp%3A%2F%2Fmovies.zsw.ca%3A6969%2Fannounce%2Cudp%3A%2F%2Fmts.tvbit.co%3A6969%2Fannounce%2Cudp%3A%2F%2Fnagios.tks.sumy.ua%3A80%2Fannounce%2Cudp%3A%2F%2Fns-1.x-fins.com%3A6969%2Fannounce%2Cudp%3A%2F%2Fns389251.ovh.net%3A6969%2Fannounce%2Cudp%3A%2F%2Fopen.demonii.com%3A1337%2Fannounce%2Cudp%3A%2F%2Fopen.lolicon.eu%3A7777%2Fannounce%2Cudp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce%2Cudp%3A%2F%2Fopentor.org%3A2710%2Fannounce%2Cudp%3A%2F%2Fopentracker.i2p.rocks%3A6969%2Fannounce%2Cudp%3A%2F%2Fpeerfect.org%3A6969%2Fannounce%2Cudp%3A%2F%2Fpublic-tracker.zooki.xyz%3A6969%2Fannounce%2Cudp%3A%2F%2Fpublic.popcorn-tracker.org%3A6969%2Fannounce%2Cudp%3A%2F%2Fpublic.publictracker.xyz%3A6969%2Fannounce%2Cudp%3A%2F%2Fpublictracker.xyz%3A6969%2Fannounce%2Cudp%3A%2F%2Fretracker.akado-ural.ru%3A80%2Fannounce%2Cudp%3A%2F%2Fretracker.hotplug.ru%3A2710%2Fannounce%2Cudp%3A%2F%2Fretracker.lanta-net.ru%3A2710%2Fannounce%2Cudp%3A%2F%2Fretracker.netbynet.ru%3A2710%2Fannounce%2Cudp%3A%2F%2Fretracker.nts.su%3A2710%2Fannounce%2Cudp%3A%2F%2Fretracker.sevstar.net%3A2710%2Fannounce%2Cudp%3A%2F%2Frutorrent.frontline-mod.com%3A6969%2Fannounce%2Cudp%3A%2F%2Fsd-161673.dedibox.fr%3A6969%2Fannounce%2Cudp%3A%2F%2Fshadowshq.eddie4.nl%3A6969%2Fannounce%2Cudp%3A%2F%2Fshadowshq.yi.org%3A6969%2Fannounce%2Cudp%3A%2F%2Fstorage.groupees.com%3A6969%2Fannounce%2Cudp%3A%2F%2Ft1.leech.ie%3A1337%2Fannounce%2Cudp%3A%2F%2Ft2.leech.ie%3A1337%2Fannounce%2Cudp%3A%2F%2Ft3.leech.ie%3A1337%2Fannounce%2Cudp%3A%2F%2Fteamspeak.value-wolf.org%3A6969%2Fannounce%2Cudp%3A%2F%2Fthetracker.org%3A80%2Fannounce%2Cudp%3A%2F%2Ftorrent.tdjs.tech%3A6969%2Fannounce%2Cudp%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce%2Cudp%3A%2F%2Ftr.cili001.com%3A8070%2Fannounce%2Cudp%3A%2F%2Ftr2.ysagin.top%3A2710%2Fannounce%2Cudp%3A%2F%2Ftracker-udp.gbitt.info%3A80%2Fannounce%2Cudp%3A%2F%2Ftracker-v6.zooki.xyz%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.0x.tf%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.aletorrenty.pl%3A2710%2Fannounce%2Cudp%3A%2F%2Ftracker.altrosky.nl%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.army%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.beeimg.com%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.birkenwald.de%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.bittor.pw%3A1337%2Fannounce%2Cudp%3A%2F%2Ftracker.blacksparrowmedia.net%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.coppersurfer.tk%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.cyberia.is%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.dler.org%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.ds.is%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.dyne.org%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.e-utp.net%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.eddie4.nl%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.ex.ua%3A80%2Fannounce%2Cudp%3A%2F%2Ftracker.filemail.com%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.flashtorrents.org%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.fortu.io%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.grepler.com%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.internetwarriors.net%3A1337%2Fannounce%2Cudp%3A%2F%2Ftracker.kali.org%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.kamigami.org%3A2710%2Fannounce%2Cudp%3A%2F%2Ftracker.kuroy.me%3A5944%2Fannounce%2Cudp%3A%2F%2Ftracker.leechers-paradise.org%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.lelux.fi%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.moeking.me%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.open-internet.nl%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce%2Cudp%3A%2F%2Ftracker.piratepublic.com%3A1337%2Fannounce%2Cudp%3A%2F%2Ftracker.publictracker.xyz%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.shkinev.me%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.skynetcloud.site%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.skyts.net%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.swateam.org.uk%3A2710%2Fannounce%2Cudp%3A%2F%2Ftracker.tiny-vps.com%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce%2Cudp%3A%2F%2Ftracker.tvunderground.org.ru%3A3218%2Fannounce%2Cudp%3A%2F%2Ftracker.uw0.xyz%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.v6speed.org%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.vulnix.sh%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.yoshi210.com%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.zemoj.com%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.zerobytes.xyz%3A1337%2Fannounce%2Cudp%3A%2F%2Ftracker.zum.bi%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker0.ufibox.com%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce%2Cudp%3A%2F%2Ftracker2.christianbro.pw%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker2.dler.org%3A80%2Fannounce%2Cudp%3A%2F%2Ftracker2.indowebster.com%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce%2Cudp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce%2Cudp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce%2Cudp%3A%2F%2Ftracker6.dler.org%3A2710%2Fannounce%2Cudp%3A%2F%2Ftsundere.pw%3A6969%2Fannounce%2Cudp%3A%2F%2Fu.wwwww.wtf%3A1%2Fannounce%2Cudp%3A%2F%2Fultra.zt.ua%3A6969%2Fannounce%2Cudp%3A%2F%2Fus-tracker.publictracker.xyz%3A6969%2Fannounce%2Cudp%3A%2F%2Fvalakas.rollo.dnsabr.com%3A2710%2Fannounce%2Cudp%3A%2F%2Fvibe.community%3A6969%2Fannounce%2Cudp%3A%2F%2Fwassermann.online%3A6969%2Fannounce%2Cudp%3A%2F%2Fwww.loushao.net%3A8080%2Fannounce%2Cudp%3A%2F%2Fwww.midea123.z-media.com.cn%3A6969%2Fannounce%2Cudp%3A%2F%2Fwww.torrent.eu.org%3A451%2Fannounce%2Cudp%3A%2F%2Fzephir.monocul.us%3A6969%2Fannounce%2Cudp%3A%2F%2Fzer0day.ch%3A1337%2Fannounce%2Cudp%3A%2F%2Fzer0day.to%3A1337%2Fannounce&tr=http%3A%2F%2F104.238.198.186%3A8000%2Fannounce&tr=http%3A%2F%2F185.83.215.123%3A6969%2Fannounce&tr=http%3A%2F%2F62.210.202.61%2Fannounce&tr=http%3A%2F%2F87.110.238.140%3A6969%2Fannounce&tr=http%3A%2F%2F95.211.168.204%3A2710%2Fannounce&tr=http%3A%2F%2Fmail2.zelenaya.net%2Fannounce&tr=http%3A%2F%2Ftracker.electro-torrent.pl%2Fannounce&tr=http%3A%2F%2Ftracker.nyaa.uk%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.torrentyorg.pl%2Fannounce&tr=http%3A%2F%2Ftracker01.loveapp.com%3A6789%2Fannounce&tr=https%3A%2F%2Ftracker.nanoha.org%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.com%3A6869%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:eee41931ce6b536fe72f8af5fa1ffd8f8c56e588&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">1.7GB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">14</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">8</span></td>
			<td nowrap="nowrap" align="center">65</td>
			<td align="center"><a href="/topics/list/user_id/743372">黑暗路基艾尔</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 13:32			<span style="display: none;">2022/02/21 13:32</span></td>
			<td width="6%" align="center">
			<a class="sort-1" 
				href="/topics/list/sort_id/1">
				<font color=black>其他</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/550" >
				萝莉社活动室</a></span>				
				<a href="/topics/view/593384_Pixiv_282_500P.html"  target="_blank" >
				[獸耳娘噠萌進化·Pixiv][第282期][500P]|ω•)</a>
				<span style="color: gray;">約1條評論</span>			</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:DQKGXNS6NSTEUROHGCUTYA3NU66ZA7LW&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Fanidex.moe%3A6969%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=http%3A%2F%2Fopen.acgnxtracker.com%3A80%2Fannounce&tr=http%3A%2F%2Fsukebei.tracker.wf%3A8888%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%3A80%2Fannounce&tr=http%3A%2F%2Ft.overflow.biz%3A6969%2Fannounce&tr=http%3A%2F%2Ftorrenttracker.nwc.acsalaska.net%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.anirena.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.dler.org%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.gbitt.info%2Fannounce&tr=http%3A%2F%2Ftracker.internetwarriors.net%3A1337%2Fannounce&tr=http%3A%2F%2Ftracker.lelux.fi%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&tr=http%3A%2F%2Ftracker.zerobytes.xyz%3A1337%2Fannounce&tr=http%3A%2F%2Ftracker2.dler.org%3A80%2Fannounce&tr=http%3A%2F%2Fvps02.net.orel.ru%3A80%2Fannounce&tr=https%3A%2F%2F1337.abcvg.info%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.lelux.fi%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.nanoha.org%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.tamersunion.org%3A443%2Fannounce&tr=wss%3A%2F%2Ftracker.openwebtorrent.com">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:1c146bb65e6ca64a45c730a93c036da7bd907d76&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">1.3GB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">8</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">1</span></td>
			<td nowrap="nowrap" align="center">18</td>
			<td align="center"><a href="/topics/list/user_id/556597">恋之星第六季</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 13:30			<span style="display: none;">2022/02/21 13:30</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/805" >
				DBD制作组</a></span>				
				<a href="/topics/view/593383_DBD-Raws_Dragon_Ball_Shen_Long_no_Densetsu_1080P_BDRip_HEVC-10bit_FLAC_MKV.html"  target="_blank" >
				[DBD-Raws][龙珠剧场版：神龙传说/Dragon Ball: Shen Long no Densetsu/ドラゴンボール 神龍の伝説][1080P][BDRip][HEVC-10bit][FLAC][MKV]</a>
				<span style="color: gray;">約3條評論</span>			</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:4545SC4N7RQQGXDYBZ26BD4WPT645NRD&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2F1337.abcvg.info%2Fannounce&tr=http%3A%2F%2F178.175.143.27%2Fannounce&tr=http%3A%2F%2F5rt.tace.ru%3A60889%2Fannounce&tr=http%3A%2F%2F78.30.254.12%3A2710%2Fannounce&tr=http%3A%2F%2F91.217.91.21%3A3218%2Fannounce&tr=http%3A%2F%2F%5B2001%3A1b10%3A1000%3A8101%3A0%3A242%3Aac11%3A2%5D%3A6969%2Fannounce&tr=http%3A%2F%2F%5B2001%3A470%3A1%3A189%3A0%3A1%3A2%3A3%5D%3A6969%2Fannounce&tr=http%3A%2F%2F%5B2a04%3Aac00%3A1%3A3dd8%3A%3A1%3A2710%5D%3A2710%2Fannounce&tr=http%3A%2F%2Faaa.army%3A8866%2Fannounce&tr=http%3A%2F%2Fatrack.pow7.com%2Fannounce&tr=http%3A%2F%2Fbobbialbano.com%3A6969%2Fannounce&tr=http%3A%2F%2Fbt.3kb.xyz%2Fannounce&tr=http%3A%2F%2Fbt.pusacg.org%3A8080%2Fannounce&tr=http%3A%2F%2Fderpyradio.net%3A6969%2Fannounce&tr=http%3A%2F%2Fexplodie.org%3A6969%2Fannounce&tr=http%3A%2F%2Fgrifon.info%2Fannounce&tr=http%3A%2F%2Fh4.trakx.nibba.trade%2Fannounce&tr=http%3A%2F%2Fjsb.by%3A8000%2Fannounce&tr=http%3A%2F%2Fnovaopcj.icu%3A10325%2Fannounce&tr=http%3A%2F%2Fns349743.ip-91-121-106.eu%2Fannounce&tr=http%3A%2F%2Fopen.acgnxtracker.com%2Fannounce&tr=http%3A%2F%2Fopen.touki.ru%2Fannounce.php&tr=http%3A%2F%2Fopentracker.i2p.rocks%3A6969%2Fannounce&tr=http%3A%2F%2Fopentracker.xyz%2Fannounce&tr=http%3A%2F%2Fpow7.com%2Fannounce&tr=http%3A%2F%2Fpublictracker.ch%3A6969%2Fannounce&tr=http%3A%2F%2Fretracker.hotplug.ru%3A2710%2Fannounce&tr=http%3A%2F%2Fretracker.sevstar.net%3A2710%2Fannounce&tr=http%3A%2F%2Frt.tace.ru%2Fannounce&tr=http%3A%2F%2Fsecure.pow7.com%2Fannounce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Ft.overflow.biz%3A6969%2Fannounce&tr=http%3A%2F%2Ft1.pow7.com%2Fannounce&tr=http%3A%2F%2Ft2.pow7.com%2Fannounce&tr=http%3A%2F%2Fthetracker.org%2Fannounce&tr=http%3A%2F%2Ftorrentsmd.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftorrenttracker.nwc.acsalaska.net%3A6969%2Fannounce&tr=http%3A%2F%2Ftr.cili001.com%3A8070%2Fannounce&tr=http%3A%2F%2Ftracker.anonwebz.xyz%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker.birkenwald.de%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.bittor.pw%3A1337%2Fannounce&tr=http%3A%2F%2Ftracker.bittorrent.am%2Fannounce&tr=http%3A%2F%2Ftracker.bt4g.com%3A2095%2Fannounce&tr=http%3A%2F%2Ftracker.bz%2Fannounce&tr=http%3A%2F%2Ftracker.dler.org%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.dutchtracking.nl%2Fannounce&tr=http%3A%2F%2Ftracker.files.fm%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.gbitt.info%2Fannounce&tr=http%3A%2F%2Ftracker.internetwarriors.net%3A1337%2Fannounce&tr=http%3A%2F%2Ftracker.ipv6tracker.ru%2Fannounce&tr=http%3A%2F%2Ftracker.kamigami.org%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.kuroy.me%3A5944%2Fannounce&tr=http%3A%2F%2Ftracker.lelux.fi%2Fannounce&tr=http%3A%2F%2Ftracker.moeking.me%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.noobsubs.net%2Fannounce&tr=http%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&tr=http%3A%2F%2Ftracker.sakurato.xyz%3A23333%2Fannounce&tr=http%3A%2F%2Ftracker.skyts.net%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.sloppyta.co%2Fannounce&tr=http%3A%2F%2Ftracker.trackerfix.com%2Fannounce&tr=http%3A%2F%2Ftracker.uw0.xyz%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.ygsub.com%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.zerobytes.xyz%3A1337%2Fannounce&tr=http%3A%2F%2Ftracker2.dler.org%2Fannounce&tr=http%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftrun.tom.ru%2Fannounce&tr=http%3A%2F%2Fvpn.flying-datacenter.de%3A6969%2Fannounce&tr=http%3A%2F%2Fvps02.net.orel.ru%2Fannounce&tr=http%3A%2F%2Fwww.wareztorrent.com%2Fannounce&tr=https%3A%2F%2F1337.abcvg.info%2Fannounce&tr=https%3A%2F%2F2.tracker.eu.org%2Fannounce&tr=https%3A%2F%2F3.tracker.eu.org%2Fannounce&tr=https%3A%2F%2Faaa.army%3A8866%2Fannounce&tr=https%3A%2F%2Fopen.kickasstracker.com%2Fannounce&tr=https%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=https%3A%2F%2Ftr.ready4.icu%2Fannounce&tr=https%3A%2F%2Ftracker.bt-hash.com%2Fannounce&tr=https%3A%2F%2Ftracker.gbitt.info%2Fannounce&tr=https%3A%2F%2Ftracker.hama3.net%2Fannounce&tr=https%3A%2F%2Ftracker.imgoingto.icu%2Fannounce&tr=https%3A%2F%2Ftracker.lelux.fi%2Fannounce&tr=https%3A%2F%2Ftracker.lilithraws.cf%2Fannounce&tr=https%3A%2F%2Ftracker.nitrix.me%2Fannounce&tr=https%3A%2F%2Ftracker.parrotsec.org%2Fannounce&tr=https%3A%2F%2Ftracker.sloppyta.co%2Fannounce&tr=https%3A%2F%2Ftracker.tamersunion.org%2Fannounce&tr=https%3A%2F%2Ftrakx.herokuapp.com%2Fannounce&tr=https%3A%2F%2Fw.wwwww.wtf%2Fannounce&tr=https%3A%2F%2Fwww.wareztorrent.com%2Fannounce&tr=http%3A%2F%2F1337.abcvg.info%2Fannounce%2Chttp%3A%2F%2F178.175.143.27%3A80%2Fannounce%2Chttp%3A%2F%2F5rt.tace.ru%3A60889%2Fannounce%2Chttp%3A%2F%2F78.30.254.12%3A2710%2Fannounce%2Chttp%3A%2F%2F91.217.91.21%3A3218%2Fannounce%2Chttp%3A%2F%2F%5B2001%3A1b10%3A1000%3A8101%3A0%3A242%3Aac11%3A2%5D%3A6969%2Fannounce%2Chttp%3A%2F%2F%5B2001%3A470%3A1%3A189%3A0%3A1%3A2%3A3%5D%3A6969%2Fannounce%2Chttp%3A%2F%2F%5B2a04%3Aac00%3A1%3A3dd8%3A%3A1%3A2710%5D%3A2710%2Fannounce%2Chttp%3A%2F%2Faaa.army%3A8866%2Fannounce%2Chttp%3A%2F%2Fatrack.pow7.com%3A80%2Fannounce%2Chttp%3A%2F%2Fbobbialbano.com%3A6969%2Fannounce%2Chttp%3A%2F%2Fbt.3kb.xyz%3A80%2Fannounce%2Chttp%3A%2F%2Fbt.pusacg.org%3A8080%2Fannounce%2Chttp%3A%2F%2Fderpyradio.net%3A6969%2Fannounce%2Chttp%3A%2F%2Fexplodie.org%3A6969%2Fannounce%2Chttp%3A%2F%2Fgrifon.info%3A80%2Fannounce%2Chttp%3A%2F%2Fh4.trakx.nibba.trade%3A80%2Fannounce%2Chttp%3A%2F%2Fjsb.by%3A8000%2Fannounce%2Chttp%3A%2F%2Fnovaopcj.icu%3A10325%2Fannounce%2Chttp%3A%2F%2Fns349743.ip-91-121-106.eu%3A80%2Fannounce%2Chttp%3A%2F%2Fopen.acgnxtracker.com%3A80%2Fannounce%2Chttp%3A%2F%2Fopen.touki.ru%3A80%2Fannounce.php%2Chttp%3A%2F%2Fopentracker.i2p.rocks%3A6969%2Fannounce%2Chttp%3A%2F%2Fopentracker.xyz%3A80%2Fannounce%2Chttp%3A%2F%2Fpow7.com%3A80%2Fannounce%2Chttp%3A%2F%2Fpublictracker.ch%3A6969%2Fannounce%2Chttp%3A%2F%2Fretracker.hotplug.ru%3A2710%2Fannounce%2Chttp%3A%2F%2Fretracker.sevstar.net%3A2710%2Fannounce%2Chttp%3A%2F%2Frt.tace.ru%3A80%2Fannounce%2Chttp%3A%2F%2Fsecure.pow7.com%3A80%2Fannounce%2Chttp%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce%2Chttp%3A%2F%2Ft.acg.rip%3A6699%2Fannounce%2Chttp%3A%2F%2Ft.nyaatracker.com%3A80%2Fannounce%2Chttp%3A%2F%2Ft.overflow.biz%3A6969%2Fannounce%2Chttp%3A%2F%2Ft1.pow7.com%3A80%2Fannounce%2Chttp%3A%2F%2Ft2.pow7.com%3A80%2Fannounce%2Chttp%3A%2F%2Fthetracker.org%3A80%2Fannounce%2Chttp%3A%2F%2Ftorrentsmd.com%3A8080%2Fannounce%2Chttp%3A%2F%2Ftorrenttracker.nwc.acsalaska.net%3A6969%2Fannounce%2Chttp%3A%2F%2Ftr.cili001.com%3A8070%2Fannounce%2Chttp%3A%2F%2Ftracker.anonwebz.xyz%3A8080%2Fannounce%2Chttp%3A%2F%2Ftracker.birkenwald.de%3A6969%2Fannounce%2Chttp%3A%2F%2Ftracker.bittor.pw%3A1337%2Fannounce%2Chttp%3A%2F%2Ftracker.bittorrent.am%3A80%2Fannounce%2Chttp%3A%2F%2Ftracker.bt4g.com%3A2095%2Fannounce%2Chttp%3A%2F%2Ftracker.bz%3A80%2Fannounce%2Chttp%3A%2F%2Ftracker.dler.org%3A6969%2Fannounce%2Chttp%3A%2F%2Ftracker.dutchtracking.nl%3A80%2Fannounce%2Chttp%3A%2F%2Ftracker.files.fm%3A6969%2Fannounce%2Chttp%3A%2F%2Ftracker.gbitt.info%3A80%2Fannounce%2Chttp%3A%2F%2Ftracker.internetwarriors.net%3A1337%2Fannounce%2Chttp%3A%2F%2Ftracker.ipv6tracker.ru%3A80%2Fannounce%2Chttp%3A%2F%2Ftracker.kamigami.org%3A2710%2Fannounce%2Chttp%3A%2F%2Ftracker.kuroy.me%3A5944%2Fannounce%2Chttp%3A%2F%2Ftracker.lelux.fi%3A80%2Fannounce%2Chttp%3A%2F%2Ftracker.moeking.me%3A6969%2Fannounce%2Chttp%3A%2F%2Ftracker.noobsubs.net%3A80%2Fannounce%2Chttp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce%2Chttp%3A%2F%2Ftracker.sakurato.xyz%3A23333%2Fannounce%2Chttp%3A%2F%2Ftracker.skyts.net%3A6969%2Fannounce%2Chttp%3A%2F%2Ftracker.sloppyta.co%3A80%2Fannounce%2Chttp%3A%2F%2Ftracker.trackerfix.com%3A80%2Fannounce%2Chttp%3A%2F%2Ftracker.uw0.xyz%3A6969%2Fannounce%2Chttp%3A%2F%2Ftracker.ygsub.com%3A6969%2Fannounce%2Chttp%3A%2F%2Ftracker.zerobytes.xyz%3A1337%2Fannounce%2Chttp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce%2Chttp%3A%2F%2Ftracker2.dler.org%3A80%2Fannounce%2Chttp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce%2Chttp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce%2Chttp%3A%2F%2Ftrun.tom.ru%3A80%2Fannounce%2Chttp%3A%2F%2Fvpn.flying-datacenter.de%3A6969%2Fannounce%2Chttp%3A%2F%2Fvps02.net.orel.ru%3A80%2Fannounce%2Chttp%3A%2F%2Fwww.wareztorrent.com%3A80%2Fannounce%2Chttps%3A%2F%2F1337.abcvg.info%3A443%2Fannounce%2Chttps%3A%2F%2F2.tracker.eu.org%3A443%2Fannounce%2Chttps%3A%2F%2F3.tracker.eu.org%3A443%2Fannounce%2Chttps%3A%2F%2Faaa.army%3A8866%2Fannounce%2Chttps%3A%2F%2Fopen.kickasstracker.com%3A443%2Fannounce%2Chttps%3A%2F%2Fopentracker.acgnx.se%3A443%2Fannounce%2Chttps%3A%2F%2Ftr.ready4.icu%3A443%2Fannounce%2Chttps%3A%2F%2Ftracker.bt-hash.com%3A443%2Fannounce%2Chttps%3A%2F%2Ftracker.gbitt.info%3A443%2Fannounce%2Chttps%3A%2F%2Ftracker.hama3.net%3A443%2Fannounce%2Chttps%3A%2F%2Ftracker.imgoingto.icu%3A443%2Fannounce%2Chttps%3A%2F%2Ftracker.lelux.fi%3A443%2Fannounce%2Chttps%3A%2F%2Ftracker.lilithraws.cf%3A443%2Fannounce%2Chttps%3A%2F%2Ftracker.nitrix.me%3A443%2Fannounce%2Chttps%3A%2F%2Ftracker.parrotsec.org%3A443%2Fannounce%2Chttps%3A%2F%2Ftracker.sloppyta.co%3A443%2Fannounce%2Chttps%3A%2F%2Ftracker.tamersunion.org%3A443%2Fannounce%2Chttps%3A%2F%2Ftrakx.herokuapp.com%3A443%2Fannounce%2Chttps%3A%2F%2Fw.wwwww.wtf%3A443%2Fannounce%2Chttps%3A%2F%2Fwww.wareztorrent.com%3A443%2Fannounce%2Cudp%3A%2F%2F104.238.198.186%3A8000%2Fannounce%2Cudp%3A%2F%2F107.150.14.110%3A6969%2Fannounce%2Cudp%3A%2F%2F109.121.134.121%3A1337%2Fannounce%2Cudp%3A%2F%2F114.55.113.60%3A6969%2Fannounce%2Cudp%3A%2F%2F128.199.70.66%3A5944%2Fannounce%2Cudp%3A%2F%2F151.80.120.114%3A2710%2Fannounce%2Cudp%3A%2F%2F168.235.67.63%3A6969%2Fannounce%2Cudp%3A%2F%2F178.33.73.26%3A2710%2Fannounce%2Cudp%3A%2F%2F182.176.139.129%3A6969%2Fannounce%2Cudp%3A%2F%2F185.5.97.139%3A8089%2Fannounce%2Cudp%3A%2F%2F185.83.215.123%3A6969%2Fannounce%2Cudp%3A%2F%2F185.86.149.205%3A1337%2Fannounce%2Cudp%3A%2F%2F188.165.253.109%3A1337%2Fannounce%2Cudp%3A%2F%2F191.101.229.236%3A1337%2Fannounce%2Cudp%3A%2F%2F194.106.216.222%3A80%2Fannounce%2Cudp%3A%2F%2F195.123.209.37%3A1337%2Fannounce%2Cudp%3A%2F%2F195.123.209.40%3A80%2Fannounce%2Cudp%3A%2F%2F208.67.16.113%3A8000%2Fannounce%2Cudp%3A%2F%2F212.1.226.176%3A2710%2Fannounce%2Cudp%3A%2F%2F212.47.227.58%3A6969%2Fannounce%2Cudp%3A%2F%2F213.163.67.56%3A1337%2Fannounce%2Cudp%3A%2F%2F37.19.5.155%3A2710%2Fannounce%2Cudp%3A%2F%2F3rt.tace.ru%3A60889%2Fannounce%2Cudp%3A%2F%2F46.4.109.148%3A6969%2Fannounce%2Cudp%3A%2F%2F47.ip-51-68-199.eu%3A6969%2Fannounce%2Cudp%3A%2F%2F5.79.249.77%3A6969%2Fannounce%2Cudp%3A%2F%2F5.79.83.193%3A6969%2Fannounce%2Cudp%3A%2F%2F51.254.244.161%3A6969%2Fannounce%2Cudp%3A%2F%2F52.58.128.163%3A6969%2Fannounce%2Cudp%3A%2F%2F61626c.net%3A6969%2Fannounce%2Cudp%3A%2F%2F62.138.0.158%3A6969%2Fannounce%2Cudp%3A%2F%2F62.212.85.66%3A2710%2Fannounce%2Cudp%3A%2F%2F74.82.52.209%3A6969%2Fannounce%2Cudp%3A%2F%2F78.30.254.12%3A2710%2Fannounce%2Cudp%3A%2F%2F85.17.19.180%3A80%2Fannounce%2Cudp%3A%2F%2F89.234.156.205%3A80%2Fannounce%2Cudp%3A%2F%2F9.rarbg.com%3A2710%2Fannounce%2Cudp%3A%2F%2F9.rarbg.me%3A2710%2Fannounce%2Cudp%3A%2F%2F9.rarbg.me%3A2780%2Fannounce%2Cudp%3A%2F%2F9.rarbg.to%3A2710%2Fannounce%2Cudp%3A%2F%2F9.rarbg.to%3A2730%2Fannounce%2Cudp%3A%2F%2F91.216.110.52%3A451%2Fannounce%2Cudp%3A%2F%2F91.218.230.81%3A6969%2Fannounce%2Cudp%3A%2F%2F94.23.183.33%3A6969%2Fannounce%2Cudp%3A%2F%2F95.211.168.204%3A2710%2Fannounce%2Cudp%3A%2F%2F%5B2001%3A1b10%3A1000%3A8101%3A0%3A242%3Aac11%3A2%5D%3A6969%2Fannounce%2Cudp%3A%2F%2F%5B2001%3A470%3A1%3A189%3A0%3A1%3A2%3A3%5D%3A6969%2Fannounce%2Cudp%3A%2F%2F%5B2a03%3A7220%3A8083%3Acd00%3A%3A1%5D%3A451%2Fannounce%2Cudp%3A%2F%2F%5B2a04%3Aac00%3A1%3A3dd8%3A%3A1%3A2710%5D%3A2710%2Fannounce%2Cudp%3A%2F%2F%5B2a04%3Ac44%3Ae00%3A32e0%3A4cf%3A6aff%3Afe00%3Aaa1%5D%3A6969%2Fannounce%2Cudp%3A%2F%2Faaa.army%3A8866%2Fannounce%2Cudp%3A%2F%2Fadm.category5.tv%3A6969%2Fannounce%2Cudp%3A%2F%2Fadmin.videoenpoche.info%3A6969%2Fannounce%2Cudp%3A%2F%2Fadminion.n-blade.ru%3A6969%2Fannounce%2Cudp%3A%2F%2Fanidex.moe%3A6969%2Fannounce%2Cudp%3A%2F%2Fapi.bitumconference.ru%3A6969%2Fannounce%2Cudp%3A%2F%2Faruacfilmes.com.br%3A6969%2Fannounce%2Cudp%3A%2F%2Fasger.xyz%3A6969%2Fannounce%2Cudp%3A%2F%2Fbclearning.top%3A6969%2Fannounce%2Cudp%3A%2F%2Fbenouworldtrip.fr%3A6969%2Fannounce%2Cudp%3A%2F%2Fbioquantum.co.za%3A6969%2Fannounce%2Cudp%3A%2F%2Fbitsparadise.info%3A6969%2Fannounce%2Cudp%3A%2F%2Fblokas.io%3A6969%2Fannounce%2Cudp%3A%2F%2Fbms-hosxp.com%3A6969%2Fannounce%2Cudp%3A%2F%2Fbt.firebit.org%3A2710%2Fannounce%2Cudp%3A%2F%2Fbt.okmp3.ru%3A2710%2Fannounce%2Cudp%3A%2F%2Fbt1.archive.org%3A6969%2Fannounce%2Cudp%3A%2F%2Fbt2.54new.com%3A8080%2Fannounce%2Cudp%3A%2F%2Fbt2.archive.org%3A6969%2Fannounce%2Cudp%3A%2F%2Fbubu.mapfactor.com%3A6969%2Fannounce%2Cudp%3A%2F%2Fcamera.lei001.com%3A6969%2Fannounce%2Cudp%3A%2F%2Fcdn-1.gamecoast.org%3A6969%2Fannounce%2Cudp%3A%2F%2Fcdn-2.gamecoast.org%3A6969%2Fannounce%2Cudp%3A%2F%2Fchanchan.uchuu.co.uk%3A6969%2Fannounce%2Cudp%3A%2F%2Fchihaya.toss.li%3A9696%2Fannounce%2Cudp%3A%2F%2Fcode2chicken.nl%3A6969%2Fannounce%2Cudp%3A%2F%2Fconcen.org%3A6969%2Fannounce%2Cudp%3A%2F%2Fcutiegirl.ru%3A6969%2Fannounce%2Cudp%3A%2F%2Fdaveking.com%3A6969%2Fannounce%2Cudp%3A%2F%2Fdenis.stalker.upeer.me%3A6969%2Fannounce%2Cudp%3A%2F%2Fdiscord.heihachi.pw%3A6969%2Fannounce%2Cudp%3A%2F%2Fdpiui.reedlan.com%3A6969%2Fannounce%2Cudp%3A%2F%2Fdrumkitx.com%3A6969%2Fannounce%2Cudp%3A%2F%2Feddie4.nl%3A6969%2Fannounce%2Cudp%3A%2F%2Fedu.uifr.ru%3A6969%2Fannounce%2Cudp%3A%2F%2Fengplus.ru%3A6969%2Fannounce%2Cudp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce%2Cudp%3A%2F%2Ffe.dealclub.de%3A6969%2Fannounce%2Cudp%3A%2F%2Fforever-tracker.zooki.xyz%3A6969%2Fannounce%2Cudp%3A%2F%2Ffree-tracker.zooki.xyz%3A6969%2Fannounce%2Cudp%3A%2F%2Fgit.vulnix.sh%3A6969%2Fannounce%2Cudp%3A%2F%2Finferno.demonoid.is%3A3391%2Fannounce%2Cudp%3A%2F%2Fipv4.tracker.harry.lu%3A80%2Fannounce%2Cudp%3A%2F%2Fipv6.tracker.zerobytes.xyz%3A16661%2Fannounce%2Cudp%3A%2F%2Fjsb.by%3A8000%2Fannounce%2Cudp%3A%2F%2Fkanal-4.de%3A6969%2Fannounce%2Cudp%3A%2F%2Fkoli.services%3A6969%2Fannounce%2Cudp%3A%2F%2Fline-net.ru%3A6969%2Fannounce%2Cudp%3A%2F%2Fln.mtahost.co%3A6969%2Fannounce%2Cudp%3A%2F%2Fmail.realliferpg.de%3A6969%2Fannounce%2Cudp%3A%2F%2Fmgtracker.org%3A2710%2Fannounce%2Cudp%3A%2F%2Fmovies.zsw.ca%3A6969%2Fannounce%2Cudp%3A%2F%2Fmts.tvbit.co%3A6969%2Fannounce%2Cudp%3A%2F%2Fnagios.tks.sumy.ua%3A80%2Fannounce%2Cudp%3A%2F%2Fns-1.x-fins.com%3A6969%2Fannounce%2Cudp%3A%2F%2Fns389251.ovh.net%3A6969%2Fannounce%2Cudp%3A%2F%2Fopen.demonii.com%3A1337%2Fannounce%2Cudp%3A%2F%2Fopen.lolicon.eu%3A7777%2Fannounce%2Cudp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce%2Cudp%3A%2F%2Fopentor.org%3A2710%2Fannounce%2Cudp%3A%2F%2Fopentracker.i2p.rocks%3A6969%2Fannounce%2Cudp%3A%2F%2Fpeerfect.org%3A6969%2Fannounce%2Cudp%3A%2F%2Fpublic-tracker.zooki.xyz%3A6969%2Fannounce%2Cudp%3A%2F%2Fpublic.popcorn-tracker.org%3A6969%2Fannounce%2Cudp%3A%2F%2Fpublic.publictracker.xyz%3A6969%2Fannounce%2Cudp%3A%2F%2Fpublictracker.xyz%3A6969%2Fannounce%2Cudp%3A%2F%2Fretracker.akado-ural.ru%3A80%2Fannounce%2Cudp%3A%2F%2Fretracker.hotplug.ru%3A2710%2Fannounce%2Cudp%3A%2F%2Fretracker.lanta-net.ru%3A2710%2Fannounce%2Cudp%3A%2F%2Fretracker.netbynet.ru%3A2710%2Fannounce%2Cudp%3A%2F%2Fretracker.nts.su%3A2710%2Fannounce%2Cudp%3A%2F%2Fretracker.sevstar.net%3A2710%2Fannounce%2Cudp%3A%2F%2Frutorrent.frontline-mod.com%3A6969%2Fannounce%2Cudp%3A%2F%2Fsd-161673.dedibox.fr%3A6969%2Fannounce%2Cudp%3A%2F%2Fshadowshq.eddie4.nl%3A6969%2Fannounce%2Cudp%3A%2F%2Fshadowshq.yi.org%3A6969%2Fannounce%2Cudp%3A%2F%2Fstorage.groupees.com%3A6969%2Fannounce%2Cudp%3A%2F%2Ft1.leech.ie%3A1337%2Fannounce%2Cudp%3A%2F%2Ft2.leech.ie%3A1337%2Fannounce%2Cudp%3A%2F%2Ft3.leech.ie%3A1337%2Fannounce%2Cudp%3A%2F%2Fteamspeak.value-wolf.org%3A6969%2Fannounce%2Cudp%3A%2F%2Fthetracker.org%3A80%2Fannounce%2Cudp%3A%2F%2Ftorrent.tdjs.tech%3A6969%2Fannounce%2Cudp%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce%2Cudp%3A%2F%2Ftr.cili001.com%3A8070%2Fannounce%2Cudp%3A%2F%2Ftr2.ysagin.top%3A2710%2Fannounce%2Cudp%3A%2F%2Ftracker-udp.gbitt.info%3A80%2Fannounce%2Cudp%3A%2F%2Ftracker-v6.zooki.xyz%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.0x.tf%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.aletorrenty.pl%3A2710%2Fannounce%2Cudp%3A%2F%2Ftracker.altrosky.nl%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.army%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.beeimg.com%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.birkenwald.de%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.bittor.pw%3A1337%2Fannounce%2Cudp%3A%2F%2Ftracker.blacksparrowmedia.net%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.coppersurfer.tk%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.cyberia.is%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.dler.org%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.ds.is%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.dyne.org%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.e-utp.net%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.eddie4.nl%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.ex.ua%3A80%2Fannounce%2Cudp%3A%2F%2Ftracker.filemail.com%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.flashtorrents.org%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.fortu.io%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.grepler.com%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.internetwarriors.net%3A1337%2Fannounce%2Cudp%3A%2F%2Ftracker.kali.org%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.kamigami.org%3A2710%2Fannounce%2Cudp%3A%2F%2Ftracker.kuroy.me%3A5944%2Fannounce%2Cudp%3A%2F%2Ftracker.leechers-paradise.org%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.lelux.fi%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.moeking.me%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.open-internet.nl%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce%2Cudp%3A%2F%2Ftracker.piratepublic.com%3A1337%2Fannounce%2Cudp%3A%2F%2Ftracker.publictracker.xyz%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.shkinev.me%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.skynetcloud.site%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.skyts.net%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.swateam.org.uk%3A2710%2Fannounce%2Cudp%3A%2F%2Ftracker.tiny-vps.com%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce%2Cudp%3A%2F%2Ftracker.tvunderground.org.ru%3A3218%2Fannounce%2Cudp%3A%2F%2Ftracker.uw0.xyz%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.v6speed.org%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.vulnix.sh%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.yoshi210.com%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.zemoj.com%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker.zerobytes.xyz%3A1337%2Fannounce%2Cudp%3A%2F%2Ftracker.zum.bi%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker0.ufibox.com%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce%2Cudp%3A%2F%2Ftracker2.christianbro.pw%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker2.dler.org%3A80%2Fannounce%2Cudp%3A%2F%2Ftracker2.indowebster.com%3A6969%2Fannounce%2Cudp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce%2Cudp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce%2Cudp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce%2Cudp%3A%2F%2Ftracker6.dler.org%3A2710%2Fannounce%2Cudp%3A%2F%2Ftsundere.pw%3A6969%2Fannounce%2Cudp%3A%2F%2Fu.wwwww.wtf%3A1%2Fannounce%2Cudp%3A%2F%2Fultra.zt.ua%3A6969%2Fannounce%2Cudp%3A%2F%2Fus-tracker.publictracker.xyz%3A6969%2Fannounce%2Cudp%3A%2F%2Fvalakas.rollo.dnsabr.com%3A2710%2Fannounce%2Cudp%3A%2F%2Fvibe.community%3A6969%2Fannounce%2Cudp%3A%2F%2Fwassermann.online%3A6969%2Fannounce%2Cudp%3A%2F%2Fwww.loushao.net%3A8080%2Fannounce%2Cudp%3A%2F%2Fwww.midea123.z-media.com.cn%3A6969%2Fannounce%2Cudp%3A%2F%2Fwww.torrent.eu.org%3A451%2Fannounce%2Cudp%3A%2F%2Fzephir.monocul.us%3A6969%2Fannounce%2Cudp%3A%2F%2Fzer0day.ch%3A1337%2Fannounce%2Cudp%3A%2F%2Fzer0day.to%3A1337%2Fannounce&tr=http%3A%2F%2F104.238.198.186%3A8000%2Fannounce&tr=http%3A%2F%2F185.83.215.123%3A6969%2Fannounce&tr=http%3A%2F%2F62.210.202.61%2Fannounce&tr=http%3A%2F%2F87.110.238.140%3A6969%2Fannounce&tr=http%3A%2F%2F95.211.168.204%3A2710%2Fannounce&tr=http%3A%2F%2Fmail2.zelenaya.net%2Fannounce&tr=http%3A%2F%2Ftracker.electro-torrent.pl%2Fannounce&tr=http%3A%2F%2Ftracker.nyaa.uk%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.torrentyorg.pl%2Fannounce&tr=http%3A%2F%2Ftracker01.loveapp.com%3A6789%2Fannounce&tr=https%3A%2F%2Ftracker.nanoha.org%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.com%3A6869%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:e779d90b8dfc61035c780e75e08f967cfdceb623&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">1.8GB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">26</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">34</span></td>
			<td nowrap="nowrap" align="center">41</td>
			<td align="center"><a href="/topics/list/user_id/743372">黑暗路基艾尔</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 13:24			<span style="display: none;">2022/02/21 13:24</span></td>
			<td width="6%" align="center">
			<a class="sort-12" 
				href="/topics/list/sort_id/12">
				<font color=brown>特攝</font></a>
			</td>
			<td class="title">
								
								
				<a href="/topics/view/593382_KRSUB_Revice_08_1080P.html"  target="_blank" >
				[KRSUB][假面骑士Revice][08][全家的假日 天国与地狱！？][1080P][官方中文][外挂字幕]</a>
				<span style="color: gray;">約1條評論</span>			</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:7K5L7Y5I2VUOB6OXKG5PW5L6BQ6XDR72&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2F104.238.198.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Fopen.nyaatorrents.info%3A6544%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=https%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.com%3A6869%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce&tr=https%3A%2F%2Ftr.bangumi.moe%3A9696%2Fannounce&tr=http%3A%2F%2F94.228.192.98%2Fannounce&tr=http%3A%2F%2Ftracker.btcake.com%2Fannounce&tr=http%3A%2F%2Ftracker.ktxp.com%3A6868%2Fannounce&tr=http%3A%2F%2Ftracker.ktxp.com%3A7070%2Fannounce&tr=http%3A%2F%2Fbt.sc-ol.com%3A2710%2Fannounce&tr=http%3A%2F%2Fbtfile.sdo.com%3A6961%2Fannounce&tr=http%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker3.torrentino.com%2Fannounce&tr=http%3A%2F%2Ftracker2.torrentino.com%2Fannounce&tr=http%3A%2F%2Fpubt.net%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.tfile.me%2Fannounce&tr=http%3A%2F%2Fbigfoot1942.sektori.org%3A6969%2Fannounce&tr=http%3A%2F%2F208.67.16.113%3A8000%2Fannonuce&tr=https%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.kamigami.org%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.kisssub.org%3A2015%2Fannounce&tr=http%3A%2F%2Fsukebei.tracker.wf%3A8888%2Fannounce&tr=https%3A%2F%2F1337.abcvg.info%2Fannounce&tr=https%3A%2F%2Ftracker.fastdownload.xyz%2Fannounce&tr=https%3A%2F%2Fopentracker.xyz%2Fannounce&tr=http%3A%2F%2Fopen.trackerlist.xyz%2Fannounce&tr=https%3A%2F%2Ft.quic.ws%2Fannounce&tr=https%3A%2F%2Ftracker.gbitt.info%2Fannounce&tr=http%3A%2F%2Fvps02.net.orel.ru%2Fannounce&tr=http%3A%2F%2Ftracker2.dler.org%2Fannounce&tr=https%3A%2F%2Ftracker.vectahosting.eu%3A2053%2Fannounce&tr=https%3A%2F%2Ftracker.parrotsec.org%2Fannounce&tr=http%3A%2F%2Fgwp2-v19.rinet.ru%2Fannounce&tr=http%3A%2F%2Ftracker.files.fm%3A6969%2Fannounce&tr=http%3A%2F%2Fbt1.xxxxbt.cc%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2F93.158.213.92%3A1337%2Fannounce&tr=http%3A%2F%2F176.113.71.19%3A6961%2Fannounce&tr=http%3A%2F%2F184.105.151.164%3A6969%2Fannounce&tr=http%3A%2F%2F78.30.254.12%3A2710%2Fannounce&tr=http%3A%2F%2F86.62.124.78%2Fannounce&tr=http%3A%2F%2F95.107.48.115%2Fannounce&tr=http%3A%2F%2F193.148.69.107%2Fannounce&tr=http%3A%2F%2F54.39.98.124%2Fannounce&tr=http%3A%2F%2F182.150.53.61%3A8080%2Fannounce&tr=http%3A%2F%2F175.24.22.206%3A11450%2Fannounce&tr=http%3A%2F%2F13.70.4.194%3A6969%2Fannounce&tr=http%3A%2F%2Ftr.cili001.com%3A8070%2Fannounce&tr=https%3A%2F%2Ftracker.nanoha.org%2Fannounce&tr=https%3A%2F%2Ftracker.nitrix.me%2Fannounce&tr=http%3A%2F%2Ftracker.nyap2p.com%3A8080%2Fannounce&tr=http%3A%2F%2Fh4.trakx.nibba.trade%2Fannounce&tr=http%3A%2F%2Ftracker.gbitt.info%2Fannounce&tr=http%3A%2F%2Ftracker810.xyz%3A11450%2Fannounce&tr=http%3A%2F%2Fopen.acgnxtracker.com%2Fannounce&tr=http%3A%2F%2F51.38.230.101%2Fannounce&tr=http%3A%2F%2F176.113.71.60%3A6961%2Fannounce&tr=http%3A%2F%2F176.113.68.67%3A6961%2Fannounce&tr=http%3A%2F%2F51.15.55.204%3A1337%2Fannounce&tr=http%3A%2F%2F78.46.225.225%3A2710%2Fannounce&tr=http%3A%2F%2F185.83.214.123%3A6969%2Fannounce&tr=http%3A%2F%2F91.207.136.85%2Fannounce&tr=https%3A%2F%2Ftracker.nyaa.tk%2Fannounce&tr=http%3A%2F%2Ftracker.zerobytes.xyz%3A1337%2Fannounce&tr=http%3A%2F%2Ftracker.internetwarriors.net%3A1337%2Fannounce&tr=http%3A%2F%2Ftracker.btsync.gq%3A233%2Fannounce&tr=http%3A%2F%2Fmail2.zelenaya.net%2Fannounce&tr=http%3A%2F%2Fexplodie.org%3A6969%2Fannounce&tr=https%3A%2F%2Ftracker.parrotlinux.org%2Fannounce&tr=http%3A%2F%2Ftracker.skyts.net%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&tr=http%3A%2F%2Ftracker.bz%2Fannounce&tr=http%3A%2F%2Ftracker.bt4g.com%3A2095%2Fannounce&tr=http%3A%2F%2Fretracker.sevstar.net%3A2710%2Fannounce&tr=http%3A%2F%2Fpow7.com%2Fannounce&tr=https%3A%2F%2Ftracker.sloppyta.co%2Fannounce&tr=https%3A%2F%2Ftracker.lelux.fi%2Fannounce&tr=http%3A%2F%2Fwww.loushao.net%3A8080%2Fannounce&tr=http%3A%2F%2Ftrun.tom.ru%2Fannounce&tr=http%3A%2F%2Ftracker.yoshi210.com%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.ygsub.com%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.lelux.fi%2Fannounce&tr=http%3A%2F%2Ftracker.gcvchp.com%3A2710%2Fannounce&tr=http%3A%2F%2Fp4p.arenabg.com%3A1337%2Fannounce&tr=http%3A%2F%2Fopentracker.i2p.rocks%3A6969%2Fannounce&tr=wss%3A%2F%2Ftracker.sloppyta.co%3A443%2Fannounce&tr=wss%3A%2F%2Ftracker.openwebtorrent.com%3A443%2Fannounce&tr=wss%3A%2F%2Ftracker.novage.com.ua%3A443%2Fannounce&tr=http%3A%2F%2Ftracker.ipv6tracker.ru%2Fannounce&tr=https%3A%2F%2Ftracker6.lelux.fi%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:fababfe3a8d568e0f9d751bafb757e0c3d71c7fa&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">706.2MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">3</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">1</span></td>
			<td nowrap="nowrap" align="center">11</td>
			<td align="center"><a href="/topics/list/user_id/260865">_、未来、</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 13:10			<span style="display: none;">2022/02/21 13:10</span></td>
			<td width="6%" align="center">
			<a class="sort-6" 
				href="/topics/list/sort_id/6">
				<font color=blue>日劇</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/648" >
				魔星字幕团</a></span>				
				<a href="/topics/view/593381_MagicStar_EP06_WEBDL_1080p.html"  target="_blank" >
				[MagicStar] 鹿枫堂四色日和 / 鹿楓堂よついろ日和 EP06 [WEBDL] [1080p]【生】【附日字】</a>
				<span style="color: gray;">約1條評論</span>			</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:3ICGVTPNJNQ57Q53MXL4RIQVTUQTXYSY&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:da046acded4b61dfc3bb65d7c8a2159d213be258&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">490.2MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">1</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">0</span></td>
			<td nowrap="nowrap" align="center">5</td>
			<td align="center"><a href="/topics/list/user_id/105217">jackieliiy</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 13:10			<span style="display: none;">2022/02/21 13:10</span></td>
			<td width="6%" align="center">
			<a class="sort-6" 
				href="/topics/list/sort_id/6">
				<font color=blue>日劇</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/648" >
				魔星字幕团</a></span>				
				<a href="/topics/view/593380_MagicStar_EP06_WEBDL_1080p.html"  target="_blank" >
				[MagicStar] 如果，有一所只有帅哥的高中 / もしも、イケメンだけの高校があったら EP06 [WEBDL] [1080p]【生】【附日字】</a>
				<span style="color: gray;">約1條評論</span>			</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:2SK52ZYVUJNCA5JIWJHTB7Q235ETTFEJ&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:d495dd6715a25a207528b24f30fe1adf49399489&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">710.7MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">0</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">3</span></td>
			<td nowrap="nowrap" align="center">5</td>
			<td align="center"><a href="/topics/list/user_id/105217">jackieliiy</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 11:38			<span style="display: none;">2022/02/21 11:38</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/657" >
				LoliHouse</a></span>				
				<a href="/topics/view/593379_LoliHouse_Shingeki_no_Kyojin_-_82_WebRip_1080p_HEVC-10bit_AAC.html"  target="_blank" >
				[豌豆字幕组&LoliHouse] 进击的巨人 / Shingeki no Kyojin - 82 [WebRip 1080p HEVC-10bit AAC][简繁内封字幕]</a>
				<span style="color: gray;">約3條評論</span>			</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:OQ7B3OT3IZOEIMNBGCUJ6JHGXPD22JQJ&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Fopen.nyaatorrents.info%3A6544%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%3A80%2Fannounce&tr=http%3A%2F%2F104.238.198.186%3A8000%2Fannounce&tr=https%3A%2F%2Ftr.bangumi.moe%3A9696%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce&tr=http%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Fopen.acgnxtracker.com%2Fannounce&tr=http%3A%2F%2Fsukebei.tracker.wf%3A8888%2Fannounce&tr=http%3A%2F%2Ftracker.kamigami.org%3A2710%2Fannounce&tr=https%3A%2F%2Ftracker.nanoha.org%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:743e1dba7b465c4431a130a89f24e6bbc7ad2609&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">498MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">107</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">80</span></td>
			<td nowrap="nowrap" align="center">411</td>
			<td align="center"><a href="/topics/list/user_id/689653">LoliHouse</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 11:00			<span style="display: none;">2022/02/21 11:00</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/801" >
				NC-Raws</a></span>				
				<a href="/topics/view/593378_NC-Raws_WEB_Ore_Tsushima_-_88_Baha_1920x1080_AVC_AAC_MP4.html"  target="_blank" >
				[NC-Raws] 叫我對大哥 (WEB版) / Ore, Tsushima - 88 (Baha 1920x1080 AVC AAC MP4)</a>
				<span style="color: gray;">約3條評論</span>			</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:OM7V4VHYIBWUZ4CJLZOTADUIKWXEE3IV&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Fopen.nyaatorrents.info%3A6544%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=https%3A%2F%2Ftr.bangumi.moe%3A9696%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:733f5e54f8406d4cf0495e5d300e8855ae426d15&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">24MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">13</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">2</span></td>
			<td nowrap="nowrap" align="center">50</td>
			<td align="center"><a href="/topics/list/user_id/637871">九十九朔夜</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 10:45			<span style="display: none;">2022/02/21 10:45</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/669" >
				喵萌奶茶屋</a></span>				
				<a href="/topics/view/593377_01_Sono_Bisque_Doll_wa_Koi_wo_Suru_07_1080p.html"  target="_blank" >
				【喵萌奶茶屋】★01月新番★[更衣人偶坠入爱河/恋上换装娃娃/Sono Bisque Doll wa Koi wo Suru][07][1080p][简体][招募翻译校对]</a>
				<span style="color: gray;">約1條評論</span>			</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:FCNZDEJXIPWGOWPL6ZRYOPAOCPNPIM56&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Fopen.nyaatorrents.info%3A6544%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%3A80%2Fannounce&tr=https%3A%2F%2Fopen.acgnxtracker.com%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce&tr=http%3A%2F%2Fsukebei.tracker.wf%3A8888%2Fannounce&tr=http%3A%2F%2Fp4p.arenabg.com%3A1337%2Fannounce&tr=http%3A%2F%2Fopenbittorrent.com%3A80%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:289b91913743ec6759ebf663873c0e13daf433be&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">326.4MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">23</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">2</span></td>
			<td nowrap="nowrap" align="center">111</td>
			<td align="center"><a href="/topics/list/user_id/693094">nekomoekissaten</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 10:45			<span style="display: none;">2022/02/21 10:45</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/669" >
				喵萌奶茶屋</a></span>				
				<a href="/topics/view/593376_01_Sono_Bisque_Doll_wa_Koi_wo_Suru_07_720p.html"  target="_blank" >
				【喵萌奶茶屋】★01月新番★[更衣人偶坠入爱河/恋上换装娃娃/Sono Bisque Doll wa Koi wo Suru][07][720p][简体][招募翻译校对]</a>
				<span style="color: gray;">約1條評論</span>			</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:SVYQ4W2HQLPKBPBGAPWNU4Y6IOHDKKQM&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Fopen.nyaatorrents.info%3A6544%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%3A80%2Fannounce&tr=https%3A%2F%2Fopen.acgnxtracker.com%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce&tr=http%3A%2F%2Fsukebei.tracker.wf%3A8888%2Fannounce&tr=http%3A%2F%2Fp4p.arenabg.com%3A1337%2Fannounce&tr=http%3A%2F%2Fopenbittorrent.com%3A80%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:95710e5b4782dea0bc2603ecda731e438e352a0c&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">146.9MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">8</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">15</span></td>
			<td nowrap="nowrap" align="center">31</td>
			<td align="center"><a href="/topics/list/user_id/693094">nekomoekissaten</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 09:05			<span style="display: none;">2022/02/21 09:05</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/767" >
				天月動漫&發佈組</a></span>				
				<a href="/topics/view/593375_Skymoon-Raws_Shingeki_no_Kyojin_-_The_Final_Season_-_23_ViuTV_WEB-DL_1080p_AVC_AAC_MP4_ASSx2.html"  target="_blank" >
				[Skymoon-Raws] 進擊的巨人 第四季 / Shingeki no Kyojin - The Final Season - 23 [ViuTV][WEB-DL][1080p][AVC AAC][繁體外掛][MP4+ASSx2](正式版本)</a>
				<span style="color: gray;">約3條評論</span>			</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:PESM4NE2LFSDOEYOUKZ53WOXKXT4UQQV&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=http%3A%2F%2Fanidex.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.anirena.com%3A80%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:7924ce349a596437130ea2b3ddd9d755e7ca4215&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">407.7MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">18</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">3</span></td>
			<td nowrap="nowrap" align="center">89</td>
			<td align="center"><a href="/topics/list/user_id/730809">Laputa</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 09:04			<span style="display: none;">2022/02/21 09:04</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/767" >
				天月動漫&發佈組</a></span>				
				<a href="/topics/view/593374_Skymoon-Raws_Shingeki_no_Kyojin_-_The_Final_Season_-_23_ViuTV_WEB-RIP_1080p_AVC_AAC_CHT_SRTx2_MKV.html"  target="_blank" >
				[Skymoon-Raws] 進擊的巨人 第四季 / Shingeki no Kyojin - The Final Season - 23 [ViuTV][WEB-RIP][1080p][AVC AAC][CHT][SRTx2][MKV](先行版本)</a>
							</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:HVWIMFAQAOBCDMGA3XCXUQR3SGYWNEK3&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=http%3A%2F%2Fanidex.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.anirena.com%3A80%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:3d6c861410038221b0c0ddc57a423b91b166915b&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">392.5MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">6</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">0</span></td>
			<td nowrap="nowrap" align="center">15</td>
			<td align="center"><a href="/topics/list/user_id/730809">Laputa</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 08:07			<span style="display: none;">2022/02/21 08:07</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/767" >
				天月動漫&發佈組</a></span>				
				<a href="/topics/view/593373_Sasaki_to_Miyano_-_07_1080P.html"  target="_blank" >
				[天月搬運組] 佐佐木與宮野 / Sasaki to Miyano - 07 [1080P][簡繁日外掛]</a>
				<span style="color: gray;">約1條評論</span>			</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:U3SJWI4HYXRT3V7NQKLN5MHH76G5LRXZ&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=http%3A%2F%2Fanidex.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.anirena.com%3A80%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:a6e49b2387c5e33dd7ed8296deb0e7ff8dd5c6f9&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">171MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">7</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">0</span></td>
			<td nowrap="nowrap" align="center">32</td>
			<td align="center"><a href="/topics/list/user_id/730809">Laputa</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 06:57			<span style="display: none;">2022/02/21 06:57</span></td>
			<td width="6%" align="center">
			<a class="sort-43" 
				href="/topics/list/sort_id/43">
				<b><font color=purple>動漫音樂</font></b></a>
			</td>
			<td class="title">
								
								
				<a href="/topics/view/593372_2022_01_19_Fate_Grand_Order_-_-_MP3_320K.html"  target="_blank" >
				[2022.01.19] 劇場版「Fate/Grand Order -終局特異点 冠位時間神殿ソロモン-」オリジナルサウンドトラック [MP3 320K]</a>
							</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:2EWBSHMT5TFS7F5G2IJKBD2RBYT2QA5U&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=http%3A%2F%2F208.67.16.113%3A8000%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:d12c191d93eccb2f97a6d212a08f510e27a803b4&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">165.4MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">19</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">1</span></td>
			<td nowrap="nowrap" align="center">152</td>
			<td align="center"><a href="/topics/list/user_id/677029">Mashin</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 06:57			<span style="display: none;">2022/02/21 06:57</span></td>
			<td width="6%" align="center">
			<a class="sort-43" 
				href="/topics/list/sort_id/43">
				<b><font color=purple>動漫音樂</font></b></a>
			</td>
			<td class="title">
								
								
				<a href="/topics/view/593371_2022_01_19_Fate_Grand_Order_-_-_FLAC.html"  target="_blank" >
				[2022.01.19] 劇場版「Fate/Grand Order -終局特異点 冠位時間神殿ソロモン-」オリジナルサウンドトラック [FLAC]</a>
							</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:UJ5CVWPUFK3JF2IBYD55QTYI2V2RL4UN&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=http%3A%2F%2F208.67.16.113%3A8000%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:a27a2ad9f42ab692e901c0fbd84f08d57515f28d&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">350.4MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">12</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">0</span></td>
			<td nowrap="nowrap" align="center">67</td>
			<td align="center"><a href="/topics/list/user_id/677029">Mashin</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 05:24			<span style="display: none;">2022/02/21 05:24</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/49" >
				华盟字幕社</a></span>				
				<a href="/topics/view/593370_3_07_MP4_720p.html"  target="_blank" >
				[澄空学园&雪飘工作室&华盟字幕社]擅长捉弄的高木同学3 第07话 MP4 720p</a>
				<span style="color: gray;">約3條評論</span>			</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:5QB4JWDYCQMHXKO7TQVTYRPE7AUIXUT7&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2F104.238.198.186%3A8000%2Fannounce&tr=http%3A%2F%2F94.228.192.98%2Fannounce&tr=http%3A%2F%2Ftracker.btcake.com%2Fannounce&tr=http%3A%2F%2Ftracker.ktxp.com%3A6868%2Fannounce&tr=http%3A%2F%2Ftracker.ktxp.com%3A7070%2Fannounce&tr=http%3A%2F%2Fbt.sc-ol.com%3A2710%2Fannounce&tr=http%3A%2F%2Fbtfile.sdo.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker3.torrentino.com%2Fannounce&tr=http%3A%2F%2Ftracker2.torrentino.com%2Fannounce&tr=http%3A%2F%2Fpubt.net%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.tfile.me%2Fannounce&tr=http%3A%2F%2Fbigfoot1942.sektori.org%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce&tr=http%3A%2F%2F184.105.151.166%3A6969%2Fannounce&tr=http%3A%2F%2F213.183.51.12%3A6969%2Fannounce&tr=http%3A%2F%2F51.222.84.64%3A1337%2Fannounce&tr=http%3A%2F%2F51.81.200.170%3A6699%2Fannounce&tr=http%3A%2F%2F93.158.213.92%3A1337%2Fannounce&tr=http%3A%2F%2F%5B2001%3A1b10%3A1000%3A8101%3A0%3A242%3Aac11%3A2%5D%3A6969%2Fannounce&tr=http%3A%2F%2F%5B2001%3A470%3A1%3A189%3A0%3A1%3A2%3A3%5D%3A6969%2Fannounce&tr=http%3A%2F%2F%5B2a04%3Aac00%3A1%3A3dd8%3A%3A1%3A2710%5D%3A2710%2Fannounce&tr=http%3A%2F%2Fbt.3dmgame.com%3A2710%2Fannounce&tr=http%3A%2F%2Fbt.ali213.net%3A8000%2Fannounce&tr=http%3A%2F%2Fbt.okmp3.ru%3A2710%2Fannounce&tr=http%3A%2F%2Fbt.unionpeer.org%3A777%2Fannounce&tr=http%3A%2F%2Fbt2.54new.com%3A8080%2Fannounce&tr=http%3A%2F%2Fbttracker.debian.org%3A6969%2Fannounce&tr=http%3A%2F%2Fexplodie.org%3A6969%2Fannounce&tr=http%3A%2F%2Fgooger.cc%3A1337%2Fannounce&tr=http%3A%2F%2Fmixfiend.com%3A6969%2Fannounce&tr=http%3A%2F%2Fmvgroup.org%3A2710%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=http%3A%2F%2Fopentracker.i2p.rocks%3A6969%2Fannounce&tr=http%3A%2F%2Fretracker.hotplug.ru%3A2710%2Fannounce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fsukebei.tracker.wf%3A8888%2Fannounce&tr=http%3A%2F%2Ftorrent.arjlover.net%3A2710%2Fannounce&tr=http%3A%2F%2Ftorrent.resonatingmedia.com%3A6969%2Fannounce&tr=http%3A%2F%2Ftorrent.rus.ec%3A2710%2Fannounce&tr=http%3A%2F%2Ftorrentsmd.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftr.cili001.com%3A8070%2Fannounce&tr=http%3A%2F%2Ftracker.ali213.net%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker.birkenwald.de%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.bittor.pw%3A1337%2Fannounce&tr=http%3A%2F%2Ftracker.breizh.pm%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.bt4g.com%3A2095%2Fannounce&tr=http%3A%2F%2Ftracker.dler.org%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.files.fm%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.gigatorrents.ws%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.grepler.com%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.loadbt.com%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.minglong.org%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&tr=http%3A%2F%2Ftracker.pussytorrents.org%3A3000%2Fannounce&tr=http%3A%2F%2Ftracker.tvunderground.org.ru%3A3218%2Fannounce&tr=http%3A%2F%2Ftracker.xdvdz.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.yoshi210.com%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.zerobytes.xyz%3A1337%2Fannounce&tr=http%3A%2F%2Ftracker3.dler.org%3A2710%2Fannounce&tr=http%3A%2F%2Fwww.mvgroup.org%3A2710%2Fannounce&tr=http%3A%2F%2Fwww.thetradersden.org%2Fforums%2Ftracker%3A80%2Fannounce.php&tr=http%3A%2F%2Fwww.tvnihon.com%3A6969%2Fannounce&tr=http%3A%2F%2F1337.abcvg.info%2Fannounce&tr=http%3A%2F%2F157.90.169.123%2Fannounce&tr=http%3A%2F%2F185.148.3.231%2Fannounce&tr=http%3A%2F%2F60-fps.org%2Fbt%3A80%2Fannounce.php&tr=http%3A%2F%2F95.107.48.115%2Fannounce&tr=http%3A%2F%2Fall4nothin.net%2Fannounce.php&tr=http%3A%2F%2Fbaibako.tv%2Fannounce&tr=http%3A%2F%2Fbig-boss-tracker.net%2Fannounce.php&tr=http%3A%2F%2Fbithq.org%2Fannounce.php&tr=http%3A%2F%2Fbluebird-hd.org%2Fannounce.php&tr=http%3A%2F%2Fbt-club.ws%2Fannounce&tr=http%3A%2F%2Fbt.rghost.net%2Fannounce&tr=http%3A%2F%2Fbtx.anifilm.tv%2Fannounce.php&tr=http%3A%2F%2Fcarbon-bonsai-621.appspot.com%2Fannounce&tr=http%3A%2F%2Fdata-bg.net%2Fannounce.php&tr=http%3A%2F%2Ffxtt.ru%2Fannounce&tr=http%3A%2F%2Fh4.trakx.nibba.trade%2Fannounce&tr=http%3A%2F%2Firrenhaus.dyndns.dk%2Fannounce.php&tr=http%3A%2F%2Fkinorun.com%2Fannounce.php&tr=http%3A%2F%2Fmasters-tb.com%2Fannounce.php&tr=http%3A%2F%2Fmediaclub.tv%2Fannounce.php&tr=http%3A%2F%2Fmilanesitracker.tekcities.com%2Fannounce&tr=http%3A%2F%2Fns349743.ip-91-121-106.eu%2Fannounce&tr=http%3A%2F%2Fopen.acgnxtracker.com%2Fannounce&tr=http%3A%2F%2Fopentracker.xyz%2Fannounce&tr=http%3A%2F%2Fretracker.telecom.by%2Fannounce&tr=http%3A%2F%2Fsecure.pow7.com%2Fannounce&tr=http%3A%2F%2Fsiambit.com%2Fannounce.php&tr=http%3A%2F%2Ft1.pow7.com%2Fannounce&tr=http%3A%2F%2Ft2.pow7.com%2Fannounce&tr=http%3A%2F%2Ftorrent-team.net%2Fannounce.php&tr=http%3A%2F%2Ftorrent.mp3quran.net%2Fannounce.php&tr=http%3A%2F%2Ftorrentzilla.org%2Fannounce&tr=http%3A%2F%2Ftorrentzilla.org%2Fannounce.php&tr=http%3A%2F%2Ftr.kxmp.cf%2Fannounce&tr=http%3A%2F%2Ftracker.anirena.com%2Fannounce&tr=http%3A%2F%2Ftracker.gbitt.info%2Fannounce&tr=http%3A%2F%2Ftracker.ipv6tracker.org%2Fannounce&tr=http%3A%2F%2Ftracker.ipv6tracker.ru%2Fannounce&tr=http%3A%2F%2Ftracker.lelux.fi%2Fannounce&tr=http%3A%2F%2Ftracker.noobsubs.net%2Fannounce&tr=http%3A%2F%2Ftracker.pow7.com%2Fannounce&tr=http%3A%2F%2Ftracker.trackerfix.com%2Fannounce&tr=http%3A%2F%2Ftracker1.bt.moack.co.kr%2Fannounce&tr=http%3A%2F%2Ftracker2.dler.org%2Fannounce&tr=http%3A%2F%2Fvps02.net.orel.ru%2Fannounce&tr=http%3A%2F%2Fwww.all4nothin.net%2Fannounce.php&tr=http%3A%2F%2Fwww.legittorrents.info%2Fannounce.php&tr=http%3A%2F%2Fwww.shnflac.net%2Fannounce.php&tr=http%3A%2F%2Fwww.tribalmixes.com%2Fannounce.php&tr=http%3A%2F%2Fwww.xwt-classics.net%2Fannounce.php&tr=http%3A%2F%2Fwww.zone-torrent.net%2Fannounce.php&tr=https%3A%2F%2F1337.abcvg.info%2Fannounce&tr=https%3A%2F%2Fcarbon-bonsai-621.appspot.com%2Fannounce&tr=https%3A%2F%2Fmytracker.fly.dev%2Fannounce&tr=https%3A%2F%2Fopen.kickasstracker.com%2Fannounce&tr=https%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=https%3A%2F%2Fopentracker.xyz%2Fannounce&tr=https%3A%2F%2Ftorrent.ubuntu.com%2Fannounce&tr=https%3A%2F%2Ftorrents.linuxmint.com%2Fannounce.php&tr=https%3A%2F%2Ftr.torland.ga%2Fannounce&tr=https%3A%2F%2Ftracker.bt-hash.com%2Fannounce&tr=https%3A%2F%2Ftracker.coalition.space%2Fannounce&tr=https%3A%2F%2Ftracker.iriseden.eu%2Fannounce&tr=https%3A%2F%2Ftracker.iriseden.fr%2Fannounce&tr=https%3A%2F%2Ftracker.kuroy.me%2Fannounce&tr=https%3A%2F%2Ftracker.lelux.fi%2Fannounce&tr=https%3A%2F%2Ftracker.lilithraws.cf%2Fannounce&tr=https%3A%2F%2Ftracker.nanoha.org%2Fannounce&tr=https%3A%2F%2Ftracker.nitrix.me%2Fannounce&tr=https%3A%2F%2Ftracker.shittyurl.org%2Fannounce&tr=https%3A%2F%2Ftracker.tamersunion.org%2Fannounce&tr=https%3A%2F%2Ftrakx.herokuapp.com%2Fannounce&tr=http%3A%2F%2F78.30.254.12%3A2710%2Fannounce&tr=http%3A%2F%2Fbuny.uk%3A6969%2Fannounce&tr=http%3A%2F%2Fipv4announce.sktorrent.eu%3A6969%2Fannounce&tr=http%3A%2F%2Fopenbittorrent.com%2Fannounce&tr=http%3A%2F%2Fp4p.arenabg.com%3A1337%2Fannounce&tr=http%3A%2F%2Frt.optizone.ru%2Fannounce&tr=http%3A%2F%2Ft.overflow.biz%3A6969%2Fannounce&tr=http%3A%2F%2Ftorrenttracker.nwc.acsalaska.net%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.moeking.me%3A6969%2Fannounce&tr=http%3A%2F%2Ftrackme.theom.nz%2Fannounce&tr=https%3A%2F%2Fqbittorrent-tracker.120181311.xyz%2Fannounce&tr=https%3A%2F%2Ft1.tokhmi.xyz%2Fannounce&tr=https%3A%2F%2Ftr.ready4.icu%2Fannounce&tr=https%3A%2F%2Ftracker.foreverpirates.co%2Fannounce&tr=https%3A%2F%2Ftracker.imgoingto.icu%2Fannounce&tr=https%3A%2F%2Ftracker.madassnews.net%2Fannounce&tr=https%3A%2F%2Ftracker.parrotsec.org%2Fannounce&tr=https%3A%2F%2Ftrackme.theom.nz%2Fannounce&tr=https%3A%2F%2Fw.wwwww.wtf%2Fannounce&tr=wss%3A%2F%2Ftracker.openwebtorrent.com%3A443%2Fannounce&tr=http%3A%2F%2Fretracker.sevstar.net%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.uw0.xyz%3A6969%2Fannounce&tr=http%3A%2F%2F207.241.226.111%3A6969%2Fannounce&tr=http%3A%2F%2F207.241.231.226%3A6969%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:ec03c4d87814187ba9df9c2b3c45e4f8288bd27f&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">154MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">20</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">12</span></td>
			<td nowrap="nowrap" align="center">97</td>
			<td align="center"><a href="/topics/list/user_id/32769">CAMOE</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 05:24			<span style="display: none;">2022/02/21 05:24</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/49" >
				华盟字幕社</a></span>				
				<a href="/topics/view/593369_3_06_MP4_720p.html"  target="_blank" >
				[澄空学园&雪飘工作室&华盟字幕社]擅长捉弄的高木同学3 第06话 MP4 720p</a>
				<span style="color: gray;">約1條評論</span>			</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:ZA7FQVHFUF64RAL3AJTUGS5RKEYBYFC5&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2F104.238.198.186%3A8000%2Fannounce&tr=http%3A%2F%2F94.228.192.98%2Fannounce&tr=http%3A%2F%2Ftracker.btcake.com%2Fannounce&tr=http%3A%2F%2Ftracker.ktxp.com%3A6868%2Fannounce&tr=http%3A%2F%2Ftracker.ktxp.com%3A7070%2Fannounce&tr=http%3A%2F%2Fbt.sc-ol.com%3A2710%2Fannounce&tr=http%3A%2F%2Fbtfile.sdo.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker3.torrentino.com%2Fannounce&tr=http%3A%2F%2Ftracker2.torrentino.com%2Fannounce&tr=http%3A%2F%2Fpubt.net%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.tfile.me%2Fannounce&tr=http%3A%2F%2Fbigfoot1942.sektori.org%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce&tr=http%3A%2F%2F184.105.151.166%3A6969%2Fannounce&tr=http%3A%2F%2F213.183.51.12%3A6969%2Fannounce&tr=http%3A%2F%2F51.222.84.64%3A1337%2Fannounce&tr=http%3A%2F%2F51.81.200.170%3A6699%2Fannounce&tr=http%3A%2F%2F93.158.213.92%3A1337%2Fannounce&tr=http%3A%2F%2F%5B2001%3A1b10%3A1000%3A8101%3A0%3A242%3Aac11%3A2%5D%3A6969%2Fannounce&tr=http%3A%2F%2F%5B2001%3A470%3A1%3A189%3A0%3A1%3A2%3A3%5D%3A6969%2Fannounce&tr=http%3A%2F%2F%5B2a04%3Aac00%3A1%3A3dd8%3A%3A1%3A2710%5D%3A2710%2Fannounce&tr=http%3A%2F%2Fbt.3dmgame.com%3A2710%2Fannounce&tr=http%3A%2F%2Fbt.ali213.net%3A8000%2Fannounce&tr=http%3A%2F%2Fbt.okmp3.ru%3A2710%2Fannounce&tr=http%3A%2F%2Fbt.unionpeer.org%3A777%2Fannounce&tr=http%3A%2F%2Fbt2.54new.com%3A8080%2Fannounce&tr=http%3A%2F%2Fbttracker.debian.org%3A6969%2Fannounce&tr=http%3A%2F%2Fexplodie.org%3A6969%2Fannounce&tr=http%3A%2F%2Fgooger.cc%3A1337%2Fannounce&tr=http%3A%2F%2Fmixfiend.com%3A6969%2Fannounce&tr=http%3A%2F%2Fmvgroup.org%3A2710%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=http%3A%2F%2Fopentracker.i2p.rocks%3A6969%2Fannounce&tr=http%3A%2F%2Fretracker.hotplug.ru%3A2710%2Fannounce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fsukebei.tracker.wf%3A8888%2Fannounce&tr=http%3A%2F%2Ftorrent.arjlover.net%3A2710%2Fannounce&tr=http%3A%2F%2Ftorrent.resonatingmedia.com%3A6969%2Fannounce&tr=http%3A%2F%2Ftorrent.rus.ec%3A2710%2Fannounce&tr=http%3A%2F%2Ftorrentsmd.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftr.cili001.com%3A8070%2Fannounce&tr=http%3A%2F%2Ftracker.ali213.net%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker.birkenwald.de%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.bittor.pw%3A1337%2Fannounce&tr=http%3A%2F%2Ftracker.breizh.pm%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.bt4g.com%3A2095%2Fannounce&tr=http%3A%2F%2Ftracker.dler.org%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.files.fm%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.gigatorrents.ws%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.grepler.com%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.loadbt.com%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.minglong.org%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&tr=http%3A%2F%2Ftracker.pussytorrents.org%3A3000%2Fannounce&tr=http%3A%2F%2Ftracker.tvunderground.org.ru%3A3218%2Fannounce&tr=http%3A%2F%2Ftracker.xdvdz.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.yoshi210.com%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.zerobytes.xyz%3A1337%2Fannounce&tr=http%3A%2F%2Ftracker3.dler.org%3A2710%2Fannounce&tr=http%3A%2F%2Fwww.mvgroup.org%3A2710%2Fannounce&tr=http%3A%2F%2Fwww.thetradersden.org%2Fforums%2Ftracker%3A80%2Fannounce.php&tr=http%3A%2F%2Fwww.tvnihon.com%3A6969%2Fannounce&tr=http%3A%2F%2F1337.abcvg.info%2Fannounce&tr=http%3A%2F%2F157.90.169.123%2Fannounce&tr=http%3A%2F%2F185.148.3.231%2Fannounce&tr=http%3A%2F%2F60-fps.org%2Fbt%3A80%2Fannounce.php&tr=http%3A%2F%2F95.107.48.115%2Fannounce&tr=http%3A%2F%2Fall4nothin.net%2Fannounce.php&tr=http%3A%2F%2Fbaibako.tv%2Fannounce&tr=http%3A%2F%2Fbig-boss-tracker.net%2Fannounce.php&tr=http%3A%2F%2Fbithq.org%2Fannounce.php&tr=http%3A%2F%2Fbluebird-hd.org%2Fannounce.php&tr=http%3A%2F%2Fbt-club.ws%2Fannounce&tr=http%3A%2F%2Fbt.rghost.net%2Fannounce&tr=http%3A%2F%2Fbtx.anifilm.tv%2Fannounce.php&tr=http%3A%2F%2Fcarbon-bonsai-621.appspot.com%2Fannounce&tr=http%3A%2F%2Fdata-bg.net%2Fannounce.php&tr=http%3A%2F%2Ffxtt.ru%2Fannounce&tr=http%3A%2F%2Fh4.trakx.nibba.trade%2Fannounce&tr=http%3A%2F%2Firrenhaus.dyndns.dk%2Fannounce.php&tr=http%3A%2F%2Fkinorun.com%2Fannounce.php&tr=http%3A%2F%2Fmasters-tb.com%2Fannounce.php&tr=http%3A%2F%2Fmediaclub.tv%2Fannounce.php&tr=http%3A%2F%2Fmilanesitracker.tekcities.com%2Fannounce&tr=http%3A%2F%2Fns349743.ip-91-121-106.eu%2Fannounce&tr=http%3A%2F%2Fopen.acgnxtracker.com%2Fannounce&tr=http%3A%2F%2Fopentracker.xyz%2Fannounce&tr=http%3A%2F%2Fretracker.telecom.by%2Fannounce&tr=http%3A%2F%2Fsecure.pow7.com%2Fannounce&tr=http%3A%2F%2Fsiambit.com%2Fannounce.php&tr=http%3A%2F%2Ft1.pow7.com%2Fannounce&tr=http%3A%2F%2Ft2.pow7.com%2Fannounce&tr=http%3A%2F%2Ftorrent-team.net%2Fannounce.php&tr=http%3A%2F%2Ftorrent.mp3quran.net%2Fannounce.php&tr=http%3A%2F%2Ftorrentzilla.org%2Fannounce&tr=http%3A%2F%2Ftorrentzilla.org%2Fannounce.php&tr=http%3A%2F%2Ftr.kxmp.cf%2Fannounce&tr=http%3A%2F%2Ftracker.anirena.com%2Fannounce&tr=http%3A%2F%2Ftracker.gbitt.info%2Fannounce&tr=http%3A%2F%2Ftracker.ipv6tracker.org%2Fannounce&tr=http%3A%2F%2Ftracker.ipv6tracker.ru%2Fannounce&tr=http%3A%2F%2Ftracker.lelux.fi%2Fannounce&tr=http%3A%2F%2Ftracker.noobsubs.net%2Fannounce&tr=http%3A%2F%2Ftracker.pow7.com%2Fannounce&tr=http%3A%2F%2Ftracker.trackerfix.com%2Fannounce&tr=http%3A%2F%2Ftracker1.bt.moack.co.kr%2Fannounce&tr=http%3A%2F%2Ftracker2.dler.org%2Fannounce&tr=http%3A%2F%2Fvps02.net.orel.ru%2Fannounce&tr=http%3A%2F%2Fwww.all4nothin.net%2Fannounce.php&tr=http%3A%2F%2Fwww.legittorrents.info%2Fannounce.php&tr=http%3A%2F%2Fwww.shnflac.net%2Fannounce.php&tr=http%3A%2F%2Fwww.tribalmixes.com%2Fannounce.php&tr=http%3A%2F%2Fwww.xwt-classics.net%2Fannounce.php&tr=http%3A%2F%2Fwww.zone-torrent.net%2Fannounce.php&tr=https%3A%2F%2F1337.abcvg.info%2Fannounce&tr=https%3A%2F%2Fcarbon-bonsai-621.appspot.com%2Fannounce&tr=https%3A%2F%2Fmytracker.fly.dev%2Fannounce&tr=https%3A%2F%2Fopen.kickasstracker.com%2Fannounce&tr=https%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=https%3A%2F%2Fopentracker.xyz%2Fannounce&tr=https%3A%2F%2Ftorrent.ubuntu.com%2Fannounce&tr=https%3A%2F%2Ftorrents.linuxmint.com%2Fannounce.php&tr=https%3A%2F%2Ftr.torland.ga%2Fannounce&tr=https%3A%2F%2Ftracker.bt-hash.com%2Fannounce&tr=https%3A%2F%2Ftracker.coalition.space%2Fannounce&tr=https%3A%2F%2Ftracker.iriseden.eu%2Fannounce&tr=https%3A%2F%2Ftracker.iriseden.fr%2Fannounce&tr=https%3A%2F%2Ftracker.kuroy.me%2Fannounce&tr=https%3A%2F%2Ftracker.lelux.fi%2Fannounce&tr=https%3A%2F%2Ftracker.lilithraws.cf%2Fannounce&tr=https%3A%2F%2Ftracker.nanoha.org%2Fannounce&tr=https%3A%2F%2Ftracker.nitrix.me%2Fannounce&tr=https%3A%2F%2Ftracker.shittyurl.org%2Fannounce&tr=https%3A%2F%2Ftracker.tamersunion.org%2Fannounce&tr=https%3A%2F%2Ftrakx.herokuapp.com%2Fannounce&tr=http%3A%2F%2F78.30.254.12%3A2710%2Fannounce&tr=http%3A%2F%2Fbuny.uk%3A6969%2Fannounce&tr=http%3A%2F%2Fipv4announce.sktorrent.eu%3A6969%2Fannounce&tr=http%3A%2F%2Fopenbittorrent.com%2Fannounce&tr=http%3A%2F%2Fp4p.arenabg.com%3A1337%2Fannounce&tr=http%3A%2F%2Frt.optizone.ru%2Fannounce&tr=http%3A%2F%2Ft.overflow.biz%3A6969%2Fannounce&tr=http%3A%2F%2Ftorrenttracker.nwc.acsalaska.net%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.moeking.me%3A6969%2Fannounce&tr=http%3A%2F%2Ftrackme.theom.nz%2Fannounce&tr=https%3A%2F%2Fqbittorrent-tracker.120181311.xyz%2Fannounce&tr=https%3A%2F%2Ft1.tokhmi.xyz%2Fannounce&tr=https%3A%2F%2Ftr.ready4.icu%2Fannounce&tr=https%3A%2F%2Ftracker.foreverpirates.co%2Fannounce&tr=https%3A%2F%2Ftracker.imgoingto.icu%2Fannounce&tr=https%3A%2F%2Ftracker.madassnews.net%2Fannounce&tr=https%3A%2F%2Ftracker.parrotsec.org%2Fannounce&tr=https%3A%2F%2Ftrackme.theom.nz%2Fannounce&tr=https%3A%2F%2Fw.wwwww.wtf%2Fannounce&tr=wss%3A%2F%2Ftracker.openwebtorrent.com%3A443%2Fannounce&tr=http%3A%2F%2Fretracker.sevstar.net%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.uw0.xyz%3A6969%2Fannounce&tr=http%3A%2F%2F207.241.226.111%3A6969%2Fannounce&tr=http%3A%2F%2F207.241.231.226%3A6969%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:c83e5854e5a17dc8817b0267434bb151301c145d&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">142.4MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">17</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">1</span></td>
			<td nowrap="nowrap" align="center">117</td>
			<td align="center"><a href="/topics/list/user_id/32769">CAMOE</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 03:15			<span style="display: none;">2022/02/21 03:15</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/801" >
				NC-Raws</a></span>				
				<a href="/topics/view/593368_NC-Raws_Yami_Shibai_S10_-_07_B-Global_1920x1080_HEVC_AAC_MKV.html"  target="_blank" >
				[NC-Raws] 暗芝居 第十季 / Yami Shibai S10 - 07 (B-Global 1920x1080 HEVC AAC MKV)</a>
							</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:PBVZDPIPG7TCXGLY7ZQIJ4WIZ53DDPOQ&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Fopen.nyaatorrents.info%3A6544%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=https%3A%2F%2Ftr.bangumi.moe%3A9696%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:786b91bd0f37e62b9978fe6084f2c8cf7631bdd0&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">121.9MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">1</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">0</span></td>
			<td nowrap="nowrap" align="center">9</td>
			<td align="center"><a href="/topics/list/user_id/637871">九十九朔夜</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 03:15			<span style="display: none;">2022/02/21 03:15</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/801" >
				NC-Raws</a></span>				
				<a href="/topics/view/593367_NC-Raws_Yami_Shibai_S10_-_07_Baha_1920x1080_AVC_AAC_MP4.html"  target="_blank" >
				[NC-Raws] 闇芝居 第十季 / Yami Shibai S10 - 07 (Baha 1920x1080 AVC AAC MP4)</a>
				<span style="color: gray;">約1條評論</span>			</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:LXG2ZLSNXW4PVMN6YAVX32SIRV33W7TF&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Fopen.nyaatorrents.info%3A6544%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=https%3A%2F%2Ftr.bangumi.moe%3A9696%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:5dcdacae4dbdb8fab1bec02b7dea488d77bb7e65&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">88.1MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">12</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">2</span></td>
			<td nowrap="nowrap" align="center">60</td>
			<td align="center"><a href="/topics/list/user_id/637871">九十九朔夜</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 02:34			<span style="display: none;">2022/02/21 02:34</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/520" >
				豌豆字幕组</a></span>				
				<a href="/topics/view/593366_Shingeki_no_Kyojin_82_1080P_MP4.html"  target="_blank" >
				【豌豆字幕组】[进击的巨人 / Shingeki_no_Kyojin][82][简体][1080P][MP4]</a>
				<span style="color: gray;">約1條評論</span>			</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:O3P722GL2KV4JIQASLBK2PTQHGORJV4F&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Fmgtracker.org%3A2710%2Fannounce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=https%3A%2F%2Ftr.bangumi.moe%3A9696%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2F104.238.198.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.kamigami.org%3A2710%2Fannounce&tr=https%3A%2F%2Ftracker.imgoingto.icu%2Fannounce&tr=https%3A%2F%2Ftr.ready4.icu%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:76dffd68cbd2abc4a20092c2ad3e70399d14d785&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">368.1MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">89</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">5</span></td>
			<td nowrap="nowrap" align="center">652</td>
			<td align="center"><a href="/topics/list/user_id/420797">春音爱良aira</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 02:34			<span style="display: none;">2022/02/21 02:34</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/520" >
				豌豆字幕组</a></span>				
				<a href="/topics/view/593365_Shingeki_no_Kyojin_82_1080P_MP4.html"  target="_blank" >
				【豌豆字幕組】[進擊的巨人 / Shingeki_no_Kyojin][82][繁體][1080P][MP4]</a>
				<span style="color: gray;">約2條評論</span>			</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:BEAQSAOQ4VRWHFBYLFLRCX3ZT2CGYM5D&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Fmgtracker.org%3A2710%2Fannounce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=https%3A%2F%2Ftr.bangumi.moe%3A9696%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2F104.238.198.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.kamigami.org%3A2710%2Fannounce&tr=https%3A%2F%2Ftracker.imgoingto.icu%2Fannounce&tr=https%3A%2F%2Ftr.ready4.icu%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:09010901d0e5636394385957115f799e846c33a3&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">367.9MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">77</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">34</span></td>
			<td nowrap="nowrap" align="center">712</td>
			<td align="center"><a href="/topics/list/user_id/420797">春音爱良aira</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 02:26			<span style="display: none;">2022/02/21 02:26</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/666" >
				中肯字幕組</a></span>				
				<a href="/topics/view/593364_2001_01-03_BIG5_MP4_1255x940_DVD_V2.html"  target="_blank" >
				【中肯字幕組】【懷舊老番】【刃牙 2001】【01-03】【BIG5_MP4】【1255x940】【DVD】V2</a>
				<span style="color: gray;">約3條評論</span>			</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:5XOTUCTH664DVTTFNON37EPOLFJSIV5S&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:eddd3a0a67f7b83ace656b9bbf91ee59532457b2&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">1.3GB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">3</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">1</span></td>
			<td nowrap="nowrap" align="center">25</td>
			<td align="center"><a href="/topics/list/user_id/665734">yellow9395</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 02:02			<span style="display: none;">2022/02/21 02:02</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/801" >
				NC-Raws</a></span>				
				<a href="/topics/view/593363_NC-Raws_The_Final_Season_Kyojin_F_Part_2_-_23_Baha_1920x1080_AVC_AAC_MP4.html"  target="_blank" >
				[NC-Raws] 進擊的巨人 The Final Season / Kyojin F Part 2 - 23 (Baha 1920x1080 AVC AAC MP4)</a>
				<span style="color: gray;">約1條評論</span>			</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:WFGS324XAAG3FCWY4RJSHDVAMO3FWZ5F&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Fopen.nyaatorrents.info%3A6544%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=https%3A%2F%2Ftr.bangumi.moe%3A9696%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:b14d2deb97000db28ad8e453238ea063b65b67a5&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">561.1MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">37</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">16</span></td>
			<td nowrap="nowrap" align="center">239</td>
			<td align="center"><a href="/topics/list/user_id/637871">九十九朔夜</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 01:45			<span style="display: none;">2022/02/21 01:45</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/803" >
				Lilith-Raws</a></span>				
				<a href="/topics/view/593362_Lilith-Raws_Shingeki_no_Kyojin_-_The_Final_Season_-_23_Baha_WEB-DL_1080p_AVC_AAC_CHT_MP4.html"  target="_blank" >
				[Lilith-Raws] 進擊的巨人 / Shingeki no Kyojin - The Final Season - 23 [Baha][WEB-DL][1080p][AVC AAC][CHT][MP4]</a>
				<span style="color: gray;">約1條評論</span>			</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:YN4OIPDP5BKATXLAYTOQMY4HCDYNC7UV&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Fopen.nyaatorrents.info%3A6544%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=https%3A%2F%2Ftracker.lilithraws.org%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.com%3A6869%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:c378e43c6fe85409dd60c4dd06638710f0d17e95&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">561MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">88</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">5</span></td>
			<td nowrap="nowrap" align="center">1009</td>
			<td align="center"><a href="/topics/list/user_id/727168">lilithraws</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 01:37			<span style="display: none;">2022/02/21 01:37</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/801" >
				NC-Raws</a></span>				
				<a href="/topics/view/593361_NC-Raws_The_Final_Season_Part_2_Kyojin_F_Part_2_-_07_B-Global_1920x1080_HEVC_AAC_MKV.html"  target="_blank" >
				[NC-Raws] 进击的巨人 The Final Season Part 2 / Kyojin F Part 2 - 07 (B-Global 1920x1080 HEVC AAC MKV)</a>
							</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:NEVVA5RABEDD4O5IKHGKHGQTZW6DTO6N&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Fopen.nyaatorrents.info%3A6544%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=https%3A%2F%2Ftr.bangumi.moe%3A9696%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:692b50762009063e3ba851cca39a13cdbc39bbcd&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">654MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">25</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">25</span></td>
			<td nowrap="nowrap" align="center">44</td>
			<td align="center"><a href="/topics/list/user_id/637871">九十九朔夜</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 01:36			<span style="display: none;">2022/02/21 01:36</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/801" >
				NC-Raws</a></span>				
				<a href="/topics/view/593360_NC-Raws_The_Final_Season_Kyojin_F_Part_2_-_23_Baha_1920x1080_AVC_AAC_MP4.html"  target="_blank" >
				[NC-Raws] 進擊的巨人 The Final Season / Kyojin F Part 2 - 23 (Baha 1920x1080 AVC AAC MP4)</a>
				<span style="color: gray;">約2條評論</span>			</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:5GVMFIRQF2FMMM6SIHFYR6NFBRIFKEW4&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Fopen.nyaatorrents.info%3A6544%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=https%3A%2F%2Ftr.bangumi.moe%3A9696%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:e9aac2a2302e8ac633d241cb88f9a50c505512dc&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">428.4MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">35</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">8</span></td>
			<td nowrap="nowrap" align="center">383</td>
			<td align="center"><a href="/topics/list/user_id/637871">九十九朔夜</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 00:33			<span style="display: none;">2022/02/21 00:33</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/803" >
				Lilith-Raws</a></span>				
				<a href="/topics/view/593359_Lilith-Raws_Sasaki_to_Miyano_-_07_Baha_WEB-DL_1080p_AVC_AAC_CHT_MP4.html"  target="_blank" >
				[Lilith-Raws] 佐佐木與宮野 / Sasaki to Miyano - 07 [Baha][WEB-DL][1080p][AVC AAC][CHT][MP4]</a>
							</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:S37Y2SCO2CEYZVQMNFUTCZKN5Z65YIGW&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Fopen.nyaatorrents.info%3A6544%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=https%3A%2F%2Ftracker.lilithraws.org%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.com%3A6869%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:96ff8d484ed0898cd60c696931654dee7ddc20d6&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">463.4MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">17</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">2</span></td>
			<td nowrap="nowrap" align="center">185</td>
			<td align="center"><a href="/topics/list/user_id/727168">lilithraws</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 00:33			<span style="display: none;">2022/02/21 00:33</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/803" >
				Lilith-Raws</a></span>				
				<a href="/topics/view/593358_Lilith-Raws_Futsal_Boys%21%21%21%21%21_-_07_Baha_WEB-DL_1080p_AVC_AAC_CHT_MP4.html"  target="_blank" >
				[Lilith-Raws] Futsal Boys!!!!! - 07 [Baha][WEB-DL][1080p][AVC AAC][CHT][MP4]</a>
							</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:2T3PMHYP25BHZC3LBIXGMJPWJNMBERNR&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Fopen.nyaatorrents.info%3A6544%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=https%3A%2F%2Ftracker.lilithraws.org%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.com%3A6869%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:d4f6f61f0fd7427c8b6b0a2e6625f64b581245b1&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">380.7MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">22</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">7</span></td>
			<td nowrap="nowrap" align="center">189</td>
			<td align="center"><a href="/topics/list/user_id/727168">lilithraws</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 00:26			<span style="display: none;">2022/02/21 00:26</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/604" >
				c.c动漫</a></span>				
				<a href="/topics/view/593357_c_c_1_-_07_BIG5_1080P_MP4.html"  target="_blank" >
				[c.c動漫][1月新番][鏽色鎧甲-黎明][07][BIG5][1080P][MP4]</a>
				<span style="color: gray;">約2條評論</span>			</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:2T6NE6K4TOXONRIEYT2Y6ZSKF4FWHTMG&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Fopen.nyaatorrents.info%3A6544%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:d4fcd2795c9baee6c504c4f58f664a2f0b63cd86&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">491.7MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">19</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">6</span></td>
			<td nowrap="nowrap" align="center">175</td>
			<td align="center"><a href="/topics/list/user_id/55561">lleeopen</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 00:08			<span style="display: none;">2022/02/21 00:08</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/657" >
				LoliHouse</a></span>				
				<a href="/topics/view/593356_LoliHouse_Ryman_s_Club_-_04_WebRip_1080p_HEVC-10bit_AAC.html"  target="_blank" >
				[LoliHouse] 白领羽球部 / Ryman's Club - 04 [WebRip 1080p HEVC-10bit AAC][简繁内封字幕]</a>
				<span style="color: gray;">約1條評論</span>			</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:H56SM2JMRSTLGDQ3XTIRI5CYZFTKJZMK&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Fopen.nyaatorrents.info%3A6544%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=http%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce&tr=http%3A%2F%2F104.238.198.186%3A8000%2Fannounce&tr=http%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&tr=https%3A%2F%2Ftr.bangumi.moe%3A9696%2Fannounce&tr=http%3A%2F%2F208.67.16.113%3A8000%2Fannounce&tr=http%3A%2F%2Fbtfile.sdo.com%3A6961%2Fannounce&tr=http%3A%2F%2Fp4p.arenabg.com%3A1337%2Fannounce&tr=http%3A%2F%2Fexplodie.org%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.internetwarriors.net%3A1337%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:3f7d26692c8ca6b30e1bbcd1147458c966a4e58a&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">336.8MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">10</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">0</span></td>
			<td nowrap="nowrap" align="center">65</td>
			<td align="center"><a href="/topics/list/user_id/689653">LoliHouse</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 00:01			<span style="display: none;">2022/02/21 00:01</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/801" >
				NC-Raws</a></span>				
				<a href="/topics/view/593355_NC-Raws_Sasaki_to_Miyano_-_07_B-Global_1920x1080_HEVC_AAC_MKV.html"  target="_blank" >
				[NC-Raws] 佐佐木与宮野 / Sasaki to Miyano - 07 (B-Global 1920x1080 HEVC AAC MKV)</a>
							</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:NZSTACAVBW6OI4GWYVQOCPYEHSDLI7VY&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Fopen.nyaatorrents.info%3A6544%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=https%3A%2F%2Ftr.bangumi.moe%3A9696%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:6e653008150dbce470d6c560e13f043c86b47eb8&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">646MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">3</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">0</span></td>
			<td nowrap="nowrap" align="center">9</td>
			<td align="center"><a href="/topics/list/user_id/637871">九十九朔夜</a></td>
		</tr>
			<tr class="">
			<td width="98">
			今天 00:01			<span style="display: none;">2022/02/21 00:01</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/801" >
				NC-Raws</a></span>				
				<a href="/topics/view/593354_NC-Raws_Sasaki_to_Miyano_-_07_Baha_1920x1080_AVC_AAC_MP4.html"  target="_blank" >
				[NC-Raws] 佐佐木與宮野 / Sasaki to Miyano - 07 (Baha 1920x1080 AVC AAC MP4)</a>
							</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:KK33TBJAPWYX66NMHKHP5SF4J3RJD7CR&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Fopen.nyaatorrents.info%3A6544%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=https%3A%2F%2Ftr.bangumi.moe%3A9696%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:52b7b985207db17f79ac3a8efec8bc4ee291fc51&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">463.6MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">10</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">5</span></td>
			<td nowrap="nowrap" align="center">149</td>
			<td align="center"><a href="/topics/list/user_id/637871">九十九朔夜</a></td>
		</tr>
			<tr class="">
			<td width="98">
			昨天 23:54			<span style="display: none;">2022/02/20 23:54</span></td>
			<td width="6%" align="center">
			<a class="sort-6" 
				href="/topics/list/sort_id/6">
				<font color=blue>日劇</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/648" >
				魔星字幕团</a></span>				
				<a href="/topics/view/593353_MagicStar_EP06_WEBDL_1080p.html"  target="_blank" >
				[MagicStar] 后宫婚。 / ハレ婚。 EP06 [WEBDL] [1080p]【生】【附日字】</a>
							</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:DYFIOXORYERNHVI6MTT7UMJ57ID7DON5&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:1e0a875dd1c122d3d51e64e7fa313dfa07f1b9bd&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">486.5MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">4</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">2</span></td>
			<td nowrap="nowrap" align="center">34</td>
			<td align="center"><a href="/topics/list/user_id/105217">jackieliiy</a></td>
		</tr>
			<tr class="">
			<td width="98">
			昨天 23:54			<span style="display: none;">2022/02/20 23:54</span></td>
			<td width="6%" align="center">
			<a class="sort-6" 
				href="/topics/list/sort_id/6">
				<font color=blue>日劇</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/648" >
				魔星字幕团</a></span>				
				<a href="/topics/view/593352_MagicStar_EP06_WEBDL_1080p.html"  target="_blank" >
				[MagicStar] 封刃师 / 封刃師 EP06 [WEBDL] [1080p]【生】【附日字】</a>
							</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:A74XR64BXORLYPKJ34VKAOODPCT7KA4H&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:07f978fb81bba2bc3d49df2aa039c378a7f50387&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">485.2MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">6</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">0</span></td>
			<td nowrap="nowrap" align="center">33</td>
			<td align="center"><a href="/topics/list/user_id/105217">jackieliiy</a></td>
		</tr>
			<tr class="">
			<td width="98">
			昨天 23:49			<span style="display: none;">2022/02/20 23:49</span></td>
			<td width="6%" align="center">
			<a class="sort-6" 
				href="/topics/list/sort_id/6">
				<font color=blue>日劇</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/648" >
				魔星字幕团</a></span>				
				<a href="/topics/view/593351_MagicStar_EP07_WEBDL_1080p.html"  target="_blank" >
				[MagicStar] 悠闲的赤胴铃之助 / まったり！赤胴鈴之助 EP07 [WEBDL] [1080p]【生】【附日字】</a>
							</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:YEPIC5C6QH2C67SETIVJ2F6TLLBJ45U6&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:c11e81745e81f42f7e449a2a9d17d35ac29e769e&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">476.7MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">2</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">0</span></td>
			<td nowrap="nowrap" align="center">17</td>
			<td align="center"><a href="/topics/list/user_id/105217">jackieliiy</a></td>
		</tr>
			<tr class="">
			<td width="98">
			昨天 23:42			<span style="display: none;">2022/02/20 23:42</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/804" >
				霜庭云花Sub</a></span>				
				<a href="/topics/view/593350_Sub_Arifureta_Shokugyou_de_Sekai_Saikyou_S2_06_1080P_AVC.html"  target="_blank" >
				[霜庭云花Sub&森之屋动画组][平凡职业成就世界最强 第二季 / Arifureta Shokugyou de Sekai Saikyou S2][06][1080P][AVC][简日内嵌]</a>
							</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:S3KUWXXYTEJVVRK6OXMI3ONB36LAXA5W&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Fopen.nyaatorrents.info%3A6544%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=http%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&tr=wss%3A%2F%2Ftracker.openwebtorrent.com%3A443%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce&tr=https%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.com%3A6869%2Fannounce&tr=http%3A%2F%2Fp4p.arenabg.com%3A1337%2Fannounce&tr=http%3A%2F%2Fopenbittorrent.com%2Fannounce&tr=http%3A%2F%2Fopentracker.i2p.rocks%3A6969%2Fannounce&tr=https%3A%2F%2Ftrakx.herokuapp.com%2Fannounce&tr=https%3A%2F%2Ftracker.tamersunion.org%2Fannounce&tr=https%3A%2F%2Ftracker.lilithraws.cf%2Fannounce&tr=https%3A%2F%2Ftracker.iriseden.fr%2Fannounce&tr=https%3A%2F%2Ftracker.iriseden.eu%2Fannounce&tr=http%3A%2F%2Ftracker.gbitt.info%2Fannounce&tr=http%3A%2F%2Ftracker.dler.org%3A6969%2Fannounce&tr=http%3A%2F%2Fh4.trakx.nibba.trade%2Fannounce&tr=http%3A%2F%2Ffxtt.ru%2Fannounce&tr=http%3A%2F%2Fexplodie.org%3A6969%2Fannounce&tr=https%3A%2F%2Fw.wwwww.wtf%2Fannounce&tr=https%3A%2F%2Ftracker.nitrix.me%2Fannounce&tr=https%3A%2F%2Ftracker.lelux.fi%2Fannounce&tr=https%3A%2F%2Ftracker.foreverpirates.co%2Fannounce&tr=https%3A%2F%2Ftr.torland.ga%2Fannounce&tr=https%3A%2F%2Fmytracker.fly.dev%2Fannounce&tr=https%3A%2F%2F1337.abcvg.info%2Fannounce&tr=http%3A%2F%2Fvps02.net.orel.ru%2Fannounce&tr=http%3A%2F%2Ftracker2.dler.org%2Fannounce&tr=http%3A%2F%2Ftracker1.bt.moack.co.kr%2Fannounce&tr=http%3A%2F%2Ftracker.zerobytes.xyz%3A1337%2Fannounce&tr=http%3A%2F%2Ftracker.noobsubs.net%2Fannounce&tr=http%3A%2F%2Ftracker.loadbt.com%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.lelux.fi%2Fannounce&tr=http%3A%2F%2Ftracker.files.fm%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.bt4g.com%3A2095%2Fannounce&tr=http%3A%2F%2Ftracker.breizh.pm%3A6969%2Fannounce&tr=http%3A%2F%2Ft.overflow.biz%3A6969%2Fannounce&tr=http%3A%2F%2Fretracker.sevstar.net%3A2710%2Fannounce&tr=http%3A%2F%2Fopen.acgnxtracker.com%2Fannounce&tr=http%3A%2F%2Fgooger.cc%3A1337%2Fannounce&tr=http%3A%2F%2F104.238.198.186%3A8000%2Fannounce&tr=http%3A%2F%2F94.228.192.98%2Fannounce&tr=http%3A%2F%2Ftracker.btcake.com%2Fannounce&tr=http%3A%2F%2Ftracker.ktxp.com%3A6868%2Fannounce&tr=http%3A%2F%2Ftracker.ktxp.com%3A7070%2Fannounce&tr=http%3A%2F%2Fbt.sc-ol.com%3A2710%2Fannounce&tr=http%3A%2F%2Fbtfile.sdo.com%3A6961%2Fannounce&tr=http%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker3.torrentino.com%2Fannounce&tr=http%3A%2F%2Ftracker2.torrentino.com%2Fannounce&tr=http%3A%2F%2Fpubt.net%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.tfile.me%2Fannounce&tr=http%3A%2F%2Fbigfoot1942.sektori.org%3A6969%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannoun&tr=http%3A%2F%2Ftracker.openbittorrent.com%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%2Fannounce&tr=http%3A%2F%2Fwww.36dm.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.36dm.com%3A2710%2Fannounce&tr=http%3A%2F%2F%5B2001%3A1b10%3A1000%3A8101%3A0%3A242%3Aac11%3A2%5D%3A6969%2Fannounce&tr=http%3A%2F%2F%5B2001%3A470%3A1%3A189%3A0%3A1%3A2%3A3%5D%3A6969%2Fannounce&tr=http%3A%2F%2F%5B2a04%3Aac00%3A1%3A3dd8%3A%3A1%3A2710%5D%3A2710%2Fannounce&tr=http%3A%2F%2F1337.abcvg.info%2Fannounce&tr=http%3A%2F%2F185.83.215.123%3A6969%2Fannounce&tr=http%3A%2F%2F62.210.202.61%2Fannounce&tr=http%3A%2F%2F78.30.254.12%3A2710%2Fannounce&tr=http%3A%2F%2F87.110.238.140%3A6969%2Fannounce&tr=http%3A%2F%2F95.211.168.204%3A2710%2Fannounce&tr=http%3A%2F%2Fmail2.zelenaya.net%2Fannounce&tr=http%3A%2F%2Fretracker.hotplug.ru%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.birkenwald.de%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.bz%2Fannounce&tr=http%3A%2F%2Ftracker.electro-torrent.pl%2Fannounce&tr=http%3A%2F%2Ftracker.moeking.me%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.nyaa.uk%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.torrentyorg.pl%2Fannounce&tr=http%3A%2F%2Ftracker.uw0.xyz%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker01.loveapp.com%3A6789%2Fannounce&tr=https%3A%2F%2F2.tracker.eu.org%2Fannounce&tr=https%3A%2F%2F3.tracker.eu.org%2Fannounce&tr=https%3A%2F%2Ftracker.gbitt.info%2Fannounce&tr=https%3A%2F%2Ftracker.nanoha.org%2Fannounce&tr=https%3A%2F%2Fseeders-paradise.org%2Fannounce&tr=http%3A%2F%2Ftracker.tfile.co%2Fannounce&tr=http%3A%2F%2Fprestige.minimafia.nl%3A443%2Fannounce&tr=http%3A%2F%2Fprivate.minimafia.nl%3A443%2Fannounce&tr=https%3A%2F%2Fopentracker.xyz%2Fannounce&tr=https%3A%2F%2Ft.quic.ws%2Fannounce&tr=https%3A%2F%2Ftracker.vectahosting.eu%3A2053%2Fannounce&tr=https%3A%2F%2Ftracker.parrotsec.org%2Fannounce&tr=http%3A%2F%2Fgwp2-v19.rinet.ru%2Fannounce&tr=http%3A%2F%2Fbt1.xxxxbt.cc%3A6969%2Fannounce&tr=http%3A%2F%2F93.158.213.92%3A1337%2Fannounce&tr=http%3A%2F%2F176.113.71.19%3A6961%2Fannounce&tr=http%3A%2F%2F184.105.151.164%3A6969%2Fannounce&tr=http%3A%2F%2F86.62.124.78%2Fannounce&tr=http%3A%2F%2F95.107.48.115%2Fannounce&tr=http%3A%2F%2F193.148.69.107%2Fannounce&tr=http%3A%2F%2F54.39.98.124%2Fannounce&tr=http%3A%2F%2Ftorrentclub.tech%3A6969%2Fannounce&tr=https%3A%2F%2Fopentracker.co%2Fannounce&tr=http%3A%2F%2Ftracker.tvunderground.org.ru%3A3218%2Fannounce&tr=http%3A%2F%2Fnewtoncity.org%3A6969%2Fannounce&tr=https%3A%2F%2Ftracker.publictorrent.net%2Fannounce&tr=http%3A%2F%2Fsub4all.org%3A2710%2Fannounce&tr=http%3A%2F%2Fagusiq-torrents.pl%3A6969%2Fannounce&tr=http%3A%2F%2Fbt.nnm-club.info%3A2710%2Fannounce&tr=http%3A%2F%2F121.14.98.151%3A9090%2Fannounce&tr=http%3A%2F%2Fbt.dmhy.net%2Fannonuce&tr=http%3A%2F%2Fshare.dmhy.me%2Fannonuce&tr=https%3A%2F%2Ftr.bangumi.moe%3A9696%2Fannounce&tr=https%3A%2F%2F1.tracker.animmouse.tk%2Fannounce&tr=https%3A%2F%2F2.tracker.animmouse.tk%2Fannounce&tr=https%3A%2F%2Ftracker.fastdownload.xyz%2Fannounce&tr=http%3A%2F%2Ftorrent.nwps.ws%2Fannounce&tr=http%3A%2F%2Fopen.trackerlist.xyz%2Fannounce&tr=http%3A%2F%2Fpeersteers.org%2Fannounce&tr=http%3A%2F%2Ftracker.publictorrent.net%2Fannounce&tr=http%3A%2F%2Fretracker.local%2Fannounce.php&tr=http%3A%2F%2Ftracker.ipv6tracker.org%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:96d54b5ef899135ac55e75d88db9a1df960b83b6&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">358.8MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">8</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">0</span></td>
			<td nowrap="nowrap" align="center">51</td>
			<td align="center"><a href="/topics/list/user_id/751592">霜庭云花Sub</a></td>
		</tr>
			<tr class="">
			<td width="98">
			昨天 23:36			<span style="display: none;">2022/02/20 23:36</span></td>
			<td width="6%" align="center">
			<a class="sort-43" 
				href="/topics/list/sort_id/43">
				<b><font color=purple>動漫音樂</font></b></a>
			</td>
			<td class="title">
								
								
				<a href="/topics/view/593349_2022_02_23_TV_OP2_Vaundy_MP3_320K.html"  target="_blank" >
				[2022.02.23] TVアニメ「王様ランキング」OP2テーマ「裸の勇者」／Vaundy [MP3 320K]</a>
							</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:5O25M7X4EUOP5UC6W4JGJ3SMLL6V2PTE&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=http%3A%2F%2F208.67.16.113%3A8000%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:ebb5d67efc251cfed05eb71264ee4c5afd5d3e64&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">32.9MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">21</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">2</span></td>
			<td nowrap="nowrap" align="center">193</td>
			<td align="center"><a href="/topics/list/user_id/677029">Mashin</a></td>
		</tr>
			<tr class="">
			<td width="98">
			昨天 23:36			<span style="display: none;">2022/02/20 23:36</span></td>
			<td width="6%" align="center">
			<a class="sort-43" 
				href="/topics/list/sort_id/43">
				<b><font color=purple>動漫音樂</font></b></a>
			</td>
			<td class="title">
								
								
				<a href="/topics/view/593348_2022_02_23_TV_OP2_Vaundy_FLAC_48kHz_24bit.html"  target="_blank" >
				[2022.02.23] TVアニメ「王様ランキング」OP2テーマ「裸の勇者」／Vaundy [FLAC 48kHz/24bit]</a>
							</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:TBIE3CPPU3Y4OOZGWPFRB2Z66HSNKQHN&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=http%3A%2F%2F208.67.16.113%3A8000%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:98504d89efa6f1c73b26b3cb10eb3ef1e4d540ed&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">178.7MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">12</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">0</span></td>
			<td nowrap="nowrap" align="center">92</td>
			<td align="center"><a href="/topics/list/user_id/677029">Mashin</a></td>
		</tr>
			<tr class="">
			<td width="98">
			昨天 23:00			<span style="display: none;">2022/02/20 23:00</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/303" >
				动漫国字幕组</a></span>				
				<a href="/topics/view/593347_10_Part_6_19_1080P_MP4.html"  target="_blank" >
				【动漫国字幕组】★10月新番[鲁邦三世 Part 6][19][1080P][简体][MP4]</a>
							</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:UHABHAR454JKQXGSUIZM52OFJ4DR7OQT&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Fmgtracker.org%3A2710%2Fannounce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=https%3A%2F%2Ftr.bangumi.moe%3A9696%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2F104.238.198.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.kamigami.org%3A2710%2Fannounce&tr=https%3A%2F%2Ftracker.imgoingto.icu%2Fannounce&tr=https%3A%2F%2Ftr.ready4.icu%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:a1c013823cef12a85cd2a232cee9c54f071fba13&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">348.4MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">2</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">3</span></td>
			<td nowrap="nowrap" align="center">28</td>
			<td align="center"><a href="/topics/list/user_id/420797">春音爱良aira</a></td>
		</tr>
			<tr class="">
			<td width="98">
			昨天 23:00			<span style="display: none;">2022/02/20 23:00</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/303" >
				动漫国字幕组</a></span>				
				<a href="/topics/view/593346_10_Part_6_19_1080P_MP4.html"  target="_blank" >
				【動漫國字幕組】★10月新番[魯邦三世 Part 6][19][1080P][繁體][MP4]</a>
				<span style="color: gray;">約2條評論</span>			</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:YCFWPEMF46NGSJMQTONLEFEB77P6SQZD&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Fmgtracker.org%3A2710%2Fannounce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=https%3A%2F%2Ftr.bangumi.moe%3A9696%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2F104.238.198.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.kamigami.org%3A2710%2Fannounce&tr=https%3A%2F%2Ftracker.imgoingto.icu%2Fannounce&tr=https%3A%2F%2Ftr.ready4.icu%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:c08b679185e79a6925909b9ab21481ffdfe94323&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">348.4MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">4</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">5</span></td>
			<td nowrap="nowrap" align="center">97</td>
			<td align="center"><a href="/topics/list/user_id/420797">春音爱良aira</a></td>
		</tr>
			<tr class="">
			<td width="98">
			昨天 23:00			<span style="display: none;">2022/02/20 23:00</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/303" >
				动漫国字幕组</a></span>				
				<a href="/topics/view/593345_10_Part_6_19_720P_MP4.html"  target="_blank" >
				【动漫国字幕组】★10月新番[鲁邦三世 Part 6][19][720P][简体][MP4]</a>
							</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:632S6QTCXFSN6BZMCCA2PZ6437UY37XC&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Fmgtracker.org%3A2710%2Fannounce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=https%3A%2F%2Ftr.bangumi.moe%3A9696%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2F104.238.198.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.kamigami.org%3A2710%2Fannounce&tr=https%3A%2F%2Ftracker.imgoingto.icu%2Fannounce&tr=https%3A%2F%2Ftr.ready4.icu%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:f6f52f4262b964df072c1081a7e7dcdfe98dfee2&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">180.6MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">6</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">3</span></td>
			<td nowrap="nowrap" align="center">36</td>
			<td align="center"><a href="/topics/list/user_id/420797">春音爱良aira</a></td>
		</tr>
			<tr class="">
			<td width="98">
			昨天 23:00			<span style="display: none;">2022/02/20 23:00</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/303" >
				动漫国字幕组</a></span>				
				<a href="/topics/view/593344_10_Part_6_19_720P_MP4.html"  target="_blank" >
				【動漫國字幕組】★10月新番[魯邦三世 Part 6][19][720P][繁體][MP4]</a>
							</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:6PXDHI55JHWNRQJA5TSPYJ6BGN4W7S24&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Fmgtracker.org%3A2710%2Fannounce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=https%3A%2F%2Ftr.bangumi.moe%3A9696%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2F104.238.198.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.kamigami.org%3A2710%2Fannounce&tr=https%3A%2F%2Ftracker.imgoingto.icu%2Fannounce&tr=https%3A%2F%2Ftr.ready4.icu%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:f3ee33a3bd49ecd8c120ece4fc27c133796fcb5c&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">180.6MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">7</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">1</span></td>
			<td nowrap="nowrap" align="center">126</td>
			<td align="center"><a href="/topics/list/user_id/420797">春音爱良aira</a></td>
		</tr>
			<tr class="">
			<td width="98">
			昨天 22:59			<span style="display: none;">2022/02/20 22:59</span></td>
			<td width="6%" align="center">
			<a class="sort-2" 
				href="/topics/list/sort_id/2">
				<font color=red>動畫</font></a>
			</td>
			<td class="title">
								
				<span class="tag">
				<a  href="/topics/list/team_id/804" >
				霜庭云花Sub</a></span>				
				<a href="/topics/view/593343_Sub_Arifureta_Shokugyou_de_Sekai_Saikyou_S2_06_1080P_AVC.html"  target="_blank" >
				[霜庭云花Sub&森之屋动画组][平凡職業造就世界最強 第二季 / Arifureta Shokugyou de Sekai Saikyou S2][06][1080P][AVC][繁日内嵌]</a>
							</td>
			<td nowrap="nowrap" align="center">
				<a class="download-arrow arrow-magnet" title="磁力下載" href="magnet:?xt=urn:btih:J2DK2EVWLGBG2ESIU66VX6YN7YTYPMUA&dn=&tr=http%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=udp%3A%2F%2F104.143.10.186%3A8000%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=https%3A%2F%2Ft-115.rhcloud.com%2Fonly_for_ylbud&tr=http%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=http%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker1.itzmx.com%3A8080%2Fannounce&tr=udp%3A%2F%2Ftracker2.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=udp%3A%2F%2Ftracker4.itzmx.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftr.bangumi.moe%3A6969%2Fannounce&tr=http%3A%2F%2Ft.nyaatracker.com%2Fannounce&tr=http%3A%2F%2Fopen.nyaatorrents.info%3A6544%2Fannounce&tr=http%3A%2F%2Ft2.popgo.org%3A7456%2Fannonce&tr=http%3A%2F%2Fshare.camoe.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Ftracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=http%3A%2F%2Ftracker3.itzmx.com%3A6961%2Fannounce&tr=http%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&tr=wss%3A%2F%2Ftracker.openwebtorrent.com%3A443%2Fannounce&tr=http%3A%2F%2Ft.acg.rip%3A6699%2Fannounce&tr=https%3A%2F%2Fopentracker.acgnx.se%2Fannounce&tr=http%3A%2F%2Fopentracker.acgnx.com%3A6869%2Fannounce&tr=http%3A%2F%2Fp4p.arenabg.com%3A1337%2Fannounce&tr=http%3A%2F%2Fopenbittorrent.com%2Fannounce&tr=http%3A%2F%2Fopentracker.i2p.rocks%3A6969%2Fannounce&tr=https%3A%2F%2Ftrakx.herokuapp.com%2Fannounce&tr=https%3A%2F%2Ftracker.tamersunion.org%2Fannounce&tr=https%3A%2F%2Ftracker.lilithraws.cf%2Fannounce&tr=https%3A%2F%2Ftracker.iriseden.fr%2Fannounce&tr=https%3A%2F%2Ftracker.iriseden.eu%2Fannounce&tr=http%3A%2F%2Ftracker.gbitt.info%2Fannounce&tr=http%3A%2F%2Ftracker.dler.org%3A6969%2Fannounce&tr=http%3A%2F%2Fh4.trakx.nibba.trade%2Fannounce&tr=http%3A%2F%2Ffxtt.ru%2Fannounce&tr=http%3A%2F%2Fexplodie.org%3A6969%2Fannounce&tr=https%3A%2F%2Fw.wwwww.wtf%2Fannounce&tr=https%3A%2F%2Ftracker.nitrix.me%2Fannounce&tr=https%3A%2F%2Ftracker.lelux.fi%2Fannounce&tr=https%3A%2F%2Ftracker.foreverpirates.co%2Fannounce&tr=https%3A%2F%2Ftr.torland.ga%2Fannounce&tr=https%3A%2F%2Fmytracker.fly.dev%2Fannounce&tr=https%3A%2F%2F1337.abcvg.info%2Fannounce&tr=http%3A%2F%2Fvps02.net.orel.ru%2Fannounce&tr=http%3A%2F%2Ftracker2.dler.org%2Fannounce&tr=http%3A%2F%2Ftracker1.bt.moack.co.kr%2Fannounce&tr=http%3A%2F%2Ftracker.zerobytes.xyz%3A1337%2Fannounce&tr=http%3A%2F%2Ftracker.noobsubs.net%2Fannounce&tr=http%3A%2F%2Ftracker.loadbt.com%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.lelux.fi%2Fannounce&tr=http%3A%2F%2Ftracker.files.fm%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.bt4g.com%3A2095%2Fannounce&tr=http%3A%2F%2Ftracker.breizh.pm%3A6969%2Fannounce&tr=http%3A%2F%2Ft.overflow.biz%3A6969%2Fannounce&tr=http%3A%2F%2Fretracker.sevstar.net%3A2710%2Fannounce&tr=http%3A%2F%2Fopen.acgnxtracker.com%2Fannounce&tr=http%3A%2F%2Fgooger.cc%3A1337%2Fannounce&tr=http%3A%2F%2F104.238.198.186%3A8000%2Fannounce&tr=http%3A%2F%2F94.228.192.98%2Fannounce&tr=http%3A%2F%2Ftracker.btcake.com%2Fannounce&tr=http%3A%2F%2Ftracker.ktxp.com%3A6868%2Fannounce&tr=http%3A%2F%2Ftracker.ktxp.com%3A7070%2Fannounce&tr=http%3A%2F%2Fbt.sc-ol.com%3A2710%2Fannounce&tr=http%3A%2F%2Fbtfile.sdo.com%3A6961%2Fannounce&tr=http%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker3.torrentino.com%2Fannounce&tr=http%3A%2F%2Ftracker2.torrentino.com%2Fannounce&tr=http%3A%2F%2Fpubt.net%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.tfile.me%2Fannounce&tr=http%3A%2F%2Fbigfoot1942.sektori.org%3A6969%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannoun&tr=http%3A%2F%2Ftracker.openbittorrent.com%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%2Fannounce&tr=http%3A%2F%2Fwww.36dm.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.36dm.com%3A2710%2Fannounce&tr=http%3A%2F%2F%5B2001%3A1b10%3A1000%3A8101%3A0%3A242%3Aac11%3A2%5D%3A6969%2Fannounce&tr=http%3A%2F%2F%5B2001%3A470%3A1%3A189%3A0%3A1%3A2%3A3%5D%3A6969%2Fannounce&tr=http%3A%2F%2F%5B2a04%3Aac00%3A1%3A3dd8%3A%3A1%3A2710%5D%3A2710%2Fannounce&tr=http%3A%2F%2F1337.abcvg.info%2Fannounce&tr=http%3A%2F%2F185.83.215.123%3A6969%2Fannounce&tr=http%3A%2F%2F62.210.202.61%2Fannounce&tr=http%3A%2F%2F78.30.254.12%3A2710%2Fannounce&tr=http%3A%2F%2F87.110.238.140%3A6969%2Fannounce&tr=http%3A%2F%2F95.211.168.204%3A2710%2Fannounce&tr=http%3A%2F%2Fmail2.zelenaya.net%2Fannounce&tr=http%3A%2F%2Fretracker.hotplug.ru%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.birkenwald.de%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.bz%2Fannounce&tr=http%3A%2F%2Ftracker.electro-torrent.pl%2Fannounce&tr=http%3A%2F%2Ftracker.moeking.me%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.nyaa.uk%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.torrentyorg.pl%2Fannounce&tr=http%3A%2F%2Ftracker.uw0.xyz%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker01.loveapp.com%3A6789%2Fannounce&tr=https%3A%2F%2F2.tracker.eu.org%2Fannounce&tr=https%3A%2F%2F3.tracker.eu.org%2Fannounce&tr=https%3A%2F%2Ftracker.gbitt.info%2Fannounce&tr=https%3A%2F%2Ftracker.nanoha.org%2Fannounce&tr=https%3A%2F%2Fseeders-paradise.org%2Fannounce&tr=http%3A%2F%2Ftracker.tfile.co%2Fannounce&tr=http%3A%2F%2Fprestige.minimafia.nl%3A443%2Fannounce&tr=http%3A%2F%2Fprivate.minimafia.nl%3A443%2Fannounce&tr=https%3A%2F%2Fopentracker.xyz%2Fannounce&tr=https%3A%2F%2Ft.quic.ws%2Fannounce&tr=https%3A%2F%2Ftracker.vectahosting.eu%3A2053%2Fannounce&tr=https%3A%2F%2Ftracker.parrotsec.org%2Fannounce&tr=http%3A%2F%2Fgwp2-v19.rinet.ru%2Fannounce&tr=http%3A%2F%2Fbt1.xxxxbt.cc%3A6969%2Fannounce&tr=http%3A%2F%2F93.158.213.92%3A1337%2Fannounce&tr=http%3A%2F%2F176.113.71.19%3A6961%2Fannounce&tr=http%3A%2F%2F184.105.151.164%3A6969%2Fannounce&tr=http%3A%2F%2F86.62.124.78%2Fannounce&tr=http%3A%2F%2F95.107.48.115%2Fannounce&tr=http%3A%2F%2F193.148.69.107%2Fannounce&tr=http%3A%2F%2F54.39.98.124%2Fannounce&tr=http%3A%2F%2Ftorrentclub.tech%3A6969%2Fannounce&tr=https%3A%2F%2Fopentracker.co%2Fannounce&tr=http%3A%2F%2Ftracker.tvunderground.org.ru%3A3218%2Fannounce&tr=http%3A%2F%2Fnewtoncity.org%3A6969%2Fannounce&tr=https%3A%2F%2Ftracker.publictorrent.net%2Fannounce&tr=http%3A%2F%2Fsub4all.org%3A2710%2Fannounce&tr=http%3A%2F%2Fagusiq-torrents.pl%3A6969%2Fannounce&tr=http%3A%2F%2Fbt.nnm-club.info%3A2710%2Fannounce&tr=http%3A%2F%2F121.14.98.151%3A9090%2Fannounce&tr=http%3A%2F%2Fbt.dmhy.net%2Fannonuce&tr=http%3A%2F%2Fshare.dmhy.me%2Fannonuce&tr=https%3A%2F%2Ftr.bangumi.moe%3A9696%2Fannounce&tr=https%3A%2F%2F1.tracker.animmouse.tk%2Fannounce&tr=https%3A%2F%2F2.tracker.animmouse.tk%2Fannounce&tr=https%3A%2F%2Ftracker.fastdownload.xyz%2Fannounce&tr=http%3A%2F%2Ftorrent.nwps.ws%2Fannounce&tr=http%3A%2F%2Fopen.trackerlist.xyz%2Fannounce&tr=http%3A%2F%2Fpeersteers.org%2Fannounce&tr=http%3A%2F%2Ftracker.publictorrent.net%2Fannounce&tr=http%3A%2F%2Fretracker.local%2Fannounce.php&tr=http%3A%2F%2Ftracker.ipv6tracker.org%2Fannounce">&nbsp;</a>
				<a class="download-pikpak" target="_blank" title="保存到PikPak可以用手機隨時觀賞" onclick="_hmt.push(['_trackEvent', 'pikpak', 'click', 'list'])" href="https://drive.mypikpak.com/landing?__add_url=magnet:?xt=urn:btih:4e86ad12b659826d1248a7bd5bfb0dfe2787b280&__source=dmhy&__campaign=dmhyh5">&nbsp;</a>
			</td>
			<td nowrap="nowrap" align="center">367.3MB</td>
						<td nowrap="nowrap" align="center"><span class="btl_1">6</span></td>
			<td nowrap="nowrap" align="center"><span class="bts_1">2</span></td>
			<td nowrap="nowrap" align="center">77</td>
			<td align="center"><a href="/topics/list/user_id/751592">霜庭云花Sub</a></td>
		</tr>
		</tbody>
</table>
<style>
.download-pikpak{
	background: url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAFAAAAAUCAYAAAFt34pBAAAAAXNSR0IArs4c6QAAAERlWElmTU0AKgAAAAgAAYdpAAQAAAABAAAAGgAAAAAAA6ABAAMAAAABAAEAAKACAAQAAAABAAAAUKADAAQAAAABAAAAFAAAAACck7LGAAAHh0lEQVRYCeVYCZSOZRR+vjH2CaNI1LGPTmUZ20zE2MmxFM4h1DkOU4lIlkQLOac9UodsUZYmCUfZUmIoWixjCROpiVG2ppikxN/z/Pf//P//+X9DnHTqnvN93/ve9373ve9973uX1wGh1gDfcX7i1I4IDjJjXKLN4wE9Avf7VA/rw4dqjpok9rmDiQOB+KJAzq8BosDnLGE4OryX8YrjaGqf0OJ4563AxL5G9EQ34N1HrS2aWGvae/kmYMNuaz/1NtC2bnD0nKk7JgEF8gOnzwDz1wEiWPokfxod/EmzpVQHfKRp8XgQr1ZsrYG+XVz9WRjV3ZpiKOEXrgdeXWw4LUNKEQjfvp61Q98m4UDfFDJNDR242DYZtdj8irPScZUoBtfFA8M6A8W5LY/PBrKPAtcUAz4cA0xaBkxebtNM7Q/UrQrUfshW4U5OpukxbkffpaOAJtRNYiVg8RM2cuQYkDzE2p8+b9/n5gMrNocz0wg1lBLGUMifaJS/n1LL4O6m9qOkazjMcF0aAhVKuxThXzHMcVGHfgZKXgUU5C7vO2zYWauAP0+7FFRUTWD1NqBrQNrgiL+V4N+U2oN8jc6cxhrP4JXu5vCklHRohlEtZgU397NMYBMPxYY9wC3luT83cnOrACPfBLZl2RquKgz0bglk7AU6JAOZ+4FlG4FWicC0FXYuStBwxszlOed+J5QDKpchLfn/7DnjXq3EnM+cS9ECW9YikxPA+zSipGpAncrADVcDVcoGWR3/DejWGBjaCXhtqeFlbBJCoPMmZ/NgO6A1ha5I+xBO1p4XxHCPW5yPqBDtp2cz4KEpwK8ngeuviUw9Oo2+6mnT2n1tzLM8OtNouzwLfM6dWPMVINeUmW3GnnUoMq9QrN8GhUgc4FvNRaWEDl7hdgJtcHeogM19DiZwy7mRVwxy6czTMsY797oS+AWkN/yJCPrBfw/E5EPjTeOctbGBrY0oXCIPRKObuf0VgYP0ab/Q0Fdu8WHjN44/ernLGZ8KNL4FWLcT6DcJ+HIsUO9hd5SBZ4TZ5xIetHj6xcIFgI8yeKhmBGm8rYDbc2Ki2d2kBzhZWzrVPxmkGUdmrjIhe7d2MKZHOLt5n1p/3CKG3EAQX/+C4ZJpMDpYg+8A5I52HwDmfQJ8vDWcR6SeXGBspAHh5FIEdaoAqa2tHfoeMSvY+2SHtauXB558C5jBkN2ePq5ZDRNk/xFg+kfcCe6IosvJU8BRhqw8wYe7YvIikkM+RS2e+B3owa2LBLMHG7bydcD2LGvPG27CyY9WuNYCsuJjvap00PWBh6nRC4C4qBp0fx7aBchPKj0v93Gx4d/Dv1h/Trp9lQCVLWntDzOA7w6asEPuBOTU11LjBfOcmf8rHzU20d/FigCDpnLFr9Pz04YiwaBphi2cn8bPrVVe9Oy7hlMkiAv8p11QaNvxfSQu5+IcH/op4aLS8bV3uCntZ2xvL9b6fSdajA4d7dQA2LWPk/PxgmxTcbtNbUAJr6B+AvDFObPaWOBtyYI60YQMI/8HO1R6OtPVJpqS7SBIUCKmRnM9Qcr/Z0uKo25SFYJdDfgVGNjhz4mkdV95iCtk2dyxE8DeH00elQ1KJWfQnUaDDknmZpXCeqFoQaB8aXvkgvMwf+/v3r6KhiQpUpG4ObV6HrG8/wI1KwLP9wJKF7e0k//7a8ad2cDwN4DODYF7mtC86cK0Q3qyDgOPzQREo5LRC/nzARP6WvVSowJT28l0LKTVouUxdeb7tAK2fktfUMecmJdPr+asoPeaf1E4SGF2sJM+pcttDDPMZJWxtqIfmb3KO/tF92VoX1N3LVSQqC4OhN0LY7T+xfC6z/1LispHpUmh3sU5xBYu4OAZem/V1F6Qo5zSn/9TkVLmyT94XrZbMH8rHfiNfSlQKVL/dkB31hcq/kMLGqVMUuDT77CWv59FP2VR8FeclkzaiE0cV0n50kJgzw9eKS6y7w9zF6k8TaEapTjDnx6lT6p39xwAFBKL8vgp18vlju/IMqFFV6yI44/l0URUlBnIcCrlCZZsMB77jjD+v2c4961jrdAaqjyNOQ4LMlrYgpGWgWkTk5nE6P+XFxE3yKpNpX+y0EsG6o57hNxLYaTdXfylWZVrdcLpqM1cde4io82ljGxyP2DOaqPQRUNSQvCyYeUWu4t5pDNL6x6AmzyJ+nYe6ekDAJXdHzDMq9ResA4oRRfTMpEVK3mrPBfdbTfZEU5bY/Nc4js3lg4qjfad+ncZ6cZlOLNF5ce5zPL8ZTgtocGNJvCF8tV9RNfnWOYz3RXIySsfF8jKezYBOiazquax1jE+ozMZAN1f6NERDgXdCMkiy5UMWmd8XCjFJbapu1gVn3SGCZQn5ULZ5eT6EB9HLYWA/MsJ+qld+7lgHtkyJUIGA83soz6m0uH/uVSHj9Gq+LgKlDtQ5VWkoAUFXe6MSnOpw7+VyvCqdbgpSf7SBf2bUNZ6umNRW+7lcgBXkb6Zuju7msDV2yIyj78cE/yHeeTwMqGjLhO0xrMKDF2wP7Vx0JW4xjze5fi9nIYfOtW/va0romwKuYZ121xdlnsF/guwv3D560U5egAAAABJRU5ErkJggg==) left center no-repeat;
    width: 80px;
    height: 20px;
    display: inline-block;
}
.download-pikpak:hover{
	border-bottom: none !important;
}
</style>
<script language="javascript">
var ts = $.tablesorter;
ts.addParser({
	id: "magicDate",
	is: function(s) {
		return false;
	},
	format: function(s) {
		s = $(s).html();
		return $.tablesorter.formatFloat((s != "") ? new Date(s.replace(new RegExp(/-/g),"/")).getTime() : "0");
	},
	type: "numeric"
});

	jQuery("#topic_list").tablesorter({ widgets: ['zebra'] });
    jQuery("#topic_list").bind("sortStart",function() { 
        jQuery("#overlay").show(); 
    }).bind("sortEnd",function() { 
        jQuery("#overlay").hide(); 
    }); 
</script>
<div class="nav_title">
			上一頁	第1頁 
		<a href="/index/index/page/2">下一頁</a>
	</div>
</div>
</div>
<div class="table">
<div class="nav_title">善意提醒</div>
<div class="content">
對於容量很小的內容請多加留意，多半是捆綁病毒和木馬的資源。<br />
對於視頻文件採用壓縮包形式發佈的，也請看清楚其真實的目的，建議堅決抵制該類下載（經驗告訴我們，這些大都是捆綁了木馬病毒或廣告資料或者是加了密碼的無良人士發佈的）。<br />
如果發現上述問題，請及時向我們<a href="http://bbs.dmhy.org/thread.php?fid-504.html">舉報</a>。
</div>
</div>

<div class="table">
<div class="nav_title">重要聲明</div>
<div class="content">
動漫花園資源網(下稱本站)只是一個網絡交流平台，所有資源均不是由本站上載或發放。本站以即時上載的方式運作，動漫花園資源網對所有留言的真實性、完整性及立場等，不會負上任何責任。一切內容只屬留言者個人意見，並不代表本網站之立場。用戶不應信賴內容，並應自行判斷內容之真實性。於有疑問的情況下，用戶應立即尋求專業意見。由於是受到「即時留言」運作方式所限制，故不能完全監察所有即時留言，若讀者發現有留言出現問題，請聯絡我們。本站有權拒絕任何人士留言及刪除、保留任何留言。切勿撰寫粗言穢語、誹謗、渲染色情暴力或人身攻擊的言論，敬請自律。本網站保留所有權利。
</div>
</div><div id="pkpk" align="center">
<a href="https://weidian.com/?userid=1822696050&distributorid=1607560422&wfr=retailer_wxh5&share_relation=ca44f380a058d285_1607560422_1"><img src="/1280pik.png?6" ></a></div><div class="footer">
<div class="top"><a href="#top" title="返回頁面頂部">TOP</a></div>
<!-- Histats.com  (div with counter) --><div id="histats_counter"></div>
<!-- Histats.com  START  (aync)-->
<script type="text/javascript">var _Hasync= _Hasync|| [];
_Hasync.push(['Histats.start', '1,3801674,4,1034,150,25,00010001']);
_Hasync.push(['Histats.fasi', '1']);
_Hasync.push(['Histats.track_hits', '']);
(function() {
var hs = document.createElement('script'); hs.type = 'text/javascript'; hs.async = true;
hs.src = ('//s10.histats.com/js15_as.js');
(document.getElementsByTagName('head')[0] || document.getElementsByTagName('body')[0]).appendChild(hs);
})();</script>
<noscript><a href="/" target="_blank"><img  src="//sstatic1.histats.com/0.gif?3801674&101" alt="" border="0"></a></noscript>
<!-- Histats.com  END  -->
<div style="clear:both;">
©2002-2021 <a href="/">動漫花園</a>。</div>
</div>
 

<div style="text-align: center; font: 11px arial; color: #AAA; clear: both;">
Processed in 0.003 second(s), 2 queries</div>
<div
	style="text-align: center; font: 11px arial; color: #AAA; clear: both;margin-bottom: 10px;">
Powered by <a style="color: white;" href="http://www.tedmind.com" target="_blank">tedmind.com</a>&nbsp;Version: 1.8.1</div>
</div>
</div>
</body>
</html>
`

const nyaaHtml = `<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="utf-8">
		<title>Browse :: Nyaa</title>

		<meta name="viewport" content="width=480px">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<link rel="shortcut icon" type="image/png" href="/static/favicon.png">
		<link rel="icon" type="image/png" href="/static/favicon.png">
		<link rel="mask-icon" href="/static/pinned-tab.svg" color="#3582F7">
		<link rel="alternate" type="application/rss+xml" href="https://nyaa.si/?page=rss" />

		<meta property="og:site_name" content="Nyaa">
		<meta property="og:title" content="Browse :: Nyaa">
		<meta property="og:image" content="/static/img/avatar/default.png">
<meta name="description" content="A BitTorrent community focused on Eastern Asian media including anime, manga, music, and more">
<meta name="keywords" content="torrents, bittorrent, torrent, anime, manga, sukebei, download, nyaa, magnet, magnets">
<meta property="og:description" content="Nyaa homepage">

		<!-- Bootstrap core CSS -->
		<!--
			Note: This has been customized at http://getbootstrap.com/customize/ to
			set the column breakpoint to tablet mode, instead of mobile. This is to
			make the navbar not look awful on tablets.
		-->
		<link href="/static/css/bootstrap.min.css?t=1608007388" rel="stylesheet" id="bsThemeLink">
		<link href="/static/css/bootstrap-xl-mod.css?t=1608007388" rel="stylesheet">
		<!--
			This theme changer script needs to be inline and right under the above stylesheet link to prevent FOUC (Flash Of Unstyled Content)
			Development version is commented out in static/js/main.js at the bottom of the file
		-->
		<script>function toggleDarkMode(){"dark"===localStorage.getItem("theme")?setThemeLight():setThemeDark()}function setThemeDark(){bsThemeLink.href="/static/css/bootstrap-dark.min.css?t=1608007388",localStorage.setItem("theme","dark"),document.body!==null&&document.body.classList.add('dark')}function setThemeLight(){bsThemeLink.href="/static/css/bootstrap.min.css?t=1608007388",localStorage.setItem("theme","light"),document.body!==null&&document.body.classList.remove('dark')}if("undefined"!=typeof Storage){var bsThemeLink=document.getElementById("bsThemeLink");"dark"===localStorage.getItem("theme")&&setThemeDark()}</script>
		<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bootstrap-select/1.12.2/css/bootstrap-select.min.css" integrity="sha256-an4uqLnVJ2flr7w0U74xiF4PJjO2N5Df91R2CUmCLCA=" crossorigin="anonymous" />
		<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css" integrity="sha256-eZrrJcwDc/3uDhsdt61sL2oOBY362qM3lon1gyExkL0=" crossorigin="anonymous" />

		<!-- Custom styles for this template -->
		<link href="/static/css/main.css?t=1645052395" rel="stylesheet">

		<!-- Core JavaScript -->
		<script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.2.1/jquery.min.js" integrity="sha256-hwg4gsxgFZhOsEEamdOYGBf13FyQuiTwlAQgxVSNgt4=" crossorigin="anonymous"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/3.3.7/js/bootstrap.min.js" integrity="sha256-U5ZEeKfGNOja007MMD3YBI0A3OSZOQbeG6z2f2Y0hu8=" crossorigin="anonymous"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/markdown-it/8.3.1/markdown-it.min.js" integrity="sha256-3WZyZQOe+ql3pLo90lrkRtALrlniGdnf//gRpW0UQks=" crossorigin="anonymous"></script>
		<!-- Modified to not apply border-radius to selectpickers and stuff so our navbar looks cool -->
		<script src="/static/js/bootstrap-select.min.js?t=1623304983"></script>
		<script src="/static/js/main.min.js?t=1623304983"></script>

		<!-- HTML5 shim and Respond.js for IE8 support of HTML5 elements and media queries -->
		<!--[if lt IE 9]>
			<script src="https://oss.maxcdn.com/html5shiv/3.7.3/html5shiv.min.js"></script>
			<script src="https://oss.maxcdn.com/respond/1.4.2/respond.min.js"></script>
		<![endif]-->

		<link rel="search" type="application/opensearchdescription+xml" title="Nyaa.si" href="/static/search.xml">
	</head>
	<body>
		<!-- Fixed navbar -->
		<nav class="navbar navbar-default navbar-static-top navbar-inverse">
			<div class="container">
				<div class="navbar-header">
					<button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar" aria-expanded="false" aria-controls="navbar">
					<span class="sr-only">Toggle navigation</span>
					<span class="icon-bar"></span>
					<span class="icon-bar"></span>
					<span class="icon-bar"></span>
					</button>
					<a class="navbar-brand" href="/">Nyaa</a>
				</div><!--/.navbar-header -->
				<div id="navbar" class="navbar-collapse collapse">
					<ul class="nav navbar-nav">
						<li ><a href="/upload">Upload</a></li>
						<li class="dropdown">
							<a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">
								Info
								<span class="caret"></span>
							</a>
							<ul class="dropdown-menu">
								<li ><a href="/rules">Rules</a></li>
								<li ><a href="/help">Help</a></li>
							</ul>
						</li>
						<li><a href="/?page=rss">RSS</a></li>
						<li><a href="https://twitter.com/NyaaV2">Twitter</a></li>
						<li><a href="//sukebei.nyaa.si">Fap</a></li>
					</ul>

					<ul class="nav navbar-nav navbar-right">
						<li class="dropdown">
							<a href="#" class="dropdown-toggle visible-lg visible-sm visible-xs" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">
								<i class="fa fa-user fa-fw"></i>
								Guest
								<span class="caret"></span>
							</a>
							<a href="#" class="dropdown-toggle hidden-lg hidden-sm hidden-xs" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">
								<i class="fa fa-user fa-fw"></i>
								<span class="caret"></span>
							</a>
							<ul class="dropdown-menu">
								<li>
									<a href="/login">
										<i class="fa fa-sign-in fa-fw"></i>
										Login
									</a>
								</li>
								<li>
									<a href="/register">
										<i class="fa fa-pencil fa-fw"></i>
										Register
									</a>
								</li>
							</ul>
						</li>
					</ul>



					<div class="search-container visible-xs visible-sm">
						<form class="navbar-form navbar-right form" action="/" method="get">

							<input type="text" class="form-control" name="q" placeholder="Search..." value="">
							<br>

							<select class="form-control" title="Filter" data-width="120px" name="f">
								<option value="0" title="No filter" selected>No filter</option>
								<option value="1" title="No remakes" >No remakes</option>
								<option value="2" title="Trusted only" >Trusted only</option>
							</select>

							<br>

							<select class="form-control" title="Category" data-width="200px" name="c">
								<option value="0_0" title="All categories" selected>
									All categories
								</option>
								<option value="1_0" title="Anime" >
									Anime
								</option>
								<option value="1_1" title="Anime - AMV" >
									- Anime Music Video
								</option>
								<option value="1_2" title="Anime - English" >
									- English-translated
								</option>
								<option value="1_3" title="Anime - Non-English" >
									- Non-English-translated
								</option>
								<option value="1_4" title="Anime - Raw" >
									- Raw
								</option>
								<option value="2_0" title="Audio" >
									Audio
								</option>
								<option value="2_1" title="Audio - Lossless" >
									- Lossless
								</option>
								<option value="2_2" title="Audio - Lossy" >
									- Lossy
								</option>
								<option value="3_0" title="Literature" >
									Literature
								</option>
								<option value="3_1" title="Literature - English" >
									- English-translated
								</option>
								<option value="3_2" title="Literature - Non-English" >
									- Non-English-translated
								</option>
								<option value="3_3" title="Literature - Raw" >
									- Raw
								</option>
								<option value="4_0" title="Live Action" >
									Live Action
								</option>
								<option value="4_1" title="Live Action - English" >
									- English-translated
								</option>
								<option value="4_2" title="Live Action - Idol/PV" >
									- Idol/Promotional Video
								</option>
								<option value="4_3" title="Live Action - Non-English" >
									- Non-English-translated
								</option>
								<option value="4_4" title="Live Action - Raw" >
									- Raw
								</option>
								<option value="5_0" title="Pictures" >
									Pictures
								</option>
								<option value="5_1" title="Pictures - Graphics" >
									- Graphics
								</option>
								<option value="5_2" title="Pictures - Photos" >
									- Photos
								</option>
								<option value="6_0" title="Software" >
									Software
								</option>
								<option value="6_1" title="Software - Apps" >
									- Applications
								</option>
								<option value="6_2" title="Software - Games" >
									- Games
								</option>
							</select>

							<br>

							<button class="btn btn-primary form-control" type="submit">
								<i class="fa fa-search fa-fw"></i> Search
							</button>
						</form>
					</div><!--/.search-container -->

					<form class="navbar-form navbar-right form" action="/" method="get">
						<div class="input-group search-container hidden-xs hidden-sm">
							<div class="input-group-btn nav-filter" id="navFilter-criteria">
								<select class="selectpicker show-tick" title="Filter" data-width="120px" name="f">
									<option value="0" title="No filter" selected>No filter</option>
									<option value="1" title="No remakes" >No remakes</option>
									<option value="2" title="Trusted only" >Trusted only</option>
								</select>
							</div>

							<div class="input-group-btn nav-filter" id="navFilter-category">
								<!--
									On narrow viewports, there isn't enough room to fit the full stuff in the selectpicker, so we show a full-width one on wide viewports, but squish it on narrow ones.
								-->
								<select class="selectpicker show-tick" title="Category" data-width="130px" name="c">
									<option value="0_0" title="All categories" selected>
										All categories
									</option>
									<option value="1_0" title="Anime" >
										Anime
									</option>
									<option value="1_1" title="Anime - AMV" >
										- Anime Music Video
									</option>
									<option value="1_2" title="Anime - English" >
										- English-translated
									</option>
									<option value="1_3" title="Anime - Non-English" >
										- Non-English-translated
									</option>
									<option value="1_4" title="Anime - Raw" >
										- Raw
									</option>
									<option value="2_0" title="Audio" >
										Audio
									</option>
									<option value="2_1" title="Audio - Lossless" >
										- Lossless
									</option>
									<option value="2_2" title="Audio - Lossy" >
										- Lossy
									</option>
									<option value="3_0" title="Literature" >
										Literature
									</option>
									<option value="3_1" title="Literature - English" >
										- English-translated
									</option>
									<option value="3_2" title="Literature - Non-English" >
										- Non-English-translated
									</option>
									<option value="3_3" title="Literature - Raw" >
										- Raw
									</option>
									<option value="4_0" title="Live Action" >
										Live Action
									</option>
									<option value="4_1" title="Live Action - English" >
										- English-translated
									</option>
									<option value="4_2" title="Live Action - Idol/PV" >
										- Idol/Promotional Video
									</option>
									<option value="4_3" title="Live Action - Non-English" >
										- Non-English-translated
									</option>
									<option value="4_4" title="Live Action - Raw" >
										- Raw
									</option>
									<option value="5_0" title="Pictures" >
										Pictures
									</option>
									<option value="5_1" title="Pictures - Graphics" >
										- Graphics
									</option>
									<option value="5_2" title="Pictures - Photos" >
										- Photos
									</option>
									<option value="6_0" title="Software" >
										Software
									</option>
									<option value="6_1" title="Software - Apps" >
										- Applications
									</option>
									<option value="6_2" title="Software - Games" >
										- Games
									</option>
								</select>
							</div>
							<input type="text" class="form-control search-bar" name="q" placeholder="Search..." value="" />
							<div class="input-group-btn search-btn">
								<button class="btn btn-primary" type="submit">
									<i class="fa fa-search fa-fw"></i>
								</button>
							</div>
						</div>
					</form>
				</div><!--/.nav-collapse -->
			</div><!--/.container -->
		</nav>

		<div class="container">





<div class="table-responsive">
	<table class="table table-bordered table-hover table-striped torrent-list">
		<thead>
			<tr>
				<th class="hdr-category text-center" style="width:80px;">Category</th>
				<th class="hdr-name" style="width:auto;">Name</th>
				<th class="hdr-comments sorting text-center" title="Comments" style="width:50px;"><a href="/?s=comments&amp;o=desc"></a><i class="fa fa-comments-o"></i></th>
				<th class="hdr-link text-center" style="width:70px;">Link</th>
				<th class="hdr-size sorting text-center" style="width:100px;"><a href="/?s=size&amp;o=desc"></a>Size</th>
				<th class="hdr-date sorting_desc text-center" title="In UTC" style="width:140px;"><a href="/?s=id&amp;o=asc"></a>Date</th>

				<th class="hdr-seeders sorting text-center" title="Seeders" style="width:50px;"><a href="/?s=seeders&amp;o=desc"></a><i class="fa fa-arrow-up" aria-hidden="true"></i></th>
				<th class="hdr-leechers sorting text-center" title="Leechers" style="width:50px;"><a href="/?s=leechers&amp;o=desc"></a><i class="fa fa-arrow-down" aria-hidden="true"></i></th>
				<th class="hdr-downloads sorting text-center" title="Completed downloads" style="width:50px;"><a href="/?s=downloads&amp;o=desc"></a><i class="fa fa-check" aria-hidden="true"></i></th>
			</tr>
		</thead>
		<tbody>
			<tr class="success">
				<td>
					<a href="/?c=1_2" title="Anime - English-translated">
						<img src="/static/img/icons/nyaa/1_2.png" alt="Anime - English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494359" title="[SubsPlease] Fantasy Bishoujo Juniku Ojisan to - 07 (1080p) [58949588].mkv">[SubsPlease] Fantasy Bishoujo Juniku Ojisan to - 07 (1080p) [58949588].mkv</a>
				</td>
				<td class="text-center">
					<a href="/download/1494359.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:1a89365b6b9ac72926245ae80fc0ae520f3e47dc&amp;dn=%5BSubsPlease%5D%20Fantasy%20Bishoujo%20Juniku%20Ojisan%20to%20-%2007%20%281080p%29%20%5B58949588%5D.mkv&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">1.4 GiB</td>
				<td class="text-center" data-timestamp="1645551116">2022-02-22 17:31</td>

				<td class="text-center">14</td>
				<td class="text-center">148</td>
				<td class="text-center">14</td>
			</tr>
			<tr class="success">
				<td>
					<a href="/?c=1_2" title="Anime - English-translated">
						<img src="/static/img/icons/nyaa/1_2.png" alt="Anime - English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494358" title="[SubsPlease] Fantasy Bishoujo Juniku Ojisan to - 07 (720p) [E8FFE27E].mkv">[SubsPlease] Fantasy Bishoujo Juniku Ojisan to - 07 (720p) [E8FFE27E].mkv</a>
				</td>
				<td class="text-center">
					<a href="/download/1494358.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:bdb1ad40dc1cb86f67ef718ef1de9b2b2f84b233&amp;dn=%5BSubsPlease%5D%20Fantasy%20Bishoujo%20Juniku%20Ojisan%20to%20-%2007%20%28720p%29%20%5BE8FFE27E%5D.mkv&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">707.2 MiB</td>
				<td class="text-center" data-timestamp="1645551081">2022-02-22 17:31</td>

				<td class="text-center">18</td>
				<td class="text-center">94</td>
				<td class="text-center">13</td>
			</tr>
			<tr class="success">
				<td>
					<a href="/?c=1_2" title="Anime - English-translated">
						<img src="/static/img/icons/nyaa/1_2.png" alt="Anime - English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494357" title="[SubsPlease] Fantasy Bishoujo Juniku Ojisan to - 07 (480p) [462897C5].mkv">[SubsPlease] Fantasy Bishoujo Juniku Ojisan to - 07 (480p) [462897C5].mkv</a>
				</td>
				<td class="text-center">
					<a href="/download/1494357.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:e020acc9512744977aedaf84bd240da197575351&amp;dn=%5BSubsPlease%5D%20Fantasy%20Bishoujo%20Juniku%20Ojisan%20to%20-%2007%20%28480p%29%20%5B462897C5%5D.mkv&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">362.1 MiB</td>
				<td class="text-center" data-timestamp="1645551072">2022-02-22 17:31</td>

				<td class="text-center">9</td>
				<td class="text-center">29</td>
				<td class="text-center">12</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=3_2" title="Literature - Non-English-translated">
						<img src="/static/img/icons/nyaa/3_2.png" alt="Literature - Non-English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494356" title="[Mangá Yuri] PT BR | Hitozuma to JK - Volume 1 (Lonely Lily)">[Mangá Yuri] PT BR | Hitozuma to JK - Volume 1 (Lonely Lily)</a>
				</td>
				<td class="text-center">
					<a href="/download/1494356.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:f91a85a9be754a1a645949fa04a2983869120362&amp;dn=%5BMang%C3%A1%20Yuri%5D%20PT%20BR%20%7C%20Hitozuma%20to%20JK%20-%20Volume%201%20%28Lonely%20Lily%29&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">178.0 MiB</td>
				<td class="text-center" data-timestamp="1645550919">2022-02-22 17:28</td>

				<td class="text-center">0</td>
				<td class="text-center">2</td>
				<td class="text-center">0</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=4_4" title="Live Action - Raw">
						<img src="/static/img/icons/nyaa/4_4.png" alt="Live Action - Raw" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494355" title="【大河ドラマアンコール】黄金の日日（４５）「天変地異」 - [1440x1080p.hevc.mkv][字]">【大河ドラマアンコール】黄金の日日（４５）「天変地異」 - [1440x1080p.hevc.mkv][字]</a>
				</td>
				<td class="text-center">
					<a href="/download/1494355.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:3e3f85234670dedda3c67e98e91f30a95fd8c6eb&amp;dn=%E3%80%90%E5%A4%A7%E6%B2%B3%E3%83%89%E3%83%A9%E3%83%9E%E3%82%A2%E3%83%B3%E3%82%B3%E3%83%BC%E3%83%AB%E3%80%91%E9%BB%84%E9%87%91%E3%81%AE%E6%97%A5%E6%97%A5%EF%BC%88%EF%BC%94%EF%BC%95%EF%BC%89%E3%80%8C%E5%A4%A9%E5%A4%89%E5%9C%B0%E7%95%B0%E3%80%8D%20-%20%5B1440x1080p.hevc.mkv%5D%5B%E5%AD%97%5D&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">589.7 MiB</td>
				<td class="text-center" data-timestamp="1645549442">2022-02-22 17:04</td>

				<td class="text-center">2</td>
				<td class="text-center">23</td>
				<td class="text-center">0</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=4_4" title="Live Action - Raw">
						<img src="/static/img/icons/nyaa/4_4.png" alt="Live Action - Raw" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494354" title="【大河ドラマアンコール】黄金の日日（４４）「呂宋遠征計画」 - [1440x1080p.hevc.mkv][字]">【大河ドラマアンコール】黄金の日日（４４）「呂宋遠征計画」 - [1440x1080p.hevc.mkv][字]</a>
				</td>
				<td class="text-center">
					<a href="/download/1494354.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:960f4750118ed3183225a55ed097df16c8d1f2ed&amp;dn=%E3%80%90%E5%A4%A7%E6%B2%B3%E3%83%89%E3%83%A9%E3%83%9E%E3%82%A2%E3%83%B3%E3%82%B3%E3%83%BC%E3%83%AB%E3%80%91%E9%BB%84%E9%87%91%E3%81%AE%E6%97%A5%E6%97%A5%EF%BC%88%EF%BC%94%EF%BC%94%EF%BC%89%E3%80%8C%E5%91%82%E5%AE%8B%E9%81%A0%E5%BE%81%E8%A8%88%E7%94%BB%E3%80%8D%20-%20%5B1440x1080p.hevc.mkv%5D%5B%E5%AD%97%5D&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">811.0 MiB</td>
				<td class="text-center" data-timestamp="1645549427">2022-02-22 17:03</td>

				<td class="text-center">1</td>
				<td class="text-center">21</td>
				<td class="text-center">0</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=4_4" title="Live Action - Raw">
						<img src="/static/img/icons/nyaa/4_4.png" alt="Live Action - Raw" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494353" title="【大河ドラマ】鎌倉殿の１３人（７）「敵か、あるいは」 - [1920x1080p.hevc.mkv][字]">【大河ドラマ】鎌倉殿の１３人（７）「敵か、あるいは」 - [1920x1080p.hevc.mkv][字]</a>
				</td>
				<td class="text-center">
					<a href="/download/1494353.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:58fdea442d943f771b14842fe51924f94ec7b297&amp;dn=%E3%80%90%E5%A4%A7%E6%B2%B3%E3%83%89%E3%83%A9%E3%83%9E%E3%80%91%E9%8E%8C%E5%80%89%E6%AE%BF%E3%81%AE%EF%BC%91%EF%BC%93%E4%BA%BA%EF%BC%88%EF%BC%97%EF%BC%89%E3%80%8C%E6%95%B5%E3%81%8B%E3%80%81%E3%81%82%E3%82%8B%E3%81%84%E3%81%AF%E3%80%8D%20-%20%5B1920x1080p.hevc.mkv%5D%5B%E5%AD%97%5D&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">850.0 MiB</td>
				<td class="text-center" data-timestamp="1645549309">2022-02-22 17:01</td>

				<td class="text-center">1</td>
				<td class="text-center">78</td>
				<td class="text-center">0</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_3" title="Anime - Non-English-translated">
						<img src="/static/img/icons/nyaa/1_3.png" alt="Anime - Non-English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494352" title="Baraou no Souretsu - 06 (WEBDL x265 1080p AAC) Ukr DVO SUB">Baraou no Souretsu - 06 (WEBDL x265 1080p AAC) Ukr DVO SUB</a>
				</td>
				<td class="text-center">
					<a href="/download/1494352.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:82e0b153f3c8873f97af72651ec939091cc1b6e6&amp;dn=Baraou%20no%20Souretsu%20-%2006%20%28WEBDL%20x265%201080p%20AAC%29%20Ukr%20DVO%20SUB&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">311.0 MiB</td>
				<td class="text-center" data-timestamp="1645548704">2022-02-22 16:51</td>

				<td class="text-center">6</td>
				<td class="text-center">1</td>
				<td class="text-center">7</td>
			</tr>
			<tr class="danger">
				<td>
					<a href="/?c=1_2" title="Anime - English-translated">
						<img src="/static/img/icons/nyaa/1_2.png" alt="Anime - English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494351" title="[EMBER] Tensai Ouji no Akaji Kokka Saisei Jutsu S01E07 [1080p] [HEVC WEBRip] (The Genius Prince&#39;s Guide to Raising a Nation Out of Debt)">[EMBER] Tensai Ouji no Akaji Kokka Saisei Jutsu S01E07 [1080p] [HEVC WEBRip] (The Genius Prince&#39;s Guide to Raising a Nation Out of Debt)</a>
				</td>
				<td class="text-center">
					<a href="/download/1494351.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:f5150e81929e2673f10fc26d68f2d319c625ffaa&amp;dn=%5BEMBER%5D%20Tensai%20Ouji%20no%20Akaji%20Kokka%20Saisei%20Jutsu%20S01E07%20%5B1080p%5D%20%5BHEVC%20WEBRip%5D%20%28The%20Genius%20Prince%27s%20Guide%20to%20Raising%20a%20Nation%20Out%20of%20Debt%29&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">326.7 MiB</td>
				<td class="text-center" data-timestamp="1645547910">2022-02-22 16:38</td>

				<td class="text-center">96</td>
				<td class="text-center">22</td>
				<td class="text-center">102</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_2" title="Anime - English-translated">
						<img src="/static/img/icons/nyaa/1_2.png" alt="Anime - English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494350#comments" class="comments" title="2 comments">
						<i class="fa fa-comments-o"></i>2</a>
					<a href="/view/1494350" title="ミラベルと魔法だらけの家">ミラベルと魔法だらけの家</a>
				</td>
				<td class="text-center">
					<a href="/download/1494350.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:b0f7257e4d82026fc66784f95e0dd911bdd33631&amp;dn=%E3%83%9F%E3%83%A9%E3%83%99%E3%83%AB%E3%81%A8%E9%AD%94%E6%B3%95%E3%81%A0%E3%82%89%E3%81%91%E3%81%AE%E5%AE%B6&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">6.6 GiB</td>
				<td class="text-center" data-timestamp="1645547688">2022-02-22 16:34</td>

				<td class="text-center">6</td>
				<td class="text-center">19</td>
				<td class="text-center">7</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_2" title="Anime - English-translated">
						<img src="/static/img/icons/nyaa/1_2.png" alt="Anime - English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494349" title="[Yameii] Deep Insanity The Lost Child - 07 [English Dub] [WEB-DL 1080p] [5CD0A02C]">[Yameii] Deep Insanity The Lost Child - 07 [English Dub] [WEB-DL 1080p] [5CD0A02C]</a>
				</td>
				<td class="text-center">
					<a href="/download/1494349.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:2dcf80985b4a01cdada2a8aa99e2619314e182f5&amp;dn=%5BYameii%5D%20Deep%20Insanity%20The%20Lost%20Child%20-%2007%20%5BEnglish%20Dub%5D%20%5BWEB-DL%201080p%5D%20%5B5CD0A02C%5D&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">1.1 GiB</td>
				<td class="text-center" data-timestamp="1645547561">2022-02-22 16:32</td>

				<td class="text-center">41</td>
				<td class="text-center">17</td>
				<td class="text-center">42</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_2" title="Anime - English-translated">
						<img src="/static/img/icons/nyaa/1_2.png" alt="Anime - English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494348" title="[Yameii] Deep Insanity The Lost Child - 07 [English Dub] [WEB-DL 720p] [42386190]">[Yameii] Deep Insanity The Lost Child - 07 [English Dub] [WEB-DL 720p] [42386190]</a>
				</td>
				<td class="text-center">
					<a href="/download/1494348.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:09eb99951a50f4d045b6eac8fecdf6828185eddc&amp;dn=%5BYameii%5D%20Deep%20Insanity%20The%20Lost%20Child%20-%2007%20%5BEnglish%20Dub%5D%20%5BWEB-DL%20720p%5D%20%5B42386190%5D&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">426.3 MiB</td>
				<td class="text-center" data-timestamp="1645547525">2022-02-22 16:32</td>

				<td class="text-center">20</td>
				<td class="text-center">9</td>
				<td class="text-center">16</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_2" title="Anime - English-translated">
						<img src="/static/img/icons/nyaa/1_2.png" alt="Anime - English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494347" title="Deep Insanity - The Lost Child - S01E07 - ENG 1080p WEB H.264 -NanDesuKa (FUNi).mkv">Deep Insanity - The Lost Child - S01E07 - ENG 1080p WEB H.264 -NanDesuKa (FUNi).mkv</a>
				</td>
				<td class="text-center">
					<a href="/download/1494347.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:444f0b425446fa219d3787c81eee20ffad4bc903&amp;dn=Deep%20Insanity%20-%20The%20Lost%20Child%20-%20S01E07%20-%20ENG%201080p%20WEB%20H.264%20-NanDesuKa%20%28FUNi%29.mkv&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">1.1 GiB</td>
				<td class="text-center" data-timestamp="1645547449">2022-02-22 16:30</td>

				<td class="text-center">10</td>
				<td class="text-center">11</td>
				<td class="text-center">14</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_2" title="Anime - English-translated">
						<img src="/static/img/icons/nyaa/1_2.png" alt="Anime - English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494346" title="Deep Insanity - The Lost Child - S01E07 - ENG 720p WEB H.264 -NanDesuKa (FUNi).mkv">Deep Insanity - The Lost Child - S01E07 - ENG 720p WEB H.264 -NanDesuKa (FUNi).mkv</a>
				</td>
				<td class="text-center">
					<a href="/download/1494346.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:c3ccc894884411d84b756c42e948ada222350eee&amp;dn=Deep%20Insanity%20-%20The%20Lost%20Child%20-%20S01E07%20-%20ENG%20720p%20WEB%20H.264%20-NanDesuKa%20%28FUNi%29.mkv&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">650.3 MiB</td>
				<td class="text-center" data-timestamp="1645547441">2022-02-22 16:30</td>

				<td class="text-center">9</td>
				<td class="text-center">5</td>
				<td class="text-center">11</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_2" title="Anime - English-translated">
						<img src="/static/img/icons/nyaa/1_2.png" alt="Anime - English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494345" title="Deep Insanity - The Lost Child - S01E07 - ENG 540p WEB H.264 -NanDesuKa (FUNi).mkv">Deep Insanity - The Lost Child - S01E07 - ENG 540p WEB H.264 -NanDesuKa (FUNi).mkv</a>
				</td>
				<td class="text-center">
					<a href="/download/1494345.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:18af6a650877e0fddfba03025a8eb35529abcf70&amp;dn=Deep%20Insanity%20-%20The%20Lost%20Child%20-%20S01E07%20-%20ENG%20540p%20WEB%20H.264%20-NanDesuKa%20%28FUNi%29.mkv&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">351.3 MiB</td>
				<td class="text-center" data-timestamp="1645547422">2022-02-22 16:30</td>

				<td class="text-center">4</td>
				<td class="text-center">2</td>
				<td class="text-center">8</td>
			</tr>
			<tr class="danger">
				<td>
					<a href="/?c=1_2" title="Anime - English-translated">
						<img src="/static/img/icons/nyaa/1_2.png" alt="Anime - English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494344" title="[YuiSubs] Kenja no Deshi wo Nanoru Kenja - 07  (x265 H.265 1080p)">[YuiSubs] Kenja no Deshi wo Nanoru Kenja - 07  (x265 H.265 1080p)</a>
				</td>
				<td class="text-center">
					<a href="/download/1494344.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:78c8eb37f1275e782c8ebc4ee523813c5972a5ee&amp;dn=%5BYuiSubs%5D%20Kenja%20no%20Deshi%20wo%20Nanoru%20Kenja%20-%2007%20%20%28x265%20H.265%201080p%29&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">337.5 MiB</td>
				<td class="text-center" data-timestamp="1645546586">2022-02-22 16:16</td>

				<td class="text-center">42</td>
				<td class="text-center">7</td>
				<td class="text-center">42</td>
			</tr>
			<tr class="danger">
				<td>
					<a href="/?c=1_2" title="Anime - English-translated">
						<img src="/static/img/icons/nyaa/1_2.png" alt="Anime - English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494343" title="[ASW] Kenja no Deshi wo Nanoru Kenja - 07 [1080p HEVC x265 10Bit][AAC]">[ASW] Kenja no Deshi wo Nanoru Kenja - 07 [1080p HEVC x265 10Bit][AAC]</a>
				</td>
				<td class="text-center">
					<a href="/download/1494343.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:c42a7a1a40ca7c72ef3403ed48c930d1b575c878&amp;dn=%5BASW%5D%20Kenja%20no%20Deshi%20wo%20Nanoru%20Kenja%20-%2007%20%5B1080p%20HEVC%20x265%2010Bit%5D%5BAAC%5D&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">271.4 MiB</td>
				<td class="text-center" data-timestamp="1645546473">2022-02-22 16:14</td>

				<td class="text-center">193</td>
				<td class="text-center">45</td>
				<td class="text-center">243</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_4" title="Anime - Raw">
						<img src="/static/img/icons/nyaa/1_4.png" alt="Anime - Raw" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494342" title="[Ohys-Raws] Kenja no Deshi o Nanoru Kenja - 07 (BS4 1280x720 x264 AAC).mp4">[Ohys-Raws] Kenja no Deshi o Nanoru Kenja - 07 (BS4 1280x720 x264 AAC).mp4</a>
				</td>
				<td class="text-center">
					<a href="/download/1494342.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:0c9ed323f8aada85b628e264dcff453f6de5ef2c&amp;dn=%5BOhys-Raws%5D%20Kenja%20no%20Deshi%20o%20Nanoru%20Kenja%20-%2007%20%28BS4%201280x720%20x264%20AAC%29.mp4&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">280.9 MiB</td>
				<td class="text-center" data-timestamp="1645545939">2022-02-22 16:05</td>

				<td class="text-center">372</td>
				<td class="text-center">159</td>
				<td class="text-center">585</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_2" title="Anime - English-translated">
						<img src="/static/img/icons/nyaa/1_2.png" alt="Anime - English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494341#comments" class="comments" title="1 comment">
						<i class="fa fa-comments-o"></i>1</a>
					<a href="/view/1494341" title="[Mad le Zisell] Tensai Ouji no Akaji Kokka Saisei Jutsu - 07 [720p].mkv">[Mad le Zisell] Tensai Ouji no Akaji Kokka Saisei Jutsu - 07 [720p].mkv</a>
				</td>
				<td class="text-center">
					<a href="/download/1494341.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:d2bc95a7c1d56ca7c827e2204a9ddc4c8cdcb09d&amp;dn=%5BMad%20le%20Zisell%5D%20Tensai%20Ouji%20no%20Akaji%20Kokka%20Saisei%20Jutsu%20-%2007%20%5B720p%5D.mkv&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">220.0 MiB</td>
				<td class="text-center" data-timestamp="1645545460">2022-02-22 15:57</td>

				<td class="text-center">9</td>
				<td class="text-center">1</td>
				<td class="text-center">10</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_3" title="Anime - Non-English-translated">
						<img src="/static/img/icons/nyaa/1_3.png" alt="Anime - Non-English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494340" title="Tensai Ouji no Akaji Kokka Saiseijutsu | Руководство гениального принца по избавлению нации от долгов [ТV-1] [2022] [07 of 12] [WEBRip] [1080p] [RUS + JAP]">Tensai Ouji no Akaji Kokka Saiseijutsu | Руководство гениального принца по избавлению нации от долгов [ТV-1] [2022] [07 of 12] [WEBRip] [1080p] [RUS + JAP]</a>
				</td>
				<td class="text-center">
					<a href="/download/1494340.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:e387a8b2e499a731edf2ac4debf56afa5ee5b234&amp;dn=Tensai%20Ouji%20no%20Akaji%20Kokka%20Saiseijutsu%20%7C%20%D0%A0%D1%83%D0%BA%D0%BE%D0%B2%D0%BE%D0%B4%D1%81%D1%82%D0%B2%D0%BE%20%D0%B3%D0%B5%D0%BD%D0%B8%D0%B0%D0%BB%D1%8C%D0%BD%D0%BE%D0%B3%D0%BE%20%D0%BF%D1%80%D0%B8%D0%BD%D1%86%D0%B0%20%D0%BF%D0%BE%20%D0%B8%D0%B7%D0%B1%D0%B0%D0%B2%D0%BB%D0%B5%D0%BD%D0%B8%D1%8E%20%D0%BD%D0%B0%D1%86%D0%B8%D0%B8%20%D0%BE%D1%82%20%D0%B4%D0%BE%D0%BB%D0%B3%D0%BE%D0%B2%20%5B%D0%A2V-1%5D%20%5B2022%5D%20%5B07%20of%2012%5D%20%5BWEBRip%5D%20%5B1080p%5D%20%5BRUS%20%2B%20JAP%5D&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">5.4 GiB</td>
				<td class="text-center" data-timestamp="1645545455">2022-02-22 15:57</td>

				<td class="text-center">2</td>
				<td class="text-center">2</td>
				<td class="text-center">1</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_3" title="Anime - Non-English-translated">
						<img src="/static/img/icons/nyaa/1_3.png" alt="Anime - Non-English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494339" title="Kenja no Deshi o Nanoru Kenja | Она представилась учеником мудреца [ТV-1] [2022] [07 of 12] [WEBRip] [1080p] [RUS + JAP]">Kenja no Deshi o Nanoru Kenja | Она представилась учеником мудреца [ТV-1] [2022] [07 of 12] [WEBRip] [1080p] [RUS + JAP]</a>
				</td>
				<td class="text-center">
					<a href="/download/1494339.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:1b00273a8519004670d837f3116e718a2bc41800&amp;dn=Kenja%20no%20Deshi%20o%20Nanoru%20Kenja%20%7C%20%D0%9E%D0%BD%D0%B0%20%D0%BF%D1%80%D0%B5%D0%B4%D1%81%D1%82%D0%B0%D0%B2%D0%B8%D0%BB%D0%B0%D1%81%D1%8C%20%D1%83%D1%87%D0%B5%D0%BD%D0%B8%D0%BA%D0%BE%D0%BC%20%D0%BC%D1%83%D0%B4%D1%80%D0%B5%D1%86%D0%B0%20%5B%D0%A2V-1%5D%20%5B2022%5D%20%5B07%20of%2012%5D%20%5BWEBRip%5D%20%5B1080p%5D%20%5BRUS%20%2B%20JAP%5D&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">5.4 GiB</td>
				<td class="text-center" data-timestamp="1645545420">2022-02-22 15:57</td>

				<td class="text-center">1</td>
				<td class="text-center">2</td>
				<td class="text-center">0</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_3" title="Anime - Non-English-translated">
						<img src="/static/img/icons/nyaa/1_3.png" alt="Anime - Non-English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494338" title="[酷漫404][鬼滅之刃 遊郭篇][01-11][1080P][WebRip][繁日雙語(內嵌+外掛)][HEVC-10bit AAC][MKV][字幕組招人內詳]">[酷漫404][鬼滅之刃 遊郭篇][01-11][1080P][WebRip][繁日雙語(內嵌+外掛)][HEVC-10bit AAC][MKV][字幕組招人內詳]</a>
				</td>
				<td class="text-center">
					<a href="/download/1494338.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:cdacd908cbc17ade7bc60cf25904c4751758304c&amp;dn=%5B%E9%85%B7%E6%BC%AB404%5D%5B%E9%AC%BC%E6%BB%85%E4%B9%8B%E5%88%83%20%E9%81%8A%E9%83%AD%E7%AF%87%5D%5B01-11%5D%5B1080P%5D%5BWebRip%5D%5B%E7%B9%81%E6%97%A5%E9%9B%99%E8%AA%9E%28%E5%85%A7%E5%B5%8C%2B%E5%A4%96%E6%8E%9B%29%5D%5BHEVC-10bit%20AAC%5D%5BMKV%5D%5B%E5%AD%97%E5%B9%95%E7%B5%84%E6%8B%9B%E4%BA%BA%E5%85%A7%E8%A9%B3%5D&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">7.7 GiB</td>
				<td class="text-center" data-timestamp="1645545127">2022-02-22 15:52</td>

				<td class="text-center">14</td>
				<td class="text-center">46</td>
				<td class="text-center">11</td>
			</tr>
			<tr class="success">
				<td>
					<a href="/?c=1_3" title="Anime - Non-English-translated">
						<img src="/static/img/icons/nyaa/1_3.png" alt="Anime - Non-English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494337" title="[PuyaSubs!] Kenja no Deshi wo Nanoru Kenja - 07 [ESP-ENG][720p][56A82F7E].mkv">[PuyaSubs!] Kenja no Deshi wo Nanoru Kenja - 07 [ESP-ENG][720p][56A82F7E].mkv</a>
				</td>
				<td class="text-center">
					<a href="/download/1494337.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:5246e30e23b4eaef00e9f2b255d70915220f1222&amp;dn=%5BPuyaSubs%21%5D%20Kenja%20no%20Deshi%20wo%20Nanoru%20Kenja%20-%2007%20%5BESP-ENG%5D%5B720p%5D%5B56A82F7E%5D.mkv&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">629.5 MiB</td>
				<td class="text-center" data-timestamp="1645545084">2022-02-22 15:51</td>

				<td class="text-center">36</td>
				<td class="text-center">10</td>
				<td class="text-center">31</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_3" title="Anime - Non-English-translated">
						<img src="/static/img/icons/nyaa/1_3.png" alt="Anime - Non-English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494336" title="[酷漫404][鬼灭之刃 游郭篇][01-11][1080P][WebRip][简日双语(内嵌+外挂)][HEVC-10bit AAC][MKV][字幕组招人内详]">[酷漫404][鬼灭之刃 游郭篇][01-11][1080P][WebRip][简日双语(内嵌+外挂)][HEVC-10bit AAC][MKV][字幕组招人内详]</a>
				</td>
				<td class="text-center">
					<a href="/download/1494336.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:71b849c72d1d922cff9e71a1bb006aa2e99942ff&amp;dn=%5B%E9%85%B7%E6%BC%AB404%5D%5B%E9%AC%BC%E7%81%AD%E4%B9%8B%E5%88%83%20%E6%B8%B8%E9%83%AD%E7%AF%87%5D%5B01-11%5D%5B1080P%5D%5BWebRip%5D%5B%E7%AE%80%E6%97%A5%E5%8F%8C%E8%AF%AD%28%E5%86%85%E5%B5%8C%2B%E5%A4%96%E6%8C%82%29%5D%5BHEVC-10bit%20AAC%5D%5BMKV%5D%5B%E5%AD%97%E5%B9%95%E7%BB%84%E6%8B%9B%E4%BA%BA%E5%86%85%E8%AF%A6%5D&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">7.7 GiB</td>
				<td class="text-center" data-timestamp="1645544992">2022-02-22 15:49</td>

				<td class="text-center">6</td>
				<td class="text-center">67</td>
				<td class="text-center">0</td>
			</tr>
			<tr class="success">
				<td>
					<a href="/?c=1_3" title="Anime - Non-English-translated">
						<img src="/static/img/icons/nyaa/1_3.png" alt="Anime - Non-English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494335" title="[PuyaSubs!] Kenja no Deshi wo Nanoru Kenja - 07 [ESP-ENG][1080p][49D62D50].mkv">[PuyaSubs!] Kenja no Deshi wo Nanoru Kenja - 07 [ESP-ENG][1080p][49D62D50].mkv</a>
				</td>
				<td class="text-center">
					<a href="/download/1494335.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:a8527667c11b399f8491aa930a980ca4e32be145&amp;dn=%5BPuyaSubs%21%5D%20Kenja%20no%20Deshi%20wo%20Nanoru%20Kenja%20-%2007%20%5BESP-ENG%5D%5B1080p%5D%5B49D62D50%5D.mkv&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">1.0 GiB</td>
				<td class="text-center" data-timestamp="1645544898">2022-02-22 15:48</td>

				<td class="text-center">33</td>
				<td class="text-center">10</td>
				<td class="text-center">24</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_3" title="Anime - Non-English-translated">
						<img src="/static/img/icons/nyaa/1_3.png" alt="Anime - Non-English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494334" title="[ES] Karakai Jouzu no Takagi-san 3 - 07 [1080p] [4457ADD3]">[ES] Karakai Jouzu no Takagi-san 3 - 07 [1080p] [4457ADD3]</a>
				</td>
				<td class="text-center">
					<a href="/download/1494334.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:e84cbbac0ac95cd4394ba8b025ae8a65c15da1bf&amp;dn=%5BES%5D%20Karakai%20Jouzu%20no%20Takagi-san%203%20-%2007%20%5B1080p%5D%20%5B4457ADD3%5D&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">672.7 MiB</td>
				<td class="text-center" data-timestamp="1645544833">2022-02-22 15:47</td>

				<td class="text-center">18</td>
				<td class="text-center">8</td>
				<td class="text-center">19</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_3" title="Anime - Non-English-translated">
						<img src="/static/img/icons/nyaa/1_3.png" alt="Anime - Non-English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494333" title="[ES] Karakai Jouzu no Takagi-san 3 - 06 [1080p] [EA480A4F]">[ES] Karakai Jouzu no Takagi-san 3 - 06 [1080p] [EA480A4F]</a>
				</td>
				<td class="text-center">
					<a href="/download/1494333.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:1249bd750f27b867a005d1902d874ce319042ef9&amp;dn=%5BES%5D%20Karakai%20Jouzu%20no%20Takagi-san%203%20-%2006%20%5B1080p%5D%20%5BEA480A4F%5D&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">676.5 MiB</td>
				<td class="text-center" data-timestamp="1645544802">2022-02-22 15:46</td>

				<td class="text-center">22</td>
				<td class="text-center">8</td>
				<td class="text-center">17</td>
			</tr>
			<tr class="danger">
				<td>
					<a href="/?c=1_2" title="Anime - English-translated">
						<img src="/static/img/icons/nyaa/1_2.png" alt="Anime - English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494332" title="[LBA]Tensai.Ouji.07.x264.AAC.720p.HARDSUB.mp4">[LBA]Tensai.Ouji.07.x264.AAC.720p.HARDSUB.mp4</a>
				</td>
				<td class="text-center">
					<a href="/download/1494332.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:f0dc513f8af89fbb8f8aa86d5b708be1fa3f6bcc&amp;dn=%5BLBA%5DTensai.Ouji.07.x264.AAC.720p.HARDSUB.mp4&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">90.1 MiB</td>
				<td class="text-center" data-timestamp="1645544690">2022-02-22 15:44</td>

				<td class="text-center">3</td>
				<td class="text-center">4</td>
				<td class="text-center">8</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_3" title="Anime - Non-English-translated">
						<img src="/static/img/icons/nyaa/1_3.png" alt="Anime - Non-English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494331" title="[DJB] She Professed Herself Pupil of the Wise Man - 07 VOSTFR [1080p].mp4">[DJB] She Professed Herself Pupil of the Wise Man - 07 VOSTFR [1080p].mp4</a>
				</td>
				<td class="text-center">
					<a href="/download/1494331.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:b28775c03f0f65aceb137a3a979a7535ce838ec7&amp;dn=%5BDJB%5D%20She%20Professed%20Herself%20Pupil%20of%20the%20Wise%20Man%20-%2007%20VOSTFR%20%5B1080p%5D.mp4&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">336.9 MiB</td>
				<td class="text-center" data-timestamp="1645544635">2022-02-22 15:43</td>

				<td class="text-center">19</td>
				<td class="text-center">6</td>
				<td class="text-center">22</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_3" title="Anime - Non-English-translated">
						<img src="/static/img/icons/nyaa/1_3.png" alt="Anime - Non-English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494330" title="[悠哈璃羽字幕社][公主连结 Re:Dive 第二季_Princess Connect!! Re Dive S2][06][x264 1080p][CHT]">[悠哈璃羽字幕社][公主连结 Re:Dive 第二季_Princess Connect!! Re Dive S2][06][x264 1080p][CHT]</a>
				</td>
				<td class="text-center">
					<a href="/download/1494330.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:1310ad565b4d960db9ad29a79fa6029494067444&amp;dn=%5B%E6%82%A0%E5%93%88%E7%92%83%E7%BE%BD%E5%AD%97%E5%B9%95%E7%A4%BE%5D%5B%E5%85%AC%E4%B8%BB%E8%BF%9E%E7%BB%93%20Re%3ADive%20%E7%AC%AC%E4%BA%8C%E5%AD%A3_Princess%20Connect%21%21%20Re%20Dive%20S2%5D%5B06%5D%5Bx264%201080p%5D%5BCHT%5D&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">327.0 MiB</td>
				<td class="text-center" data-timestamp="1645544320">2022-02-22 15:38</td>

				<td class="text-center">21</td>
				<td class="text-center">9</td>
				<td class="text-center">36</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_3" title="Anime - Non-English-translated">
						<img src="/static/img/icons/nyaa/1_3.png" alt="Anime - Non-English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494329" title="[悠哈璃羽字幕社][公主连结 Re:Dive 第二季_Princess Connect!! Re Dive S2][06][x264 1080p][CHS]">[悠哈璃羽字幕社][公主连结 Re:Dive 第二季_Princess Connect!! Re Dive S2][06][x264 1080p][CHS]</a>
				</td>
				<td class="text-center">
					<a href="/download/1494329.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:452bef62a66fc37191e2e9e720014b9b4f43b3bf&amp;dn=%5B%E6%82%A0%E5%93%88%E7%92%83%E7%BE%BD%E5%AD%97%E5%B9%95%E7%A4%BE%5D%5B%E5%85%AC%E4%B8%BB%E8%BF%9E%E7%BB%93%20Re%3ADive%20%E7%AC%AC%E4%BA%8C%E5%AD%A3_Princess%20Connect%21%21%20Re%20Dive%20S2%5D%5B06%5D%5Bx264%201080p%5D%5BCHS%5D&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">327.3 MiB</td>
				<td class="text-center" data-timestamp="1645544248">2022-02-22 15:37</td>

				<td class="text-center">9</td>
				<td class="text-center">5</td>
				<td class="text-center">13</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_2" title="Anime - English-translated">
						<img src="/static/img/icons/nyaa/1_2.png" alt="Anime - English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494328" title="[Yameii] Kenja no Deshi o Nanoru Kenja - 07 | She Professed Herself Pupil of the Wise Man [English Dub] [WEB-DL 1080p] [75CC097D]">[Yameii] Kenja no Deshi o Nanoru Kenja - 07 | She Professed Herself Pupil of the Wise Man [English Dub] [WEB-DL 1080p] [75CC097D]</a>
				</td>
				<td class="text-center">
					<a href="/download/1494328.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:e88f8b702a1fcdea9ec02775956c705275af2fd9&amp;dn=%5BYameii%5D%20Kenja%20no%20Deshi%20o%20Nanoru%20Kenja%20-%2007%20%7C%20She%20Professed%20Herself%20Pupil%20of%20the%20Wise%20Man%20%5BEnglish%20Dub%5D%20%5BWEB-DL%201080p%5D%20%5B75CC097D%5D&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">1.0 GiB</td>
				<td class="text-center" data-timestamp="1645544242">2022-02-22 15:37</td>

				<td class="text-center">103</td>
				<td class="text-center">20</td>
				<td class="text-center">135</td>
			</tr>
			<tr class="danger">
				<td>
					<a href="/?c=1_2" title="Anime - English-translated">
						<img src="/static/img/icons/nyaa/1_2.png" alt="Anime - English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494327" title="[Anime Time] Tensai Ouji no Akaji Kokka Saisei Jutsu - 07 [1080p][HEVC 10bit x265][AAC][Multi Sub] The Genius Prince&#39;s Guide to Raising a Nation Out of Debt.mkv [Weekly]">[Anime Time] Tensai Ouji no Akaji Kokka Saisei Jutsu - 07 [1080p][HEVC 10bit x265][AAC][Multi Sub] The Genius Prince&#39;s Guide to Raising a Nation Out of Debt.mkv [Weekly]</a>
				</td>
				<td class="text-center">
					<a href="/download/1494327.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:93b227cbdbc710c2e1d562f1a160e45dd6a86ed4&amp;dn=%5BAnime%20Time%5D%20Tensai%20Ouji%20no%20Akaji%20Kokka%20Saisei%20Jutsu%20-%2007%20%5B1080p%5D%5BHEVC%2010bit%20x265%5D%5BAAC%5D%5BMulti%20Sub%5D%20The%20Genius%20Prince%27s%20Guide%20to%20Raising%20a%20Nation%20Out%20of%20Debt.mkv%20%5BWeekly%5D&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">217.2 MiB</td>
				<td class="text-center" data-timestamp="1645544079">2022-02-22 15:34</td>

				<td class="text-center">48</td>
				<td class="text-center">9</td>
				<td class="text-center">61</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_2" title="Anime - English-translated">
						<img src="/static/img/icons/nyaa/1_2.png" alt="Anime - English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494326" title="Kenja no Deshi wo Nanoru Kenja - 07 - 1080p WEB H.264 -NanDesuKa (B-Global).mkv">Kenja no Deshi wo Nanoru Kenja - 07 - 1080p WEB H.264 -NanDesuKa (B-Global).mkv</a>
				</td>
				<td class="text-center">
					<a href="/download/1494326.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:5100462e63da8159219655b37e9853d6a09d6d4b&amp;dn=Kenja%20no%20Deshi%20wo%20Nanoru%20Kenja%20-%2007%20-%201080p%20WEB%20H.264%20-NanDesuKa%20%28B-Global%29.mkv&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">1.0 GiB</td>
				<td class="text-center" data-timestamp="1645544055">2022-02-22 15:34</td>

				<td class="text-center">17</td>
				<td class="text-center">3</td>
				<td class="text-center">28</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_2" title="Anime - English-translated">
						<img src="/static/img/icons/nyaa/1_2.png" alt="Anime - English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494325" title="Kenja no Deshi wo Nanoru Kenja - 07 - 1080p WEB HEVC -NanDesuKa (B-Global).mkv">Kenja no Deshi wo Nanoru Kenja - 07 - 1080p WEB HEVC -NanDesuKa (B-Global).mkv</a>
				</td>
				<td class="text-center">
					<a href="/download/1494325.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:2b20de5e08e9b187066b7e17f2d98846bd18e79c&amp;dn=Kenja%20no%20Deshi%20wo%20Nanoru%20Kenja%20-%2007%20-%201080p%20WEB%20HEVC%20-NanDesuKa%20%28B-Global%29.mkv&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">636.3 MiB</td>
				<td class="text-center" data-timestamp="1645544052">2022-02-22 15:34</td>

				<td class="text-center">29</td>
				<td class="text-center">1</td>
				<td class="text-center">43</td>
			</tr>
			<tr class="success">
				<td>
					<a href="/?c=1_2" title="Anime - English-translated">
						<img src="/static/img/icons/nyaa/1_2.png" alt="Anime - English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494324" title="[Erai-raws] Kenja no Deshi o Nanoru Kenja - 07 [1080p][Multiple Subtitle][5D7F92BF].mkv">[Erai-raws] Kenja no Deshi o Nanoru Kenja - 07 [1080p][Multiple Subtitle][5D7F92BF].mkv</a>
				</td>
				<td class="text-center">
					<a href="/download/1494324.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:c0bb6d1043c02d2f7093e444b3f20df5b7a4390e&amp;dn=%5BErai-raws%5D%20Kenja%20no%20Deshi%20o%20Nanoru%20Kenja%20-%2007%20%5B1080p%5D%5BMultiple%20Subtitle%5D%5B5D7F92BF%5D.mkv&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">1.0 GiB</td>
				<td class="text-center" data-timestamp="1645543963">2022-02-22 15:32</td>

				<td class="text-center">138</td>
				<td class="text-center">30</td>
				<td class="text-center">196</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_2" title="Anime - English-translated">
						<img src="/static/img/icons/nyaa/1_2.png" alt="Anime - English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494323" title="Kenja no Deshi wo Nanoru Kenja - S01E07 - ENG 1080p WEB H.264 -NanDesuKa (FUNi).mkv">Kenja no Deshi wo Nanoru Kenja - S01E07 - ENG 1080p WEB H.264 -NanDesuKa (FUNi).mkv</a>
				</td>
				<td class="text-center">
					<a href="/download/1494323.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:c70a00daa2ba44955236b53e077cdafbdc32b15e&amp;dn=Kenja%20no%20Deshi%20wo%20Nanoru%20Kenja%20-%20S01E07%20-%20ENG%201080p%20WEB%20H.264%20-NanDesuKa%20%28FUNi%29.mkv&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">1.0 GiB</td>
				<td class="text-center" data-timestamp="1645543959">2022-02-22 15:32</td>

				<td class="text-center">52</td>
				<td class="text-center">7</td>
				<td class="text-center">78</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_2" title="Anime - English-translated">
						<img src="/static/img/icons/nyaa/1_2.png" alt="Anime - English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494322" title="Kenja no Deshi wo Nanoru Kenja - S01E07 - 1080p WEB H.264 -NanDesuKa (FUNi).mkv">Kenja no Deshi wo Nanoru Kenja - S01E07 - 1080p WEB H.264 -NanDesuKa (FUNi).mkv</a>
				</td>
				<td class="text-center">
					<a href="/download/1494322.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:fb31ec874983e94775f733cbd07c21e4216bea59&amp;dn=Kenja%20no%20Deshi%20wo%20Nanoru%20Kenja%20-%20S01E07%20-%201080p%20WEB%20H.264%20-NanDesuKa%20%28FUNi%29.mkv&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">1.0 GiB</td>
				<td class="text-center" data-timestamp="1645543958">2022-02-22 15:32</td>

				<td class="text-center">9</td>
				<td class="text-center">1</td>
				<td class="text-center">13</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_3" title="Anime - Non-English-translated">
						<img src="/static/img/icons/nyaa/1_3.png" alt="Anime - Non-English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494321" title="Kenja no Deshi wo Nanoru Kenja - S01E07 - VOSTFR 1080p WEB x264 -NanDesuKa (WAKA).mkv">Kenja no Deshi wo Nanoru Kenja - S01E07 - VOSTFR 1080p WEB x264 -NanDesuKa (WAKA).mkv</a>
				</td>
				<td class="text-center">
					<a href="/download/1494321.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:b48a352081f35818cfbcd51aabd8870810b5fc7e&amp;dn=Kenja%20no%20Deshi%20wo%20Nanoru%20Kenja%20-%20S01E07%20-%20VOSTFR%201080p%20WEB%20x264%20-NanDesuKa%20%28WAKA%29.mkv&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">877.6 MiB</td>
				<td class="text-center" data-timestamp="1645543957">2022-02-22 15:32</td>

				<td class="text-center">60</td>
				<td class="text-center">15</td>
				<td class="text-center">59</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_3" title="Anime - Non-English-translated">
						<img src="/static/img/icons/nyaa/1_3.png" alt="Anime - Non-English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494320" title="Kenja no Deshi wo Nanoru Kenja - S01E07 - VOSTFR 720p WEB x264 -NanDesuKa (WAKA).mkv">Kenja no Deshi wo Nanoru Kenja - S01E07 - VOSTFR 720p WEB x264 -NanDesuKa (WAKA).mkv</a>
				</td>
				<td class="text-center">
					<a href="/download/1494320.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:d582f6ce0f96527095779da368ffddd48c18b43f&amp;dn=Kenja%20no%20Deshi%20wo%20Nanoru%20Kenja%20-%20S01E07%20-%20VOSTFR%20720p%20WEB%20x264%20-NanDesuKa%20%28WAKA%29.mkv&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">429.9 MiB</td>
				<td class="text-center" data-timestamp="1645543948">2022-02-22 15:32</td>

				<td class="text-center">17</td>
				<td class="text-center">6</td>
				<td class="text-center">21</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_3" title="Anime - Non-English-translated">
						<img src="/static/img/icons/nyaa/1_3.png" alt="Anime - Non-English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494319" title="[Lilith-Raws] 自稱賢者弟子的賢者 / Kenja no Deshi wo Nanoru Kenja - 07 [Baha][WEB-DL][1080p][AVC AAC][CHT][MP4]">[Lilith-Raws] 自稱賢者弟子的賢者 / Kenja no Deshi wo Nanoru Kenja - 07 [Baha][WEB-DL][1080p][AVC AAC][CHT][MP4]</a>
				</td>
				<td class="text-center">
					<a href="/download/1494319.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:a259381e36a315f5dbb228b306eb36bda6abe434&amp;dn=%5BLilith-Raws%5D%20%E8%87%AA%E7%A8%B1%E8%B3%A2%E8%80%85%E5%BC%9F%E5%AD%90%E7%9A%84%E8%B3%A2%E8%80%85%20%2F%20Kenja%20no%20Deshi%20wo%20Nanoru%20Kenja%20-%2007%20%5BBaha%5D%5BWEB-DL%5D%5B1080p%5D%5BAVC%20AAC%5D%5BCHT%5D%5BMP4%5D&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">385.7 MiB</td>
				<td class="text-center" data-timestamp="1645543944">2022-02-22 15:32</td>

				<td class="text-center">215</td>
				<td class="text-center">93</td>
				<td class="text-center">359</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_2" title="Anime - English-translated">
						<img src="/static/img/icons/nyaa/1_2.png" alt="Anime - English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494318" title="Kenja no Deshi wo Nanoru Kenja - S01E07 - 720p WEB H.264 -NanDesuKa (FUNi).mkv">Kenja no Deshi wo Nanoru Kenja - S01E07 - 720p WEB H.264 -NanDesuKa (FUNi).mkv</a>
				</td>
				<td class="text-center">
					<a href="/download/1494318.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:3229ec1e8bd229689a71d16a0d72032f94e52971&amp;dn=Kenja%20no%20Deshi%20wo%20Nanoru%20Kenja%20-%20S01E07%20-%20720p%20WEB%20H.264%20-NanDesuKa%20%28FUNi%29.mkv&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">629.7 MiB</td>
				<td class="text-center" data-timestamp="1645543939">2022-02-22 15:32</td>

				<td class="text-center">13</td>
				<td class="text-center">2</td>
				<td class="text-center">24</td>
			</tr>
			<tr class="success">
				<td>
					<a href="/?c=1_2" title="Anime - English-translated">
						<img src="/static/img/icons/nyaa/1_2.png" alt="Anime - English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494317" title="[Erai-raws] Kenja no Deshi o Nanoru Kenja - 07 [720p][Multiple Subtitle][6250A7C2].mkv">[Erai-raws] Kenja no Deshi o Nanoru Kenja - 07 [720p][Multiple Subtitle][6250A7C2].mkv</a>
				</td>
				<td class="text-center">
					<a href="/download/1494317.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:9eec148c0856bcb95fcfc1fc303e66191507d5de&amp;dn=%5BErai-raws%5D%20Kenja%20no%20Deshi%20o%20Nanoru%20Kenja%20-%2007%20%5B720p%5D%5BMultiple%20Subtitle%5D%5B6250A7C2%5D.mkv&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">630.7 MiB</td>
				<td class="text-center" data-timestamp="1645543936">2022-02-22 15:32</td>

				<td class="text-center">68</td>
				<td class="text-center">15</td>
				<td class="text-center">83</td>
			</tr>
			<tr class="success">
				<td>
					<a href="/?c=1_2" title="Anime - English-translated">
						<img src="/static/img/icons/nyaa/1_2.png" alt="Anime - English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494316#comments" class="comments" title="1 comment">
						<i class="fa fa-comments-o"></i>1</a>
					<a href="/view/1494316" title="[SubsPlease] Kenja no Deshi wo Nanoru Kenja - 07 (1080p) [0B83CD5E].mkv">[SubsPlease] Kenja no Deshi wo Nanoru Kenja - 07 (1080p) [0B83CD5E].mkv</a>
				</td>
				<td class="text-center">
					<a href="/download/1494316.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:608ccdb8e78f667e185bd10f3c0014802db45985&amp;dn=%5BSubsPlease%5D%20Kenja%20no%20Deshi%20wo%20Nanoru%20Kenja%20-%2007%20%281080p%29%20%5B0B83CD5E%5D.mkv&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">1.0 GiB</td>
				<td class="text-center" data-timestamp="1645543927">2022-02-22 15:32</td>

				<td class="text-center">974</td>
				<td class="text-center">161</td>
				<td class="text-center">1113</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_2" title="Anime - English-translated">
						<img src="/static/img/icons/nyaa/1_2.png" alt="Anime - English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494315" title="Kenja no Deshi wo Nanoru Kenja - S01E07 - ENG 720p WEB H.264 -NanDesuKa (FUNi).mkv">Kenja no Deshi wo Nanoru Kenja - S01E07 - ENG 720p WEB H.264 -NanDesuKa (FUNi).mkv</a>
				</td>
				<td class="text-center">
					<a href="/download/1494315.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:cab96bda236bc0148cb67ab17363e13af1814218&amp;dn=Kenja%20no%20Deshi%20wo%20Nanoru%20Kenja%20-%20S01E07%20-%20ENG%20720p%20WEB%20H.264%20-NanDesuKa%20%28FUNi%29.mkv&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">626.5 MiB</td>
				<td class="text-center" data-timestamp="1645543925">2022-02-22 15:32</td>

				<td class="text-center">16</td>
				<td class="text-center">3</td>
				<td class="text-center">17</td>
			</tr>
			<tr class="success">
				<td>
					<a href="/?c=1_2" title="Anime - English-translated">
						<img src="/static/img/icons/nyaa/1_2.png" alt="Anime - English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494314" title="[Erai-raws] Kenja no Deshi o Nanoru Kenja - 07 [540p][Multiple Subtitle][E85953DA].mkv">[Erai-raws] Kenja no Deshi o Nanoru Kenja - 07 [540p][Multiple Subtitle][E85953DA].mkv</a>
				</td>
				<td class="text-center">
					<a href="/download/1494314.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:cecbee27287905302ef2a4557286536ef9f35d41&amp;dn=%5BErai-raws%5D%20Kenja%20no%20Deshi%20o%20Nanoru%20Kenja%20-%2007%20%5B540p%5D%5BMultiple%20Subtitle%5D%5BE85953DA%5D.mkv&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">303.3 MiB</td>
				<td class="text-center" data-timestamp="1645543924">2022-02-22 15:32</td>

				<td class="text-center">18</td>
				<td class="text-center">5</td>
				<td class="text-center">35</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_3" title="Anime - Non-English-translated">
						<img src="/static/img/icons/nyaa/1_3.png" alt="Anime - Non-English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494313" title="Kenja no Deshi wo Nanoru Kenja - S01E07 - VOSTFR 540p WEB x264 -NanDesuKa (WAKA).mkv">Kenja no Deshi wo Nanoru Kenja - S01E07 - VOSTFR 540p WEB x264 -NanDesuKa (WAKA).mkv</a>
				</td>
				<td class="text-center">
					<a href="/download/1494313.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:e41a1e8ef53d337a4ab0d575f24dd9b2ac3be7e4&amp;dn=Kenja%20no%20Deshi%20wo%20Nanoru%20Kenja%20-%20S01E07%20-%20VOSTFR%20540p%20WEB%20x264%20-NanDesuKa%20%28WAKA%29.mkv&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">176.5 MiB</td>
				<td class="text-center" data-timestamp="1645543917">2022-02-22 15:31</td>

				<td class="text-center">4</td>
				<td class="text-center">3</td>
				<td class="text-center">5</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_2" title="Anime - English-translated">
						<img src="/static/img/icons/nyaa/1_2.png" alt="Anime - English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494312" title="[NC-Raws] 自称贤者弟子的贤者 / Kenja no Deshi wo Nanoru Kenja - 07 (B-Global 1920x1080 HEVC AAC MKV)">[NC-Raws] 自称贤者弟子的贤者 / Kenja no Deshi wo Nanoru Kenja - 07 (B-Global 1920x1080 HEVC AAC MKV)</a>
				</td>
				<td class="text-center">
					<a href="/download/1494312.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:9828c6a02b23e509e803f41b350560010c3b6731&amp;dn=%5BNC-Raws%5D%20%E8%87%AA%E7%A7%B0%E8%B4%A4%E8%80%85%E5%BC%9F%E5%AD%90%E7%9A%84%E8%B4%A4%E8%80%85%20%2F%20Kenja%20no%20Deshi%20wo%20Nanoru%20Kenja%20-%2007%20%28B-Global%201920x1080%20HEVC%20AAC%20MKV%29&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">636.2 MiB</td>
				<td class="text-center" data-timestamp="1645543915">2022-02-22 15:31</td>

				<td class="text-center">22</td>
				<td class="text-center">8</td>
				<td class="text-center">34</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_3" title="Anime - Non-English-translated">
						<img src="/static/img/icons/nyaa/1_3.png" alt="Anime - Non-English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494311" title="[NC-Raws] 自称贤者弟子的贤者（泰配版） / Kenja no Deshi (Thai version) - 07 (B-Global 1920x1080 HEVC AAC MKV)">[NC-Raws] 自称贤者弟子的贤者（泰配版） / Kenja no Deshi (Thai version) - 07 (B-Global 1920x1080 HEVC AAC MKV)</a>
				</td>
				<td class="text-center">
					<a href="/download/1494311.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:e7de1a47559d8441233715cedbf6ae14fd89fc05&amp;dn=%5BNC-Raws%5D%20%E8%87%AA%E7%A7%B0%E8%B4%A4%E8%80%85%E5%BC%9F%E5%AD%90%E7%9A%84%E8%B4%A4%E8%80%85%EF%BC%88%E6%B3%B0%E9%85%8D%E7%89%88%EF%BC%89%20%2F%20Kenja%20no%20Deshi%20%28Thai%20version%29%20-%2007%20%28B-Global%201920x1080%20HEVC%20AAC%20MKV%29&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">642.5 MiB</td>
				<td class="text-center" data-timestamp="1645543910">2022-02-22 15:31</td>

				<td class="text-center">7</td>
				<td class="text-center">4</td>
				<td class="text-center">9</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_3" title="Anime - Non-English-translated">
						<img src="/static/img/icons/nyaa/1_3.png" alt="Anime - Non-English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494310" title="[NC-Raws] 自稱賢者弟子的賢者 / Kenja no Deshi wo Nanoru Kenja - 07 (Baha 1920x1080 AVC AAC MP4)">[NC-Raws] 自稱賢者弟子的賢者 / Kenja no Deshi wo Nanoru Kenja - 07 (Baha 1920x1080 AVC AAC MP4)</a>
				</td>
				<td class="text-center">
					<a href="/download/1494310.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:cce7c753cbe376f8454fc324d63a530df01d85a1&amp;dn=%5BNC-Raws%5D%20%E8%87%AA%E7%A8%B1%E8%B3%A2%E8%80%85%E5%BC%9F%E5%AD%90%E7%9A%84%E8%B3%A2%E8%80%85%20%2F%20Kenja%20no%20Deshi%20wo%20Nanoru%20Kenja%20-%2007%20%28Baha%201920x1080%20AVC%20AAC%20MP4%29&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">385.8 MiB</td>
				<td class="text-center" data-timestamp="1645543910">2022-02-22 15:31</td>

				<td class="text-center">94</td>
				<td class="text-center">37</td>
				<td class="text-center">145</td>
			</tr>
			<tr class="success">
				<td>
					<a href="/?c=1_2" title="Anime - English-translated">
						<img src="/static/img/icons/nyaa/1_2.png" alt="Anime - English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494309" title="[SubsPlease] Kenja no Deshi wo Nanoru Kenja - 07 (720p) [F6D9812E].mkv">[SubsPlease] Kenja no Deshi wo Nanoru Kenja - 07 (720p) [F6D9812E].mkv</a>
				</td>
				<td class="text-center">
					<a href="/download/1494309.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:277e216c5c93d1e53c08d8fb4c907745312db77c&amp;dn=%5BSubsPlease%5D%20Kenja%20no%20Deshi%20wo%20Nanoru%20Kenja%20-%2007%20%28720p%29%20%5BF6D9812E%5D.mkv&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">630.0 MiB</td>
				<td class="text-center" data-timestamp="1645543907">2022-02-22 15:31</td>

				<td class="text-center">633</td>
				<td class="text-center">161</td>
				<td class="text-center">770</td>
			</tr>
			<tr class="success">
				<td>
					<a href="/?c=1_2" title="Anime - English-translated">
						<img src="/static/img/icons/nyaa/1_2.png" alt="Anime - English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494308" title="[SubsPlease] Kenja no Deshi wo Nanoru Kenja - 07 (540p) [031C0A4B].mkv">[SubsPlease] Kenja no Deshi wo Nanoru Kenja - 07 (540p) [031C0A4B].mkv</a>
				</td>
				<td class="text-center">
					<a href="/download/1494308.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:ee82fae41cc8db0b3015ed5e8f99c39a84586f74&amp;dn=%5BSubsPlease%5D%20Kenja%20no%20Deshi%20wo%20Nanoru%20Kenja%20-%2007%20%28540p%29%20%5B031C0A4B%5D.mkv&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">302.7 MiB</td>
				<td class="text-center" data-timestamp="1645543893">2022-02-22 15:31</td>

				<td class="text-center">210</td>
				<td class="text-center">64</td>
				<td class="text-center">268</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_2" title="Anime - English-translated">
						<img src="/static/img/icons/nyaa/1_2.png" alt="Anime - English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494307" title="Kenja no Deshi wo Nanoru Kenja - S01E07 - 540p WEB H.264 -NanDesuKa (FUNi).mkv">Kenja no Deshi wo Nanoru Kenja - S01E07 - 540p WEB H.264 -NanDesuKa (FUNi).mkv</a>
				</td>
				<td class="text-center">
					<a href="/download/1494307.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:635a851af6f4e25a9ca3c6fc8df1dda9a280d5fe&amp;dn=Kenja%20no%20Deshi%20wo%20Nanoru%20Kenja%20-%20S01E07%20-%20540p%20WEB%20H.264%20-NanDesuKa%20%28FUNi%29.mkv&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">346.6 MiB</td>
				<td class="text-center" data-timestamp="1645543892">2022-02-22 15:31</td>

				<td class="text-center">0</td>
				<td class="text-center">9</td>
				<td class="text-center">0</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_2" title="Anime - English-translated">
						<img src="/static/img/icons/nyaa/1_2.png" alt="Anime - English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494306" title="Kenja no Deshi wo Nanoru Kenja - S01E07 - ENG 540p WEB H.264 -NanDesuKa (FUNi).mkv">Kenja no Deshi wo Nanoru Kenja - S01E07 - ENG 540p WEB H.264 -NanDesuKa (FUNi).mkv</a>
				</td>
				<td class="text-center">
					<a href="/download/1494306.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:9285c9cb64f5eeddd956dd61207e10cad982fb9c&amp;dn=Kenja%20no%20Deshi%20wo%20Nanoru%20Kenja%20-%20S01E07%20-%20ENG%20540p%20WEB%20H.264%20-NanDesuKa%20%28FUNi%29.mkv&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">346.1 MiB</td>
				<td class="text-center" data-timestamp="1645543888">2022-02-22 15:31</td>

				<td class="text-center">7</td>
				<td class="text-center">2</td>
				<td class="text-center">14</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_3" title="Anime - Non-English-translated">
						<img src="/static/img/icons/nyaa/1_3.png" alt="Anime - Non-English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494305" title="[Lilith-Raws] 天才王子的赤字國家重生術 / Tensai Ouji no Akaji Kokka Saisei Jutsu - 07 [Baha][WEB-DL][1080p][AVC AAC][CHT][MP4]">[Lilith-Raws] 天才王子的赤字國家重生術 / Tensai Ouji no Akaji Kokka Saisei Jutsu - 07 [Baha][WEB-DL][1080p][AVC AAC][CHT][MP4]</a>
				</td>
				<td class="text-center">
					<a href="/download/1494305.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:ece5ed0439fe171cefbe247bd48b71314937a897&amp;dn=%5BLilith-Raws%5D%20%E5%A4%A9%E6%89%8D%E7%8E%8B%E5%AD%90%E7%9A%84%E8%B5%A4%E5%AD%97%E5%9C%8B%E5%AE%B6%E9%87%8D%E7%94%9F%E8%A1%93%20%2F%20Tensai%20Ouji%20no%20Akaji%20Kokka%20Saisei%20Jutsu%20-%2007%20%5BBaha%5D%5BWEB-DL%5D%5B1080p%5D%5BAVC%20AAC%5D%5BCHT%5D%5BMP4%5D&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">354.4 MiB</td>
				<td class="text-center" data-timestamp="1645543678">2022-02-22 15:27</td>

				<td class="text-center">246</td>
				<td class="text-center">104</td>
				<td class="text-center">386</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_3" title="Anime - Non-English-translated">
						<img src="/static/img/icons/nyaa/1_3.png" alt="Anime - Non-English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494304" title="[Lilith-Raws] 與變成了異世界美少女的大叔一起冒險 / Fantasy Bishoujo Juniku Ojisan to - 07 [Baha][WEB-DL][1080p][AVC AAC][CHT][MP4]">[Lilith-Raws] 與變成了異世界美少女的大叔一起冒險 / Fantasy Bishoujo Juniku Ojisan to - 07 [Baha][WEB-DL][1080p][AVC AAC][CHT][MP4]</a>
				</td>
				<td class="text-center">
					<a href="/download/1494304.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:c2d5d6f6679c169d4d946ce24fbeeaa2b9e1584f&amp;dn=%5BLilith-Raws%5D%20%E8%88%87%E8%AE%8A%E6%88%90%E4%BA%86%E7%95%B0%E4%B8%96%E7%95%8C%E7%BE%8E%E5%B0%91%E5%A5%B3%E7%9A%84%E5%A4%A7%E5%8F%94%E4%B8%80%E8%B5%B7%E5%86%92%E9%9A%AA%20%2F%20Fantasy%20Bishoujo%20Juniku%20Ojisan%20to%20-%2007%20%5BBaha%5D%5BWEB-DL%5D%5B1080p%5D%5BAVC%20AAC%5D%5BCHT%5D%5BMP4%5D&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">513.5 MiB</td>
				<td class="text-center" data-timestamp="1645543646">2022-02-22 15:27</td>

				<td class="text-center">204</td>
				<td class="text-center">86</td>
				<td class="text-center">288</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_2" title="Anime - English-translated">
						<img src="/static/img/icons/nyaa/1_2.png" alt="Anime - English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494303" title="[Henshin-Man] Hyouga Senshi Guyslugger - Episode 01 (WEB-DL)(480p)(MKV)">[Henshin-Man] Hyouga Senshi Guyslugger - Episode 01 (WEB-DL)(480p)(MKV)</a>
				</td>
				<td class="text-center">
					<a href="/download/1494303.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:8e7359d3baaf8f1dd1976b863ff9e37b220985e6&amp;dn=%5BHenshin-Man%5D%20Hyouga%20Senshi%20Guyslugger%20-%20Episode%2001%20%28WEB-DL%29%28480p%29%28MKV%29&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">370.0 MiB</td>
				<td class="text-center" data-timestamp="1645543230">2022-02-22 15:20</td>

				<td class="text-center">6</td>
				<td class="text-center">4</td>
				<td class="text-center">6</td>
			</tr>
			<tr class="danger">
				<td>
					<a href="/?c=1_2" title="Anime - English-translated">
						<img src="/static/img/icons/nyaa/1_2.png" alt="Anime - English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494302" title="[ASW] Tensai Ouji no Akaji Kokka Saisei Jutsu - 07 [1080p HEVC x265 10Bit][AAC]">[ASW] Tensai Ouji no Akaji Kokka Saisei Jutsu - 07 [1080p HEVC x265 10Bit][AAC]</a>
				</td>
				<td class="text-center">
					<a href="/download/1494302.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:704f8250b3fca833b75e84378622a9a79070c857&amp;dn=%5BASW%5D%20Tensai%20Ouji%20no%20Akaji%20Kokka%20Saisei%20Jutsu%20-%2007%20%5B1080p%20HEVC%20x265%2010Bit%5D%5BAAC%5D&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">221.4 MiB</td>
				<td class="text-center" data-timestamp="1645542963">2022-02-22 15:16</td>

				<td class="text-center">271</td>
				<td class="text-center">38</td>
				<td class="text-center">397</td>
			</tr>
			<tr class="danger">
				<td>
					<a href="/?c=1_2" title="Anime - English-translated">
						<img src="/static/img/icons/nyaa/1_2.png" alt="Anime - English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494301" title="[YuiSubs] Tensai Ouji no Akaji Kokka Saisei Jutsu - 07  (x265 H.265 1080p)">[YuiSubs] Tensai Ouji no Akaji Kokka Saisei Jutsu - 07  (x265 H.265 1080p)</a>
				</td>
				<td class="text-center">
					<a href="/download/1494301.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:bf1472014ce803f5503bcdd0ecd8b2e56b1aa1e2&amp;dn=%5BYuiSubs%5D%20Tensai%20Ouji%20no%20Akaji%20Kokka%20Saisei%20Jutsu%20-%2007%20%20%28x265%20H.265%201080p%29&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">296.3 MiB</td>
				<td class="text-center" data-timestamp="1645542912">2022-02-22 15:15</td>

				<td class="text-center">35</td>
				<td class="text-center">6</td>
				<td class="text-center">38</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=3_3" title="Literature - Raw">
						<img src="/static/img/icons/nyaa/3_3.png" alt="Literature - Raw" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494300" title="[白正男×山戸大輔] テコンダー朴 第01-06巻">[白正男×山戸大輔] テコンダー朴 第01-06巻</a>
				</td>
				<td class="text-center">
					<a href="/download/1494300.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:4f07746f2f12df65bb168cccbb92fc54cfe6523c&amp;dn=%5B%E7%99%BD%E6%AD%A3%E7%94%B7%C3%97%E5%B1%B1%E6%88%B8%E5%A4%A7%E8%BC%94%5D%20%E3%83%86%E3%82%B3%E3%83%B3%E3%83%80%E3%83%BC%E6%9C%B4%20%E7%AC%AC01-06%E5%B7%BB&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">1.0 GiB</td>
				<td class="text-center" data-timestamp="1645542373">2022-02-22 15:06</td>

				<td class="text-center">74</td>
				<td class="text-center">43</td>
				<td class="text-center">139</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_4" title="Anime - Raw">
						<img src="/static/img/icons/nyaa/1_4.png" alt="Anime - Raw" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494299" title="[Ohys-Raws] Waccha PriMagi! - 19 (AT-X 1280x720 x264 AAC).mp4">[Ohys-Raws] Waccha PriMagi! - 19 (AT-X 1280x720 x264 AAC).mp4</a>
				</td>
				<td class="text-center">
					<a href="/download/1494299.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:944932a25092e1f340f322769c8f145c4341de0e&amp;dn=%5BOhys-Raws%5D%20Waccha%20PriMagi%21%20-%2019%20%28AT-X%201280x720%20x264%20AAC%29.mp4&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">491.2 MiB</td>
				<td class="text-center" data-timestamp="1645542237">2022-02-22 15:03</td>

				<td class="text-center">96</td>
				<td class="text-center">47</td>
				<td class="text-center">151</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_2" title="Anime - English-translated">
						<img src="/static/img/icons/nyaa/1_2.png" alt="Anime - English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494298" title="[NC-Raws] 与变成了异世界美少女的好友一起冒险 / Fantasy Bishoujo - 07 (B-Global 1920x1080 HEVC AAC MKV)">[NC-Raws] 与变成了异世界美少女的好友一起冒险 / Fantasy Bishoujo - 07 (B-Global 1920x1080 HEVC AAC MKV)</a>
				</td>
				<td class="text-center">
					<a href="/download/1494298.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:449b41f77e1c9189d2f8a879e2d5f345a4a09927&amp;dn=%5BNC-Raws%5D%20%E4%B8%8E%E5%8F%98%E6%88%90%E4%BA%86%E5%BC%82%E4%B8%96%E7%95%8C%E7%BE%8E%E5%B0%91%E5%A5%B3%E7%9A%84%E5%A5%BD%E5%8F%8B%E4%B8%80%E8%B5%B7%E5%86%92%E9%99%A9%20%2F%20Fantasy%20Bishoujo%20-%2007%20%28B-Global%201920x1080%20HEVC%20AAC%20MKV%29&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">651.7 MiB</td>
				<td class="text-center" data-timestamp="1645542097">2022-02-22 15:01</td>

				<td class="text-center">26</td>
				<td class="text-center">7</td>
				<td class="text-center">38</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_3" title="Anime - Non-English-translated">
						<img src="/static/img/icons/nyaa/1_3.png" alt="Anime - Non-English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494297" title="[NC-Raws] 與變成了異世界美少女的大叔一起冒險 / Fantasy Bishoujo - 07 (Baha 1920x1080 AVC AAC MP4)">[NC-Raws] 與變成了異世界美少女的大叔一起冒險 / Fantasy Bishoujo - 07 (Baha 1920x1080 AVC AAC MP4)</a>
				</td>
				<td class="text-center">
					<a href="/download/1494297.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:5f9561f8f3ab63ad7333b5c99c1603e08b1d39e1&amp;dn=%5BNC-Raws%5D%20%E8%88%87%E8%AE%8A%E6%88%90%E4%BA%86%E7%95%B0%E4%B8%96%E7%95%8C%E7%BE%8E%E5%B0%91%E5%A5%B3%E7%9A%84%E5%A4%A7%E5%8F%94%E4%B8%80%E8%B5%B7%E5%86%92%E9%9A%AA%20%2F%20Fantasy%20Bishoujo%20-%2007%20%28Baha%201920x1080%20AVC%20AAC%20MP4%29&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">513.7 MiB</td>
				<td class="text-center" data-timestamp="1645542093">2022-02-22 15:01</td>

				<td class="text-center">128</td>
				<td class="text-center">39</td>
				<td class="text-center">230</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_2" title="Anime - English-translated">
						<img src="/static/img/icons/nyaa/1_2.png" alt="Anime - English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494296" title="Fantasy Bishoujo Juniku Ojisan to - 07 - 1080p WEB H.264 -NanDesuKa (B-Global).mkv">Fantasy Bishoujo Juniku Ojisan to - 07 - 1080p WEB H.264 -NanDesuKa (B-Global).mkv</a>
				</td>
				<td class="text-center">
					<a href="/download/1494296.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:36c7fe027ceb28bb496dc07571d67fd5633587a8&amp;dn=Fantasy%20Bishoujo%20Juniku%20Ojisan%20to%20-%2007%20-%201080p%20WEB%20H.264%20-NanDesuKa%20%28B-Global%29.mkv&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">1.0 GiB</td>
				<td class="text-center" data-timestamp="1645542079">2022-02-22 15:01</td>

				<td class="text-center">120</td>
				<td class="text-center">16</td>
				<td class="text-center">160</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_2" title="Anime - English-translated">
						<img src="/static/img/icons/nyaa/1_2.png" alt="Anime - English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494295" title="Fantasy Bishoujo Juniku Ojisan to - 07 - 1080p WEB HEVC -NanDesuKa (B-Global).mkv">Fantasy Bishoujo Juniku Ojisan to - 07 - 1080p WEB HEVC -NanDesuKa (B-Global).mkv</a>
				</td>
				<td class="text-center">
					<a href="/download/1494295.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:3646d4fc6466478f9bd0266a9b4eed88e8f13239&amp;dn=Fantasy%20Bishoujo%20Juniku%20Ojisan%20to%20-%2007%20-%201080p%20WEB%20HEVC%20-NanDesuKa%20%28B-Global%29.mkv&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">651.9 MiB</td>
				<td class="text-center" data-timestamp="1645542076">2022-02-22 15:01</td>

				<td class="text-center">89</td>
				<td class="text-center">10</td>
				<td class="text-center">160</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_3" title="Anime - Non-English-translated">
						<img src="/static/img/icons/nyaa/1_3.png" alt="Anime - Non-English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494294" title="Touch! 10(1080p)-pt-br">Touch! 10(1080p)-pt-br</a>
				</td>
				<td class="text-center">
					<a href="/download/1494294.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:d6a6f3a52c5fbd5a3e710896f9259ff90dfdd902&amp;dn=Touch%21%2010%281080p%29-pt-br&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">1.3 GiB</td>
				<td class="text-center" data-timestamp="1645541627">2022-02-22 14:53</td>

				<td class="text-center">1</td>
				<td class="text-center">10</td>
				<td class="text-center">1</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_3" title="Anime - Non-English-translated">
						<img src="/static/img/icons/nyaa/1_3.png" alt="Anime - Non-English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494293" title="Touch! 09(1080p)-pt-br">Touch! 09(1080p)-pt-br</a>
				</td>
				<td class="text-center">
					<a href="/download/1494293.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:bfdb481c3758cb05b570348d2889aa5b0eac90e8&amp;dn=Touch%21%2009%281080p%29-pt-br&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">1.3 GiB</td>
				<td class="text-center" data-timestamp="1645541580">2022-02-22 14:53</td>

				<td class="text-center">1</td>
				<td class="text-center">8</td>
				<td class="text-center">0</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_3" title="Anime - Non-English-translated">
						<img src="/static/img/icons/nyaa/1_3.png" alt="Anime - Non-English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494292" title="Touch! 08(1080p)-pt-br">Touch! 08(1080p)-pt-br</a>
				</td>
				<td class="text-center">
					<a href="/download/1494292.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:6ffe9a47ad678bab8ca4de6e19ad90124520713a&amp;dn=Touch%21%2008%281080p%29-pt-br&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">1.2 GiB</td>
				<td class="text-center" data-timestamp="1645541416">2022-02-22 14:50</td>

				<td class="text-center">3</td>
				<td class="text-center">7</td>
				<td class="text-center">2</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_3" title="Anime - Non-English-translated">
						<img src="/static/img/icons/nyaa/1_3.png" alt="Anime - Non-English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494291" title="Touch! 07(1080p)-pt-br">Touch! 07(1080p)-pt-br</a>
				</td>
				<td class="text-center">
					<a href="/download/1494291.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:622e5b348e443108a14a9bdff2344767c4cb537f&amp;dn=Touch%21%2007%281080p%29-pt-br&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">1.1 GiB</td>
				<td class="text-center" data-timestamp="1645541380">2022-02-22 14:49</td>

				<td class="text-center">2</td>
				<td class="text-center">8</td>
				<td class="text-center">1</td>
			</tr>
			<tr class="success">
				<td>
					<a href="/?c=1_3" title="Anime - Non-English-translated">
						<img src="/static/img/icons/nyaa/1_3.png" alt="Anime - Non-English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494290" title="[PuyaSubs!] Tensai Ouji no Akaji Kokka Saisei Jutsu - 07 [ESP-ENG][720p][BF718FA8].mkv">[PuyaSubs!] Tensai Ouji no Akaji Kokka Saisei Jutsu - 07 [ESP-ENG][720p][BF718FA8].mkv</a>
				</td>
				<td class="text-center">
					<a href="/download/1494290.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:50c9de434aecd233e4f78c9c775c58107da8609c&amp;dn=%5BPuyaSubs%21%5D%20Tensai%20Ouji%20no%20Akaji%20Kokka%20Saisei%20Jutsu%20-%2007%20%5BESP-ENG%5D%5B720p%5D%5BBF718FA8%5D.mkv&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">623.8 MiB</td>
				<td class="text-center" data-timestamp="1645541369">2022-02-22 14:49</td>

				<td class="text-center">46</td>
				<td class="text-center">8</td>
				<td class="text-center">31</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_3" title="Anime - Non-English-translated">
						<img src="/static/img/icons/nyaa/1_3.png" alt="Anime - Non-English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494289" title="Touch! 06(1080p)-pt-br">Touch! 06(1080p)-pt-br</a>
				</td>
				<td class="text-center">
					<a href="/download/1494289.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:2369ad1893d99803e88bf7a759161aa929e93f00&amp;dn=Touch%21%2006%281080p%29-pt-br&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">1.3 GiB</td>
				<td class="text-center" data-timestamp="1645541341">2022-02-22 14:49</td>

				<td class="text-center">1</td>
				<td class="text-center">8</td>
				<td class="text-center">1</td>
			</tr>
			<tr class="success">
				<td>
					<a href="/?c=1_3" title="Anime - Non-English-translated">
						<img src="/static/img/icons/nyaa/1_3.png" alt="Anime - Non-English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494288" title="[PuyaSubs!] Tensai Ouji no Akaji Kokka Saisei Jutsu - 07 [ESP-ENG][1080p][9596C2B0].mkv">[PuyaSubs!] Tensai Ouji no Akaji Kokka Saisei Jutsu - 07 [ESP-ENG][1080p][9596C2B0].mkv</a>
				</td>
				<td class="text-center">
					<a href="/download/1494288.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:d4c02d1d992a2022e1daae8816a329c9517f1ba0&amp;dn=%5BPuyaSubs%21%5D%20Tensai%20Ouji%20no%20Akaji%20Kokka%20Saisei%20Jutsu%20-%2007%20%5BESP-ENG%5D%5B1080p%5D%5B9596C2B0%5D.mkv&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">1.0 GiB</td>
				<td class="text-center" data-timestamp="1645541312">2022-02-22 14:48</td>

				<td class="text-center">55</td>
				<td class="text-center">18</td>
				<td class="text-center">47</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_3" title="Anime - Non-English-translated">
						<img src="/static/img/icons/nyaa/1_3.png" alt="Anime - Non-English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494287" title="[DJB] The Genius Prince&#39;s Guide to Raising a Nation Out of Debt - 07 VOSTFR [1080p].mp4">[DJB] The Genius Prince&#39;s Guide to Raising a Nation Out of Debt - 07 VOSTFR [1080p].mp4</a>
				</td>
				<td class="text-center">
					<a href="/download/1494287.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:49d3c02e2922f1ce08a4c9cdef3417fa5cfb14b9&amp;dn=%5BDJB%5D%20The%20Genius%20Prince%27s%20Guide%20to%20Raising%20a%20Nation%20Out%20of%20Debt%20-%2007%20VOSTFR%20%5B1080p%5D.mp4&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">337.6 MiB</td>
				<td class="text-center" data-timestamp="1645541131">2022-02-22 14:45</td>

				<td class="text-center">17</td>
				<td class="text-center">8</td>
				<td class="text-center">27</td>
			</tr>
			<tr class="danger">
				<td>
					<a href="/?c=1_2" title="Anime - English-translated">
						<img src="/static/img/icons/nyaa/1_2.png" alt="Anime - English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494286#comments" class="comments" title="3 comments">
						<i class="fa fa-comments-o"></i>3</a>
					<a href="/view/1494286" title="Demon.Slayer.Kimetsu.no.Yaiba.The.Movie.Mugen/Infinity.Train/Ressha-Hen.2020.1080p">Demon.Slayer.Kimetsu.no.Yaiba.The.Movie.Mugen/Infinity.Train/Ressha-Hen.2020.1080p</a>
				</td>
				<td class="text-center">
					<a href="/download/1494286.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:b351058be09a804c0a43234d571eaf22aa269d60&amp;dn=Demon.Slayer.Kimetsu.no.Yaiba.The.Movie.Mugen%2FInfinity.Train%2FRessha-Hen.2020.1080p&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">742.5 MiB</td>
				<td class="text-center" data-timestamp="1645540864">2022-02-22 14:41</td>

				<td class="text-center">0</td>
				<td class="text-center">10</td>
				<td class="text-center">1</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_2" title="Anime - English-translated">
						<img src="/static/img/icons/nyaa/1_2.png" alt="Anime - English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1494285" title="Tensai Ouji no Akaji Kokka Saisei Jutsu - 07 - 1080p WEB H.264 -NanDesuKa (B-Global).mkv">Tensai Ouji no Akaji Kokka Saisei Jutsu - 07 - 1080p WEB H.264 -NanDesuKa (B-Global).mkv</a>
				</td>
				<td class="text-center">
					<a href="/download/1494285.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:61e15c98875a361e6e48a18f09809c9f94181ba5&amp;dn=Tensai%20Ouji%20no%20Akaji%20Kokka%20Saisei%20Jutsu%20-%2007%20-%201080p%20WEB%20H.264%20-NanDesuKa%20%28B-Global%29.mkv&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">1.0 GiB</td>
				<td class="text-center" data-timestamp="1645540523">2022-02-22 14:35</td>

				<td class="text-center">35</td>
				<td class="text-center">4</td>
				<td class="text-center">57</td>
			</tr>
		</tbody>
	</table>
</div>

<div class="center">
	<nav>
  <ul class="pagination">
    <li class="disabled"><a href="#">&laquo;</a></li>
        <li class="active"><a href="#">1 <span class="sr-only">(current)</span></a></li>
        <li><a href="/?p=2">2</a></li>
        <li><a href="/?p=3">3</a></li>
        <li><a href="/?p=4">4</a></li>
        <li><a href="/?p=5">5</a></li>
        <li><a href="/?p=6">6</a></li>

    <li><a rel="next" href="/?p=2">&raquo;</a></li>
</ul>
</nav>

</div>
		</div> <!-- /container -->

		<footer style="text-align: center;">
			<p>Dark Mode: <a href="#" id="themeToggle">Toggle</a></p>
		</footer>
	</body>
</html>`
