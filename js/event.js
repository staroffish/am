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

function edit_anime(_id, prepage)
{
    var arg = JSON.stringify({'method':'edit_anime','params':[_id,prepage,false]})
    CreateXMLHttpRequest();
    xmlhttp.onreadystatechange = callhandle;
    xmlhttp.open("POST","/",true);
    xmlhttp.setRequestHeader("Content-Type","application/x-www-form-urlencoded");  //用POST的时候一定要有这句
    xmlhttp.send(arg);
}

function update_anime(_id, prepage)
{
    var elems = document.getElementsByTagName("input");
    var arg = "{\"method\":\"update_anime\",\"params\":[\"" + _id + "\",\"" + prepage + "\",true,";
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
    xmlhttp.onreadystatechange = updatehandle;
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
        document.body.innerHTML=xmlhttp.responseText 
    }
}

function updatehandle()
{
    
    if (xmlhttp.readyState == 4 && xmlhttp.status == 200)
    {
        var resp = xmlhttp.responseText
        var funcStr = resp.substr(0, resp.indexOf("("));
        var paramStr = resp.substr(resp.indexOf("(") + 1, resp.indexOf(")"));
        var fn = window[funcStr];
        var start = 0;
        var end = 0;
        start = paramStr.indexOf("'", start) + 1;
        end = paramStr.indexOf("'", start);
        var id = paramStr.substr(start, end - start);
        start = end + 1;
        start = paramStr.indexOf("'", start) + 1;
        end = paramStr.indexOf("'", start);
        var prePage = paramStr.substr(start, end - start);
        start = end + 1;
        fn(id, prePage);
    }
}