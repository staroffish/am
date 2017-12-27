var xmlhttp;
function main(arg)
{
    var arg = JSON.stringify({'method':'main'})
    CreateXMLHttpRequest();
    xmlhttp.onreadystatechange = callhandle;
    //xmlhttp.open("GET","Default.aspx?goback=yes&arg=" + arg,true);
    xmlhttp.open("POST","/",true);
    xmlhttp.setRequestHeader("Content-Type","application/x-www-form-urlencoded");  //用POST的时候一定要有这句
    xmlhttp.send(arg);
    
}
function send2(arg)
{
    var arg = JSON.stringify({'method':'main','params':[2]})
    CreateXMLHttpRequest();
    xmlhttp.onreadystatechange = callhandle;
    //xmlhttp.open("GET","Default.aspx?goback=yes&arg=" + arg,true);
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