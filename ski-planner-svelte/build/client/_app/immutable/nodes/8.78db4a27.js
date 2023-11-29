import{s as he,r as ye,n as te}from"../chunks/scheduler.e108d1fd.js";import{S as ge,i as ve,g as d,m as S,s as L,h as p,j as h,n as H,f as u,c as N,k as a,a as F,x as r,A as C,z as ee,p as Re,t as $,b as Ue,d as q,o as Ae,r as be,u as we,v as Ee,w as ke,y as Ie}from"../chunks/index.2b52f78d.js";import{g as Le}from"../chunks/navigation.e0f8f435.js";import{t as O,U as Ne,L as Pe}from"../chunks/user-service.ff76e630.js";function _e(n){let e,l;return{c(){e=d("p"),l=S(n[2]),this.h()},l(t){e=p(t,"P",{class:!0});var s=h(e);l=H(s,n[2]),s.forEach(u),this.h()},h(){a(e,"class","text-error")},m(t,s){F(t,e,s),r(e,l)},p(t,s){s&4&&Ae(l,t[2])},d(t){t&&u(e)}}}function Te(n){let e=O("register")+"",l;return{c(){l=S(e)},l(t){l=H(t,e)},m(t,s){F(t,l,s)},i:te,o:te,d(t){t&&u(l)}}}function De(n){let e,l;return e=new Pe({props:{className:"fill-current",size:30}}),{c(){be(e.$$.fragment)},l(t){we(e.$$.fragment,t)},m(t,s){Ee(e,t,s),l=!0},i(t){l||(q(e.$$.fragment,t),l=!0)},o(t){$(e.$$.fragment,t),l=!1},d(t){ke(e,t)}}}function Ve(n){let e,l,t,s,o,R=O("login.email")+"",g,_,i,U,f,P,D,le=O("login.password")+"",G,J,b,K,W,k,A,w,E,Q,B,ae=O("login.alreadyHaveAnAccount")+"",X,Y,I,re=O("login.loginLink")+"",Z,M,x,se,m=n[2]!==null&&_e(n);const ne=[De,Te],T=[];function oe(c,y){return c[3]?0:1}return w=oe(n),E=T[w]=ne[w](n),{c(){e=d("div"),l=d("form"),t=d("div"),s=d("label"),o=d("span"),g=S(R),_=L(),i=d("input"),U=L(),f=d("div"),P=d("label"),D=d("span"),G=S(le),J=L(),b=d("input"),K=L(),m&&m.c(),W=L(),k=d("div"),A=d("button"),E.c(),Q=L(),B=d("p"),X=S(ae),Y=L(),I=d("a"),Z=S(re),this.h()},l(c){e=p(c,"DIV",{class:!0});var y=h(e);l=p(y,"FORM",{class:!0});var v=h(l);t=p(v,"DIV",{class:!0});var j=h(t);s=p(j,"LABEL",{for:!0,class:!0});var ie=h(s);o=p(ie,"SPAN",{class:!0});var ce=h(o);g=H(ce,R),ce.forEach(u),ie.forEach(u),_=N(j),i=p(j,"INPUT",{id:!0,type:!0,autocomplete:!0,placeholder:!0,class:!0}),j.forEach(u),U=N(v),f=p(v,"DIV",{class:!0});var z=h(f);P=p(z,"LABEL",{for:!0,class:!0});var ue=h(P);D=p(ue,"SPAN",{class:!0});var fe=h(D);G=H(fe,le),fe.forEach(u),ue.forEach(u),J=N(z),b=p(z,"INPUT",{id:!0,type:!0,placeholder:!0,class:!0}),z.forEach(u),K=N(v),m&&m.l(v),W=N(v),k=p(v,"DIV",{class:!0});var V=h(k);A=p(V,"BUTTON",{class:!0});var de=h(A);E.l(de),de.forEach(u),Q=N(V),B=p(V,"P",{});var pe=h(B);X=H(pe,ae),pe.forEach(u),Y=N(V),I=p(V,"A",{href:!0,class:!0});var me=h(I);Z=H(me,re),me.forEach(u),V.forEach(u),v.forEach(u),y.forEach(u),this.h()},h(){a(o,"class","label-text"),a(s,"for","email-input"),a(s,"class","label"),a(i,"id","email-input"),a(i,"type","email"),a(i,"autocomplete","email"),a(i,"placeholder","email"),a(i,"class","input input-bordered"),a(t,"class","form-control"),a(D,"class","label-text"),a(P,"for","password-input"),a(P,"class","label"),a(b,"id","password-input"),a(b,"type","password"),a(b,"placeholder","password"),a(b,"class","input input-bordered"),a(f,"class","form-control"),a(A,"class","btn btn-secondary"),a(I,"href",n[0]),a(I,"class","link"),a(k,"class","form-control mt-2 gap-2"),a(l,"class","card-body"),a(e,"class","card w-96 bg-base-100 shadow-xl")},m(c,y){F(c,e,y),r(e,l),r(l,t),r(t,s),r(s,o),r(o,g),r(t,_),r(t,i),C(i,n[1].email),r(l,U),r(l,f),r(f,P),r(P,D),r(D,G),r(f,J),r(f,b),C(b,n[1].password),r(l,K),m&&m.m(l,null),r(l,W),r(l,k),r(k,A),T[w].m(A,null),r(k,Q),r(k,B),r(B,X),r(k,Y),r(k,I),r(I,Z),M=!0,x||(se=[ee(i,"input",n[6]),ee(b,"input",n[7]),ee(A,"click",n[4])],x=!0)},p(c,[y]){y&2&&i.value!==c[1].email&&C(i,c[1].email),y&2&&b.value!==c[1].password&&C(b,c[1].password),c[2]!==null?m?m.p(c,y):(m=_e(c),m.c(),m.m(l,W)):m&&(m.d(1),m=null);let v=w;w=oe(c),w!==v&&(Re(),$(T[v],1,1,()=>{T[v]=null}),Ue(),E=T[w],E||(E=T[w]=ne[w](c),E.c()),q(E,1),E.m(A,null)),(!M||y&1)&&a(I,"href",c[0])},i(c){M||(q(E),M=!0)},o(c){$(E),M=!1},d(c){c&&u(e),m&&m.d(),T[w].d(),x=!1,ye(se)}}}function Se(n,e,l){let{afterRegister:t}=e,{loginUrl:s="/login"}=e;const o={email:"",password:""};let R=null,g=!1;function _(){l(3,g=!0),Ne.newUserWithEmailAndPassword(o.email,o.password).catch(f=>{console.log("STH WENT WRONG"),l(2,R=f.code)}).then(f=>{f&&t(f)}).finally(()=>{l(3,g=!1)})}function i(){o.email=this.value,l(1,o)}function U(){o.password=this.value,l(1,o)}return n.$$set=f=>{"afterRegister"in f&&l(5,t=f.afterRegister),"loginUrl"in f&&l(0,s=f.loginUrl)},[s,o,R,g,_,t,i,U]}class He extends ge{constructor(e){super(),ve(this,e,Se,Ve,he,{afterRegister:5,loginUrl:0})}}function Be(n){let e,l,t,s,o,R='<img src="images/hero.jpeg" alt="mockup"/>',g;return t=new He({props:{afterRegister:n[0]}}),{c(){e=d("div"),l=d("div"),be(t.$$.fragment),s=L(),o=d("div"),o.innerHTML=R,this.h()},l(_){e=p(_,"DIV",{class:!0});var i=h(e);l=p(i,"DIV",{class:!0});var U=h(l);we(t.$$.fragment,U),U.forEach(u),s=N(i),o=p(i,"DIV",{class:!0,"data-svelte-h":!0}),Ie(o)!=="svelte-115o7sw"&&(o.innerHTML=R),i.forEach(u),this.h()},h(){a(l,"class","mx-auto place-self-center lg:col-span-7"),a(o,"class","hidden lg:mt-0 lg:col-span-5 lg:flex"),a(e,"class","grid lg:gap-8 xl:gap-0 lg:py-16 lg:grid-cols-12")},m(_,i){F(_,e,i),r(e,l),Ee(t,l,null),r(e,s),r(e,o),g=!0},p:te,i(_){g||(q(t.$$.fragment,_),g=!0)},o(_){$(t.$$.fragment,_),g=!1},d(_){_&&u(e),ke(t)}}}function Me(n){return[()=>{Le("/home")}]}class Ce extends ge{constructor(e){super(),ve(this,e,Me,Be,he,{})}}export{Ce as component};