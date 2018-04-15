var xmlhttp;
function main(arg)
{
    var arg = JSON.stringify({'method':'main'})
    CreateXMLHttpRequest();
    xmlhttp.onreadystatechange = callhandle;
    xmlhttp.open("POST","/",true);
    xmlhttp.setRequestHeader("Content-Type","application/x-www-form-urlencoded");  //用POST的时候一定要有这句
    xmlhttp.send(arg);
    
}
function show_anime(_id,prepage)
{
    var arg = JSON.stringify({'method':'show_anime','params':[_id,prepage]})
    CreateXMLHttpRequest();
    xmlhttp.onreadystatechange = callhandle;
    xmlhttp.open("POST","/",true);
    xmlhttp.setRequestHeader("Content-Type","application/x-www-form-urlencoded");  //用POST的时候一定要有这句
    xmlhttp.send(arg);
    
}

function show_collection()
{
    var arg = JSON.stringify({'method':'show_collection','params':[]})
    CreateXMLHttpRequest();
    xmlhttp.onreadystatechange = callhandle;
    xmlhttp.open("POST","/",true);
    xmlhttp.setRequestHeader("Content-Type","application/x-www-form-urlencoded");  //用POST的时候一定要有这句
    xmlhttp.send(arg);
    
}

function search_collection()
{
    var key = document.getElementById("searchInput").value
    var arg = JSON.stringify({'method':'show_collection','params':[key]})
    CreateXMLHttpRequest();
    xmlhttp.onreadystatechange = callhandle;
    xmlhttp.open("POST","/",true);
    xmlhttp.setRequestHeader("Content-Type","application/x-www-form-urlencoded");  //用POST的时候一定要有这句
    xmlhttp.send(arg);
    
}

function edit_anime(_id, prepage)
{
    var arg = JSON.stringify({'method':'edit_anime','params':[_id,prepage]})
    CreateXMLHttpRequest();
    xmlhttp.onreadystatechange = callhandle;
    xmlhttp.open("POST","/",true);
    xmlhttp.setRequestHeader("Content-Type","application/x-www-form-urlencoded");  //用POST的时候一定要有这句
    xmlhttp.send(arg);
}

function del_anime(_id)
{
    con = confirm("你确认你要删除这部动画吗?你真的要删除吗?你确定?你不会后悔?");
    if(con == false)
    {
        return;
    }
    var arg = JSON.stringify({'method':'del_anime','params':[_id]})
    CreateXMLHttpRequest();
    xmlhttp.onreadystatechange = callhandle;
    xmlhttp.open("POST","/",true);
    xmlhttp.setRequestHeader("Content-Type","application/x-www-form-urlencoded");  //用POST的时候一定要有这句
    xmlhttp.send(arg);
}

function update_anime(_id, prepage)
{
    var elems = document.getElementsByTagName("input");
    var arg = "{\"method\":\"update_anime\",\"params\":[\"" + _id + "\",\"" + prepage + "\",";
    for (var i = 0; i < elems.length; i++) 
    {
        if(elems[i].type == 'text')
        {
            arg += "\"" + elems[i].value + "\"";
            arg += ','
        }
    }
    arg =arg.replace(/,$/,"")
    arg += "]}";
    CreateXMLHttpRequest();
    xmlhttp.onreadystatechange = callhandle;
    xmlhttp.open("POST","/",true);
    xmlhttp.setRequestHeader("Content-Type","application/x-www-form-urlencoded");  //用POST的时候一定要有这句
    xmlhttp.send(arg);
}

function get_task()
{
    var arg = JSON.stringify({'method':'get_task','params':[]})
    CreateXMLHttpRequest();
    xmlhttp.onreadystatechange = callhandle;
    xmlhttp.open("POST","/",true);
    xmlhttp.setRequestHeader("Content-Type","application/x-www-form-urlencoded");  //用POST的时候一定要有这句
    xmlhttp.send(arg);
    // reload();
    get_task_by_ws();
}
function add_task()
{
    var magnet = document.getElementById("magnet").value;
    var savePath = document.getElementById("savePath").value;
    if(magnet == "" || savePath == "")
    {
        alert("链接或路径为空!!!");
        return;
    }
    var arg = JSON.stringify({'method':'add_task','params':[magnet, savePath]})
    CreateXMLHttpRequest();
    // xmlhttp.onreadystatechange = callhandle;
    xmlhttp.open("POST","/",true);
    xmlhttp.setRequestHeader("Content-Type","application/x-www-form-urlencoded");  //用POST的时候一定要有这句
    xmlhttp.send(arg);
}

function start_task(id)
{
    var arg = JSON.stringify({'method':'start_task','params':[id]})
    CreateXMLHttpRequest();
    // xmlhttp.onreadystatechange = callhandle;
    xmlhttp.open("POST","/",true);
    xmlhttp.setRequestHeader("Content-Type","application/x-www-form-urlencoded");  //用POST的时候一定要有这句
    xmlhttp.send(arg);
}

function pause_task(id)
{
    var arg = JSON.stringify({'method':'pause_task','params':[id]})
    CreateXMLHttpRequest();
    // xmlhttp.onreadystatechange = callhandle;
    xmlhttp.open("POST","/",true);
    xmlhttp.setRequestHeader("Content-Type","application/x-www-form-urlencoded");  //用POST的时候一定要有这句
    xmlhttp.send(arg);
}

function del_task(id)
{
    var arg = JSON.stringify({'method':'del_task','params':[id]})
    CreateXMLHttpRequest();
    // xmlhttp.onreadystatechange = callhandle;
    xmlhttp.open("POST","/",true);
    xmlhttp.setRequestHeader("Content-Type","application/x-www-form-urlencoded");  //用POST的时候一定要有这句
    xmlhttp.send(arg);
}

function show_adtask()
{
    var arg = JSON.stringify({'method':'show_adtask','params':[]})
    CreateXMLHttpRequest();
    xmlhttp.onreadystatechange = callhandle;
    xmlhttp.open("POST","/",true);
    xmlhttp.setRequestHeader("Content-Type","application/x-www-form-urlencoded");  //用POST的时候一定要有这句
    xmlhttp.send(arg);
}

function edit_adtask(id)
{
    var arg = JSON.stringify({'method':'edit_adtask','params':[id]})
    CreateXMLHttpRequest();
    xmlhttp.onreadystatechange = callhandle;
    xmlhttp.open("POST","/",true);
    xmlhttp.setRequestHeader("Content-Type","application/x-www-form-urlencoded");  //用POST的时候一定要有这句
    xmlhttp.send(arg);
}
function update_adTask(_id)
{
    var elems = document.getElementsByTagName("input");
    var arg = '{"method":"update_adTask","params":["' + _id + '",';
    
    for (var i = 0; i < elems.length; i++) 
    {
        if(elems[i].type == 'text' || elems[i].type == 'checkbox')
        {
            value = elems[i].value;
            value = value.replace(/\\/g,"\\\\");
            value = value.replace(/\"/g,"\\\"");
            if(elems[i].type == 'checkbox')
            {
                value = elems[i].checked;
            }
            arg += '"' + value + '"';
            arg += ','
        }
    }
    arg =arg.replace(/,$/,"")
    arg += "]}";
    CreateXMLHttpRequest();
    xmlhttp.onreadystatechange = callhandle;
    xmlhttp.open("POST","/",true);
    xmlhttp.setRequestHeader("Content-Type","application/x-www-form-urlencoded");  //用POST的时候一定要有这句
    xmlhttp.send(arg);
}

function delete_adTask(_id)
{
    con = confirm("确定要删除吗?");
    if(con == false)
    {
        return;
    }
    var arg = JSON.stringify({'method':'delete_adTask','params':[_id]})
    CreateXMLHttpRequest();
    xmlhttp.onreadystatechange = callhandle;
    xmlhttp.open("POST","/",true);
    xmlhttp.setRequestHeader("Content-Type","application/x-www-form-urlencoded");  //用POST的时候一定要有这句
    xmlhttp.send(arg);
}

function CreateXMLHttpRequest()
{
    if (window.ActiveXObject)
    {
        xmlhttp = new ActiveXObject("Microsoft.XMLHTTP");
    }
    else if (window.XMLHttpRequest)
    {
        xmlhttp = new XMLHttpRequest();
    }
}
function callhandle()
{
    if (xmlhttp.readyState == 4 && xmlhttp.status == 200)
    {
        document.body.innerHTML=xmlhttp.responseText;
    }
}

function reload()
{
    setTimeout(get_task, 60000);
}


function get_task_by_ws()
{
    // alert("get_task_by_ws");
    host = document.location.hostname;
    port = document.location.port;
    uri = 'ws://'+ host + ':' + port +'/rdtask';

    ws = new WebSocket(uri);
    
    ws.onmessage = function(evt) {
        div = document.getElementById("tasks");
        if(div == null)
        {
            ws.close();
            return;
        }
        div.innerHTML = evt.data;

    }
}