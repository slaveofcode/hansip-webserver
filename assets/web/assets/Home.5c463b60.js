import{d as b,u as h,a as g,c as x,r as v,o as n,b as c,e,f as u,g as a,w as r,h as d,i as o,p as k,j as S}from"./index.9dec00f6.js";import{_ as w}from"./_plugin-vue_export-helper.cdc0426e.js";const y="/logo-256.png",m=t=>(k("data-v-33f45b0e"),t=t(),S(),t),C={class:"flex flex-col"},A=m(()=>e("div",{class:"flex flex-row justify-center items-center mb-3"},[e("img",{src:y,class:"w-40"})],-1)),H=m(()=>e("h1",{class:"mb-3"},[e("span",{class:"text-green-500 font-bold"},"Hansip"),o(" File Sharing")],-1)),I={class:"flex flex-row justify-center"},N={key:0},V=o("Create Account"),j=o("Login Account"),B={key:1},F=o("Share File"),L=b({__name:"Home",setup(t){const p=h(),_=g(),i=x(()=>_.isAuthenticated),f=()=>{_.logout(),p.go(0)};return(E,l)=>{const s=v("router-link");return n(),c("div",C,[A,H,e("div",I,[u(i)?d("",!0):(n(),c("div",N,[a(s,{to:{name:"create-account"},class:"mr-2 btn btn-blue"},{default:r(()=>[V]),_:1}),a(s,{to:{name:"login-account"},class:"btn btn-blue"},{default:r(()=>[j]),_:1})])),u(i)?(n(),c("div",B,[a(s,{to:{name:"file-share"},class:"mr-2 btn btn-orange"},{default:r(()=>[F]),_:1}),e("button",{onClick:l[0]||(l[0]=R=>f()),class:"btn btn-blue"},"Logout")])):d("",!0)])])}}});const q=w(L,[["__scopeId","data-v-33f45b0e"]]);export{q as default};
