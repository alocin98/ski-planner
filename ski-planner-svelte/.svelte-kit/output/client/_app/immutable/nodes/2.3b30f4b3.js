import{s as T,x as _,z as y,y as h,K as ee,D as H,E as S,h as F,d as f,j as p,i as m,P as te,e as v,T as re,r as W,a as P,c as z,u as A,v as G,w as I,f as K,g as N,R,l as le,m as ae}from"../chunks/scheduler.e7d17915.js";import{S as D,i as E,b as x,d as w,m as C,a as g,t as b,e as j,g as J,c as O}from"../chunks/index.0b315888.js";import{t as B,T as oe,g as Q,a as U,f as se,F as ne}from"../chunks/ToolbarButton.b4dd5b23.js";function ie(n){let e,a,l;return{c(){e=H("svg"),a=H("path"),this.h()},l(r){e=S(r,"svg",{class:!0,fill:!0,viewBox:!0,xmlns:!0});var t=F(e);a=S(t,"path",{"fill-rule":!0,d:!0,"clip-rule":!0}),F(a).forEach(f),t.forEach(f),this.h()},h(){p(a,"fill-rule","evenodd"),p(a,"d","M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z"),p(a,"clip-rule","evenodd"),p(e,"class",l=n[4]),p(e,"fill","currentColor"),p(e,"viewBox","0 0 20 20"),p(e,"xmlns","http://www.w3.org/2000/svg")},m(r,t){m(r,e,t),te(e,a)},p(r,t){t&16&&l!==(l=r[4])&&p(e,"class",l)},d(r){r&&f(e)}}}function ue(n){let e,a;const l=[{name:n[0]},n[1],{class:B("ml-auto",n[2].class)}];let r={$$slots:{default:[ie,({svgSize:t})=>({4:t}),({svgSize:t})=>t?16:0]},$$scope:{ctx:n}};for(let t=0;t<l.length;t+=1)r=_(r,l[t]);return e=new oe({props:r}),e.$on("click",n[3]),{c(){x(e.$$.fragment)},l(t){w(e.$$.fragment,t)},m(t,o){C(e,t,o),a=!0},p(t,[o]){const s=o&7?Q(l,[o&1&&{name:t[0]},o&2&&U(t[1]),o&4&&{class:B("ml-auto",t[2].class)}]):{};o&48&&(s.$$scope={dirty:o,ctx:t}),e.$set(s)},i(t){a||(g(e.$$.fragment,t),a=!0)},o(t){b(e.$$.fragment,t),a=!1},d(t){j(e,t)}}}function ge(n,e,a){const l=["name"];let r=y(e,l),{name:t="Close"}=e;function o(s){ee.call(this,n,s)}return n.$$set=s=>{a(2,e=_(_({},e),h(s))),a(1,r=y(e,l)),"name"in s&&a(0,t=s.name)},e=h(e),[t,r,e,o]}class de extends D{constructor(e){super(),E(this,e,ge,ue,T,{name:0})}}const fe=n=>({}),$=n=>({close:n[5]});function q(n){let e,a;const l=[n[6],{class:n[4]}];let r={$$slots:{default:[be]},$$scope:{ctx:n}};for(let t=0;t<l.length;t+=1)r=_(r,l[t]);return e=new ne({props:r}),{c(){x(e.$$.fragment)},l(t){w(e.$$.fragment,t)},m(t,o){C(e,t,o),a=!0},p(t,o){const s=o&80?Q(l,[o&64&&U(t[6]),o&16&&{class:t[4]}]):{};o&263&&(s.$$scope={dirty:o,ctx:t}),e.$set(s)},i(t){a||(g(e.$$.fragment,t),a=!0)},o(t){b(e.$$.fragment,t),a=!1},d(t){j(e,t)}}}function V(n){let e;const a=n[7]["close-button"],l=W(a,n,n[8],$),r=l||ce(n);return{c(){r&&r.c()},l(t){r&&r.l(t)},m(t,o){r&&r.m(t,o),e=!0},p(t,o){l?l.p&&(!e||o&256)&&A(l,a,t,t[8],e?I(a,t[8],o,fe):G(t[8]),$):r&&r.p&&(!e||o&3)&&r.p(t,e?o:-1)},i(t){e||(g(r,t),e=!0)},o(t){b(r,t),e=!1},d(t){r&&r.d(t)}}}function ce(n){let e,a;return e=new de({props:{color:n[0],size:n[1]?"sm":"xs",name:"Remove badge",class:"ml-1.5 -mr-1.5"}}),e.$on("click",n[5]),{c(){x(e.$$.fragment)},l(l){w(e.$$.fragment,l)},m(l,r){C(e,l,r),a=!0},p(l,r){const t={};r&1&&(t.color=l[0]),r&2&&(t.size=l[1]?"sm":"xs"),e.$set(t)},i(l){a||(g(e.$$.fragment,l),a=!0)},o(l){b(e.$$.fragment,l),a=!1},d(l){j(e,l)}}}function be(n){let e,a,l;const r=n[7].default,t=W(r,n,n[8],null);let o=n[2]&&V(n);return{c(){t&&t.c(),e=P(),o&&o.c(),a=v()},l(s){t&&t.l(s),e=z(s),o&&o.l(s),a=v()},m(s,u){t&&t.m(s,u),m(s,e,u),o&&o.m(s,u),m(s,a,u),l=!0},p(s,u){t&&t.p&&(!l||u&256)&&A(t,r,s,s[8],l?I(r,s[8],u,null):G(s[8]),null),s[2]?o?(o.p(s,u),u&4&&g(o,1)):(o=V(s),o.c(),g(o,1),o.m(a.parentNode,a)):o&&(J(),b(o,1,1,()=>{o=null}),O())},i(s){l||(g(t,s),g(o),l=!0)},o(s){b(t,s),b(o),l=!1},d(s){s&&(f(e),f(a)),t&&t.d(s),o&&o.d(s)}}}function me(n){let e,a,l=n[3]&&q(n);return{c(){l&&l.c(),e=v()},l(r){l&&l.l(r),e=v()},m(r,t){l&&l.m(r,t),m(r,e,t),a=!0},p(r,[t]){r[3]?l?(l.p(r,t),t&8&&g(l,1)):(l=q(r),l.c(),g(l,1),l.m(e.parentNode,e)):l&&(J(),b(l,1,1,()=>{l=null}),O())},i(r){a||(g(l),a=!0)},o(r){b(l),a=!1},d(r){r&&f(e),l&&l.d(r)}}}const pe="font-medium inline-flex items-center justify-center px-2.5 py-0.5";function _e(n,e,a){const l=["color","large","dismissable"];let r=y(e,l),{$$slots:t={},$$scope:o}=e,{color:s="primary"}=e,{large:u=!1}=e,{dismissable:i=!1}=e;const c={primary:"bg-primary-100 text-primary-800 dark:bg-primary-900 dark:text-primary-300",blue:"bg-blue-100 text-blue-800 dark:bg-blue-900 dark:text-blue-300",dark:"bg-gray-100 text-gray-800 dark:bg-gray-700 dark:text-gray-300",gray:"bg-gray-100 text-gray-800 dark:bg-gray-700 dark:text-gray-300",red:"bg-red-100 text-red-800 dark:bg-red-900 dark:text-red-300",green:"bg-green-100 text-green-800 dark:bg-green-900 dark:text-green-300",yellow:"bg-yellow-100 text-yellow-800 dark:bg-yellow-900 dark:text-yellow-300",indigo:"bg-indigo-100 text-indigo-800 dark:bg-indigo-900 dark:text-indigo-300",purple:"bg-purple-100 text-purple-800 dark:bg-purple-900 dark:text-purple-300",pink:"bg-pink-100 text-pink-800 dark:bg-pink-900 dark:text-pink-300",none:""},k={primary:"bg-primary-100 text-primary-800 dark:bg-gray-700 dark:text-primary-400 border-primary-400 dark:border-primary-400",blue:"bg-blue-100 text-blue-800 dark:bg-gray-700 dark:text-blue-400 border-blue-400 dark:border-blue-400",dark:"bg-gray-100 text-gray-800 dark:bg-gray-700 dark:text-gray-400 border-gray-500 dark:border-gray-500",red:"bg-red-100 text-red-800 dark:bg-gray-700 dark:text-red-400 border-red-400 dark:border-red-400",green:"bg-green-100 text-green-800 dark:bg-gray-700 dark:text-green-400 border-green-400 dark:border-green-400",yellow:"bg-yellow-100 text-yellow-800 dark:bg-gray-700 dark:text-yellow-300 border-yellow-300 dark:border-yellow-300",indigo:"bg-indigo-100 text-indigo-800 dark:bg-gray-700 dark:text-indigo-400 border-indigo-400 dark:border-indigo-400",purple:"bg-purple-100 text-purple-800 dark:bg-gray-700 dark:text-purple-400 border-purple-400 dark:border-purple-400",pink:"bg-pink-100 text-pink-800 dark:bg-gray-700 dark:text-pink-400 border-pink-400 dark:border-pink-400",none:""},X={primary:"hover:bg-primary-200",blue:"hover:bg-blue-200",dark:"hover:bg-gray-200",red:"hover:bg-red-200",green:"hover:bg-green-200",yellow:"hover:bg-yellow-200",indigo:"hover:bg-indigo-200",purple:"hover:bg-purple-200",pink:"hover:bg-pink-200",none:""};let M,L=!0;const Y=re(),Z=d=>{d.stopPropagation(),a(3,L=!1)};return n.$$set=d=>{a(14,e=_(_({},e),h(d))),a(6,r=y(e,l)),"color"in d&&a(0,s=d.color),"large"in d&&a(1,u=d.large),"dismissable"in d&&a(2,i=d.dismissable),"$$scope"in d&&a(8,o=d.$$scope)},n.$$.update=()=>{i&&a(6,r.transition=r.transition??se,r),a(4,M=B(pe,u?"text-sm":"text-xs",r.border?`border ${k[s]}`:c[s],r.href&&X[s],r.rounded?"rounded-full":"rounded",e.class)),n.$$.dirty&8&&Y(L?"open":"close")},e=h(e),[s,u,i,L,M,Z,r,t,o]}class ke extends D{constructor(e){super(),E(this,e,_e,me,T,{color:0,large:1,dismissable:2})}}function ye(n){let e;return{c(){e=le("Default")},l(a){e=ae(a,"Default")},m(a,l){m(a,e,l)},d(a){a&&f(e)}}}function he(n){let e,a="Welcome to SvelteKit",l,r,t='Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation',o,s,u;return s=new ke({props:{$$slots:{default:[ye]},$$scope:{ctx:n}}}),{c(){e=K("h1"),e.textContent=a,l=P(),r=K("p"),r.innerHTML=t,o=P(),x(s.$$.fragment)},l(i){e=N(i,"H1",{"data-svelte-h":!0}),R(e)!=="svelte-yyjjjs"&&(e.textContent=a),l=z(i),r=N(i,"P",{"data-svelte-h":!0}),R(r)!=="svelte-jl9sbz"&&(r.innerHTML=t),o=z(i),w(s.$$.fragment,i)},m(i,c){m(i,e,c),m(i,l,c),m(i,r,c),m(i,o,c),C(s,i,c),u=!0},p(i,[c]){const k={};c&1&&(k.$$scope={dirty:c,ctx:i}),s.$set(k)},i(i){u||(g(s.$$.fragment,i),u=!0)},o(i){b(s.$$.fragment,i),u=!1},d(i){i&&(f(e),f(l),f(r),f(o)),j(s,i)}}}class Ce extends D{constructor(e){super(),E(this,e,null,he,T,{})}}export{Ce as component};
