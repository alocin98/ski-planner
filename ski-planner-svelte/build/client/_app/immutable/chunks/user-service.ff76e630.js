import{s as g,n as f}from"./scheduler.e108d1fd.js";import{S as y,i as w,B as c,C as u,j as l,f as i,k as n,a as v,x as h}from"./index.2b52f78d.js";import{s as b,c as k,a as m}from"./client.5020f725.js";import{a as p}from"./js.cookie.edb2da2a.js";function x(r){let e,t,s;return{c(){e=c("svg"),t=c("path"),s=c("animateTransform"),this.h()},l(a){e=u(a,"svg",{class:!0,version:!0,width:!0,id:!0,xmlns:!0,"xmlns:xlink":!0,x:!0,y:!0,viewBox:!0,"enable-background":!0,"xml:space":!0});var o=l(e);t=u(o,"path",{fill:!0,d:!0});var d=l(t);s=u(d,"animateTransform",{attributeName:!0,attributeType:!0,type:!0,dur:!0,from:!0,to:!0,repeatCount:!0}),l(s).forEach(i),d.forEach(i),o.forEach(i),this.h()},h(){n(s,"attributeName","transform"),n(s,"attributeType","XML"),n(s,"type","rotate"),n(s,"dur","1s"),n(s,"from","0 50 50"),n(s,"to","360 50 50"),n(s,"repeatCount","indefinite"),n(t,"fill","#fff"),n(t,"d","M73,50c0-12.7-10.3-23-23-23S27,37.3,27,50 M30.9,50c0-10.5,8.5-19.1,19.1-19.1S69.1,39.5,69.1,50"),n(e,"class",r[1]),n(e,"version","1.1"),n(e,"width",r[0]),n(e,"id","L9"),n(e,"xmlns","http://www.w3.org/2000/svg"),n(e,"xmlns:xlink","http://www.w3.org/1999/xlink"),n(e,"x","0px"),n(e,"y","0px"),n(e,"viewBox","0 0 100 100"),n(e,"enable-background","new 0 0 0 0"),n(e,"xml:space","preserve")},m(a,o){v(a,e,o),h(e,t),h(t,s)},p(a,[o]){o&2&&n(e,"class",a[1]),o&1&&n(e,"width",a[0])},i:f,o:f,d(a){a&&i(e)}}}function A(r,e,t){let{size:s=10}=e,{className:a=""}=e;return r.$$set=o=>{"size"in o&&t(0,s=o.size),"className"in o&&t(1,a=o.className)},[s,a]}class $ extends y{constructor(e){super(),w(this,e,A,x,g,{size:0,className:1})}}const S="register",E={"auth/invalid-email":"Email konnte nicht erkannt werden"},N="Datum",T="Preis",L="Geschmack",O="Kommentar",P="Stadt",_="Sun Loungery",U="Stop fighting for a spot by the pool. Just book it.",M="name",C="optional",q="unique name",z="street address",B="postal code",I="country",J="pool map",W={loginButton:"login",password:"password",email:"e-mail",noAccountYet:"Don't have an account yet? better get one!",alreadyHaveAnAccount:"Already have an account?",loginLink:"login"},j={register:S,firebaseErrors:E,date:N,price:T,taste:L,comment:O,city:P,title:_,subTitle:U,name:M,optional:C,uniqueName:q,streetAddress:z,postalCode:B,country:I,poolMap:J,login:W};let D=j;function ee(r){const e=r.split(".");let t=D;for(const s of e)if(t&&t.hasOwnProperty(s))t=t[s];else return;return t}async function F(r,e){return b(m,r,e).then(async t=>{const s={email:t.user.email??"NO-EMAIL"},a=await t.user.getIdToken();return p.set("AccessToken",a),G(s).then(()=>t)})}async function G(r){return fetch("/api/login",{credentials:"include",method:"POST",body:JSON.stringify(r)})}function H(r,e){return k(m,r,e).then(async t=>{const s={email:t.user.email??"NO-EMAIL"};return fetch("/api/user/login",{method:"POST",body:JSON.stringify(s)}).then(()=>t)})}function K(){return p.remove("AccessToken"),m.signOut()}function X(r,e){return fetch("/api/hotel/"+r+"/users",{method:"POST",body:JSON.stringify({email:e})})}function Y(r){return r||(r=window.fetch),r("/api/user",{credentials:"include"}).then(async e=>e.status===401?null:e.json()).catch(console.error)}const te={loginWithEmailAndPassword:F,newUserWithEmailAndPassword:H,getUser:Y,logout:K,generateUserAndSendPasswordEmail:X};export{$ as L,te as U,ee as t};