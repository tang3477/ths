window.onload =function(){
	var xhr = new XMLHttpRequest();
    xhr.open("POST", "http://localhost:8080/admin/anal", true);

    xhr.setRequestHeader("Content-type","application/x-www-form-urlencoded");
    xhr.send(
    	
    )

	var element=document.getElementById("tbo");
	
	xhr.onreadystatechange = function () {
            if (xhr.readyState == 4 && xhr.status == 200) {
                var json_msg=eval('(' + xhr.responseText + ')')
               console.log(json_msg)
                for(var i=0;i<json_msg.data.length;i++){
                	var trr = document.createElement("tr")
              		trr.setAttribute("id",json_msg.data[i].id)
                	var td1 = document.createElement("td")
                	var span1 = document.createElement("span")
                	span1.innerText = json_msg.data[i].nam
                	td1.appendChild(span1)
                	trr.appendChild(td1)
                	var td2 = document.createElement("td")
                	var span2 = document.createElement("span")
                	span2.innerText = json_msg.data[i].account
                	td2.appendChild(span2)
                	trr.appendChild(td2)
                	var td3 = document.createElement("td")
                	var span3 = document.createElement("span")
                	span3.innerText = new Date(parseInt(json_msg.data[i].tim) * 1000).toLocaleString().substr(0,16)
                	td3.appendChild(span3)
                	trr.appendChild(td3)
 
                	
                	var td4 = document.createElement("td")
                	var span4 = document.createElement("span")
                	span4.innerText = "查看"
                	trr.appendChild(td4)
                	span4.setAttribute("class", "look")
                	span4.addEventListener('click', function(){look(this)}, false);
                	td4.appendChild(span4)
                	element.appendChild(trr)
                	
                	
//					var para=document.createElement("p");
//					var a=document.createTextNode(json_msg.data[i].nam);
//					var ab=document.createTextNode(" ; ");
//					var b=document.createTextNode(json_msg.data[i].account)
//					var bc=document.createTextNode("  ; ");
//					var c=document.createTextNode(new Date(parseInt(json_msg.data[i].tim) * 1000).toLocaleString().substr(0,16))
//					para.appendChild(a);para.appendChild(ab);para.appendChild(b);para.appendChild(bc);para.appendChild(c)
//					var element=document.getElementById("anal");
//					element.appendChild(para);
                }
             }
           }
	
}


function look(item){
	a=item.parentNode.parentNode.id
	var xhr = new XMLHttpRequest();
    xhr.open("POST", "http://localhost:8080/admin/analid", true);
	var an_id=document.getElementById("an_id")
	an_id.setAttribute("id",a)
	console.log(an_id.id)
    xhr.setRequestHeader("Content-type","application/x-www-form-urlencoded");
    xhr.send(
    	"anid="+a
    )
    	xhr.onreadystatechange = function () {
            if (xhr.readyState == 4 && xhr.status == 200) {
                var json_msg=eval('(' + xhr.responseText + ')')
                
                var a=document.getElementById("nam")
                a.innerText=json_msg.data[0].nam

                var b=document.getElementById("lx")              
                b.innerHTML=json_msg.data[0].lx
                
                var c=document.getElementById("ppe")              
                c.innerHTML=json_msg.data[0].ppe
                
                var d=document.getElementById("dire")          
                d.innerHTML=json_msg.data[0].dire
                
                var e=document.getElementById("cntent")            
                e.innerHTML=json_msg.data[0].cntent
                
                var f=document.getElementById("wenj")
                f.innerHTML=json_msg.data[0].dizhi
                if(json_msg.data[0].adopt==1){
                	var f=document.getElementById("yess")           
                	f.innerHTML="申请成功"
                }
         }       
    }
    document.getElementById('ceng').style.display = 'block';
    document.getElementById('close').style.display = 'block';
}
    	function closeCeng() {
        	document.getElementById('ceng').style.display = 'none';
       		document.getElementById('close').style.display = 'none';
        
         
    }
    	

function yes(item){
	var xhr = new XMLHttpRequest();
	xhr.open("POST", "http://localhost:8080/admin/yes", true);
    xhr.setRequestHeader("Content-type","application/x-www-form-urlencoded");
    xhr.send(
    	"anid="+item.previousSibling.previousSibling.id
    )
    
    xhr.onreadystatechange = function () {
            if (xhr.readyState == 4 && xhr.status == 200) {
                var json_msg=eval('(' + xhr.responseText + ')')
                	if (json_msg.res==0){
                		alert("修改成功")
                	}
               }
           }
    
}

function no(item){
	var xhr = new XMLHttpRequest();
    xhr.open("POST", "http://localhost:8080/admin/no", true);
    xhr.setRequestHeader("Content-type","application/x-www-form-urlencoded");
    xhr.send(
    	"anid="+item.previousSibling.previousSibling.previousSibling.id
    )
    
    xhr.onreadystatechange = function () {
            if (xhr.readyState == 4 && xhr.status == 200) {
                var json_msg=eval('(' + xhr.responseText + ')')
                	if (json_msg.res==0){
                		alert("修改成功")
                	}
               }
           }
}
function xiazai(){
	var f=document.getElementById("wenj").innerHTML
	var xhr = new XMLHttpRequest();
    xhr.open("POST", "http://localhost:8080/admin/glxz", true);
    xhr.setRequestHeader("Content-type","application/x-www-form-urlencoded");
    xhr.send(
    	"xz="+f
    )
      xhr.onreadystatechange = function () {
            if (xhr.readyState == 4 && xhr.status == 200) {
                var json_msg=eval('(' + xhr.responseText + ')')
                	if (json_msg.res==0){
                		alert("已下载")
                	}
               }
           }
}
