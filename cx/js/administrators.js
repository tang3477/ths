window.onload =function(){
	var name=getname()
	alert(name)
	var a=document.getElementById("hy")
    a.innerHTML=a.innerHTML+name	
}
function zuxiao(){
	delCookie('name')
	window.location.href="login.html"

}
function analysis(){
	document.getElementById("wy").src='anal.html'
	
}
function receive(){
	document.getElementById("wy").src='xiazai.html'
}
function examine(){
	document.getElementById("wy").src='jianyi.html'
}
