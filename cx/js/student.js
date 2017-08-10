window.onload =function(){
	var name=getname()
	alert(name)
	var a=document.getElementById("hy")
    a.innerHTML=a.innerHTML+name	
}

function apply(){
	document.getElementById("wy").src='apply.html'
}
function result(){
	document.getElementById("wy").src='resultt.html'
	
}
function submits(){
	document.getElementById("wy").src='submit.html'
}
function zuxiao(){
	delCookie('name')
	window.location.href="login.html"
}
