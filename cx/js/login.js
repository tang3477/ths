
function login(){
  
	var name=document.getElementById("name")
	var nanm=document.getElementById("nam")
	var c=document.getElementById("xna").value
	var xhr = new XMLHttpRequest();
    xhr.open("POST", "http://localhost:8080/admin/login", true);
    xhr.setRequestHeader("Content-type","application/x-www-form-urlencoded");
	xhr.send(
                "username="+name.value+"&password="+nanm.value+"&xiala="+c
        )
	xhr.onreadystatechange = function () {
            if (xhr.readyState == 4 && xhr.status == 200) {
                var json_msg=eval('(' + xhr.responseText + ')')     
                if(json_msg.res==0){
                	if(c==0){
                	setCookie("name",name.value,10)
                	window.location.href="Student.html"
                	}else{
                		setCookie("name",name.value,10)
                	window.location.href="administrators.html"
                	}
                }else{
                	alert(json_msg.res)
                	
                }
            }

        }
}
function register(){
	window.location.href="register.html"
}
