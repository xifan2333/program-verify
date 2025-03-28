import{d as g,u as y,e as f,r as o,h as i,o as r,k as e,s as d,n as l,F as _,x as w,t as c,i as u,w as x}from"./vendor-CFUBHfbh.js";const k={class:"min-h-screen bg-gray-100 dark:bg-gray-900"},C={class:"flex items-center justify-between border-b dark:border-gray-700"},j={class:"space-y-2"},B={class:"whitespace-nowrap"},L={class:"h-16 bg-white dark:bg-gray-800 shadow-sm flex items-center justify-between px-6"},S={class:"text-lg font-semibold"},$={class:"p-6"},I=g({__name:"Layout",setup(z){const p=y(),s=f(!1),h=()=>{localStorage.removeItem("token"),p.push("/login")},m=[{path:"/",icon:"i-ri-home-4-line",label:"首页"},{path:"/products",icon:"i-ri-apps-2-line",label:"产品管理"},{path:"/licenses",icon:"i-ri-key-2-line",label:"许可证管理"}];return(n,t)=>{const v=o("router-link"),b=o("router-view");return r(),i("div",k,[e("aside",{class:l(["fixed left-0 top-0 h-screen bg-white dark:bg-gray-800 shadow-lg transition-all duration-300",s.value?"w-16":"w-64"])},[e("div",C,[e("div",{class:"overflow-hidden transition-all duration-300",style:d({width:s.value?"0":"160px"})},t[1]||(t[1]=[e("h1",{class:"text-xl font-bold text-primary whitespace-nowrap py-4 pl-4"},"软件授权",-1)]),4),e("button",{class:"icon-btn",onClick:t[0]||(t[0]=a=>s.value=!s.value)},[e("div",{class:l(s.value?"i-ri-menu-unfold-line mr-6":"i-ri-menu-fold-line mr-4")},null,2)])]),e("nav",{class:l(s.value?"px-2":"px-4")},[e("ul",j,[(r(),i(_,null,w(m,a=>e("li",{key:a.path},[u(v,{to:a.path,class:l(["flex items-center rounded-lg transition-all duration-300 hover:bg-gray-100 my-2 p-3",[{"bg-primary/10 text-primary":n.$route.path===a.path},s.value?"justify-center":"gap-3"]])},{default:x(()=>[e("div",{class:l([a.icon,"text-xl"])},null,2),e("div",{class:"overflow-hidden transition-all duration-300",style:d({width:s.value?"0":"120px"})},[e("span",B,c(a.label),1)],4)]),_:2},1032,["to","class"])])),64))])],2)],2),e("main",{class:l(["transition-all duration-300 min-h-screen",s.value?"ml-16":"ml-64"])},[e("header",L,[e("h2",S,c(n.$route.name),1),e("button",{class:"icon-btn",onClick:h},t[2]||(t[2]=[e("div",{class:"i-ri-logout-box-line"},null,-1)]))]),e("div",$,[u(b)])],2)])}}});export{I as default};
