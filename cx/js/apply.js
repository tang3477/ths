function ss(){

	var a=document.getElementById("xmname")

	var b=document.getElementById("xmtype")

	var c=document.getElementById("xmpeople")

	var d=document.getElementById("direct")
	var xhr = new XMLHttpRequest()
	var e=document.getElementById("cont")
	var s=getname()
	
	xhr.open("POST", "http://localhost:8080/admin/apply", true);
    xhr.setRequestHeader("Content-type","application/x-www-form-urlencoded");
	xhr.send(
            "xmname="+a.value+"&xmtype="+b.value+"&xmpeople="+c.value+"&direct="+d.value+"&cont="+e.value+"&acco="+s
        )
xhr.onreadystatechange = function () {
            if (xhr.readyState == 4 && xhr.status == 200) {
                var json_msg=eval('(' + xhr.responseText + ')')     
                if(json_msg.res==0){
                	alert("发送成功")
				}
        	}
        }
}
