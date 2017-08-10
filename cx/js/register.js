function myFunction(){
	var a=document.getElementById("active").value;
 
    var b=document.getElementById("zhanghao").value;
  
    var c=document.getElementById("shoujihao").value;
 
    var d=document.getElementById("email").value;
 
    var e=document.getElementById("mima").value;
 
    var xhr = new XMLHttpRequest();
	xhr.open("POST", "http://localhost:8080/admin/register", true);
    xhr.setRequestHeader("Content-type","application/x-www-form-urlencoded");
    if(a.length==0){
    	a=0
    }
	xhr.send(
                "active="+a+"&zhanghao="+b+"&shoujihao="+c+"&email="+d+"&mima="+e
     )
	xhr.onreadystatechange = function () {
            if (xhr.readyState == 4 && xhr.status == 200) {
  				var json_msg=eval('(' + xhr.responseText + ')')        
                if(json_msg.res==0){
                	alert("注册成功")
                }else{
                	alert(json_msg.res)
                	
                }
            }

        }
}	
function xy(){ 
	 document.getElementById("active").style.display="none";
}
function gl(){
	 document.getElementById("active").style.display="block";
}
