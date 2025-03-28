import{d as de,e as n,f as ie,h as u,o as b,k as t,B as A,n as m,t as c,C as x,j as d,v as ue,D as B,q as p,F as z,x as N,p as Q}from"./vendor-CFUBHfbh.js";import{u as be}from"./index-DmvhG2uU.js";import{a as w,A as h}from"./config-BhZ7YNfN.js";import{u as M,w as ce}from"./xlsx-ejz-vvm2.js";const pe={class:"space-y-6"},ge={class:"flex-between mb-4"},ve={class:"flex gap-2"},me={class:"relative"},ye={key:0,class:"absolute top-full right-0 mt-1 w-32 bg-white dark:bg-gray-800 rounded-lg shadow-lg border dark:border-gray-700 z-10"},xe={class:"py-1"},fe=["disabled"],ke={class:"bg-white dark:bg-gray-800 rounded-lg shadow-sm mb-4 overflow-hidden"},we={class:"flex items-center justify-between p-4 border-b dark:border-gray-700"},he={class:"p-4 space-y-4"},_e={class:"grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4"},Ce={class:"space-y-2"},Se={class:"space-y-2"},Ue={class:"space-y-2"},Pe={class:"flex items-center space-x-2"},Ee={class:"space-y-2"},De={class:"flex items-center space-x-2"},Te={class:"card"},$e={key:0,class:"flex-center py-8"},Ve={key:1,class:"overflow-x-auto"},Fe={class:"w-full"},Re={class:"px-4 py-2"},Le={class:"px-4 py-2"},Oe={class:"px-4 py-2"},Ae={class:"px-4 py-2"},Me={class:"px-4 py-2"},je={class:"px-4 py-2"},Ie={class:"flex gap-2"},qe=["onClick","disabled"],Be=["onClick"],ze={class:"flex justify-between items-center mt-4 px-4"},Ne={class:"flex items-center gap-2"},Qe=["value"],We={key:0,class:"flex gap-2"},Ge=["disabled"],He={class:"px-3 py-1"},Je=["disabled"],Ke={key:1,class:"text-sm text-gray-600 dark:text-gray-400"},Xe={key:0,class:"fixed inset-0 bg-black/50 flex-center"},Ye={class:"bg-white dark:bg-gray-800 rounded-lg p-6 w-96"},Ze={class:"flex justify-end gap-2 my-4"},et={key:1,class:"fixed inset-0 bg-black/50 flex-center"},tt={class:"bg-white dark:bg-gray-800 rounded-lg p-6 w-96"},at={class:"flex justify-end gap-2 my-4"},dt=de({__name:"Products",setup(st){const o=be(),$=n([]),V=n(!1),_=n(!1),C=n(!1),S=n(!1),U=n(null),P=n(!1),E=n(!1),r=n({status:"enabled",name:"",min_price:"",max_price:"",start_date:"",end_date:""}),y=n("enabled"),l=n({name:"",price:0}),i=n(1),g=n(10),D=n(0),F=n(!1),W=[10,20,50,100,-1],G=()=>{const s=new URLSearchParams(window.location.search);r.value={status:s.get("status")||"",name:s.get("name")||"",min_price:s.get("min_price")||"",max_price:s.get("max_price")||"",start_date:s.get("start_date")||"",end_date:s.get("end_date")||""},i.value=parseInt(s.get("page")||"1"),g.value=parseInt(s.get("page_size")||"10")},R=()=>{const s=new URLSearchParams;r.value.status&&s.set("status",r.value.status),r.value.name&&s.set("name",r.value.name),r.value.min_price&&s.set("min_price",r.value.min_price),r.value.max_price&&s.set("max_price",r.value.max_price),r.value.start_date&&s.set("start_date",r.value.start_date),r.value.end_date&&s.set("end_date",r.value.end_date),s.set("page",i.value.toString()),s.set("page_size",g.value.toString());const e=`${window.location.pathname}?${s.toString()}`;window.history.pushState({},"",e)},H=s=>{g.value=s,F.value=s===-1,i.value=1,R(),v()},L=()=>{i.value=1,R(),v()},j=s=>{i.value=s,R(),v()},v=async()=>{V.value=!0;try{const s={...F.value?{}:{page:i.value.toString(),page_size:g.value.toString()},...r.value.status&&{status:r.value.status},...r.value.name&&{name:r.value.name},...r.value.min_price&&{min_price:r.value.min_price},...r.value.max_price&&{max_price:r.value.max_price},...r.value.start_date&&{start_date:r.value.start_date},...r.value.end_date&&{end_date:r.value.end_date}},e=await w.get(h.PRODUCTS.LIST,s);$.value=e.data.items,D.value=e.data.total}catch(s){o.error(s instanceof Error?s.message:"获取产品列表失败，请稍后重试")}finally{V.value=!1}},I=()=>!l.value.name||l.value.price<=0?(o.warning("请填写完整的产品信息"),!1):!0,J=async()=>{if(I())try{await w.post(h.PRODUCTS.CREATE,l.value),C.value=!1,l.value={name:"",price:0},v()}catch(s){o.error(s instanceof Error?s.message:"创建失败，请稍后重试")}},K=async()=>{if(U.value&&I())try{const s=await w.put(h.PRODUCTS.UPDATE(U.value.id),l.value);s.status===200?(S.value=!1,U.value=null,l.value={name:"",price:0},v()):o.error(s.message||"更新失败")}catch(s){o.error(s instanceof Error?s.message:"更新失败，请稍后重试")}},X=async s=>{try{const e=await w.put(h.PRODUCTS.UPDATE(s),{status:"disabled"});e.status===200?(o.success(e.message),v()):o.error(e.message||"禁用产品失败")}catch(e){o.error(e instanceof Error?e.message:"操作失败，请稍后重试")}},Y=async s=>{try{const e=await w.put(h.PRODUCTS.UPDATE(s),{status:"enabled"});e.status===200?(o.success(e.message),v()):o.error(e.message||"重新启用产品失败")}catch(e){o.error(e instanceof Error?e.message:"操作失败，请稍后重试")}},Z=s=>{U.value=s,l.value={name:s.name,price:s.price},S.value=!0},ee=()=>{r.value={status:"enabled",name:"",min_price:"",max_price:"",start_date:"",end_date:""},L()},te=async()=>{_.value=!0;try{const s=["ID","名称","价格","状态","创建时间"],e=$.value.map(k=>[k.id,k.name,k.price,k.status==="enabled"?"启用":"禁用",new Date(k.created_at).toLocaleString()]),a=M.book_new(),T=M.aoa_to_sheet([s,...e]),le=[{wch:8},{wch:20},{wch:10},{wch:8},{wch:20}];T["!cols"]=le,M.book_append_sheet(a,T,"产品列表");const ne=ce(a,{bookType:"xlsx",type:"array"}),oe=new Blob([ne],{type:"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"}),f=document.createElement("a"),q=URL.createObjectURL(oe);f.setAttribute("href",q),f.setAttribute("download",`products_${new Date().toISOString().slice(0,10)}.xlsx`),f.style.visibility="hidden",document.body.appendChild(f),f.click(),document.body.removeChild(f),URL.revokeObjectURL(q),o.success("导出成功")}catch(s){o.error(s instanceof Error?s.message:"导出失败，请稍后重试")}finally{_.value=!1}},O=s=>{y.value=s,r.value.status=s==="all"?"":s,E.value=!1,L()},ae=()=>{switch(y.value){case"enabled":return"启用";case"disabled":return"禁用";case"all":return"启用状态";default:return"启用状态"}},se=()=>{switch(y.value){case"enabled":return"i-ri-checkbox-circle-line";case"disabled":return"i-ri-close-circle-line";case"all":return"i-ri-list-check";default:return"i-ri-filter-3-line"}},re=()=>{switch(y.value){case"enabled":return"text-green-600";case"disabled":return"text-red-600";case"all":return"text-blue-600";default:return"text-gray-600"}};return ie(()=>{G(),v()}),(s,e)=>(b(),u("div",pe,[t("div",ge,[e[28]||(e[28]=t("h1",{class:"text-2xl font-bold"},"产品管理",-1)),t("div",ve,[t("div",me,[t("button",{class:"flex items-center gap-2 px-3 py-1.5 rounded-lg border dark:border-gray-700 hover:bg-gray-100 dark:hover:bg-gray-800 transition-colors",onClick:e[0]||(e[0]=a=>E.value=!E.value)},[t("div",{class:m([se(),re()])},null,2),t("span",null,c(ae()),1),e[22]||(e[22]=t("div",{class:"i-ri-arrow-down-s-line text-sm"},null,-1))]),E.value?(b(),u("div",ye,[t("div",xe,[t("button",{class:m(["w-full px-4 py-2 text-left hover:bg-gray-100 dark:hover:bg-gray-700 flex items-center gap-2",y.value==="enabled"?"text-green-600":"text-gray-600 dark:text-gray-300"]),onClick:e[1]||(e[1]=a=>O("enabled"))},e[23]||(e[23]=[t("div",{class:"i-ri-checkbox-circle-line"},null,-1),x(" 启用 ")]),2),t("button",{class:m(["w-full px-4 py-2 text-left hover:bg-gray-100 dark:hover:bg-gray-700 flex items-center gap-2",y.value==="disabled"?"text-red-600":"text-gray-600 dark:text-gray-300"]),onClick:e[2]||(e[2]=a=>O("disabled"))},e[24]||(e[24]=[t("div",{class:"i-ri-close-circle-line"},null,-1),x(" 禁用 ")]),2),t("button",{class:m(["w-full px-4 py-2 text-left hover:bg-gray-100 dark:hover:bg-gray-700 flex items-center gap-2",y.value==="all"?"text-blue-600":"text-gray-600 dark:text-gray-300"]),onClick:e[3]||(e[3]=a=>O("all"))},e[25]||(e[25]=[t("div",{class:"i-ri-list-check"},null,-1),x(" 全部 ")]),2)])])):A("",!0)]),t("button",{class:"btn bg-blue-400 text-white hover:bg-blue-500",onClick:te,disabled:_.value},[e[26]||(e[26]=t("div",{class:"i-ri-download-line mr-1"},null,-1)),x(" "+c(_.value?"导出中...":"导出"),1)],8,fe),t("button",{class:"btn btn-primary",onClick:e[4]||(e[4]=a=>C.value=!0)},e[27]||(e[27]=[t("i",{class:"i-ri-add-line"},null,-1),x(" 创建产品 ")]))])]),t("div",ke,[t("div",we,[e[29]||(e[29]=t("h3",{class:"text-lg font-medium text-gray-900 dark:text-gray-100"},"筛选条件",-1)),t("button",{class:"text-gray-500 hover:text-gray-700 dark:text-gray-400 dark:hover:text-gray-200 transition-colors",onClick:e[5]||(e[5]=a=>P.value=!P.value)},[t("div",{class:m(["i-ri-arrow-up-s-line text-xl",{"rotate-180":!P.value}])},null,2)])]),d(t("div",he,[t("div",_e,[t("div",Ce,[e[31]||(e[31]=t("label",{class:"block text-sm font-medium text-gray-700 dark:text-gray-300"},"状态",-1)),d(t("select",{"onUpdate:modelValue":e[6]||(e[6]=a=>r.value.status=a),class:"block w-full px-3 py-2 bg-white dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-primary-500"},e[30]||(e[30]=[t("option",{value:""},"全部",-1),t("option",{value:"enabled"},"启用",-1),t("option",{value:"disabled"},"禁用",-1)]),512),[[B,r.value.status]])]),t("div",Se,[e[32]||(e[32]=t("label",{class:"block text-sm font-medium text-gray-700 dark:text-gray-300"},"名称",-1)),d(t("input",{"onUpdate:modelValue":e[7]||(e[7]=a=>r.value.name=a),type:"text",placeholder:"输入产品名称",class:"block w-full px-3 py-2 bg-white dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-primary-500"},null,512),[[p,r.value.name]])]),t("div",Ue,[e[34]||(e[34]=t("label",{class:"block text-sm font-medium text-gray-700 dark:text-gray-300"},"价格范围",-1)),t("div",Pe,[d(t("input",{"onUpdate:modelValue":e[8]||(e[8]=a=>r.value.min_price=a),type:"number",placeholder:"最低价",class:"block w-full px-3 py-2 bg-white dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-primary-500"},null,512),[[p,r.value.min_price]]),e[33]||(e[33]=t("span",{class:"text-gray-500 dark:text-gray-400"},"-",-1)),d(t("input",{"onUpdate:modelValue":e[9]||(e[9]=a=>r.value.max_price=a),type:"number",placeholder:"最高价",class:"block w-full px-3 py-2 bg-white dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-primary-500"},null,512),[[p,r.value.max_price]])])]),t("div",Ee,[e[36]||(e[36]=t("label",{class:"block text-sm font-medium text-gray-700 dark:text-gray-300"},"创建时间",-1)),t("div",De,[d(t("input",{"onUpdate:modelValue":e[10]||(e[10]=a=>r.value.start_date=a),type:"date",class:"block w-full px-3 py-2 bg-white dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-primary-500"},null,512),[[p,r.value.start_date]]),e[35]||(e[35]=t("span",{class:"text-gray-500 dark:text-gray-400"},"-",-1)),d(t("input",{"onUpdate:modelValue":e[11]||(e[11]=a=>r.value.end_date=a),type:"date",class:"block w-full px-3 py-2 bg-white dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-primary-500"},null,512),[[p,r.value.end_date]])])])]),t("div",{class:"flex justify-end space-x-3 mt-4 pt-4 border-t dark:border-gray-700"},[t("button",{onClick:ee,class:"btn mr-4"},e[37]||(e[37]=[t("div",{class:"i-ri-refresh-line mr-1"},null,-1),x(" 重置 ")])),t("button",{onClick:L,class:"btn btn-primary"},e[38]||(e[38]=[t("div",{class:"i-ri-filter-3-line mr-1"},null,-1),x(" 筛选 ")]))])],512),[[ue,P.value]])]),t("div",Te,[V.value?(b(),u("div",$e,e[39]||(e[39]=[t("div",{class:"i-ri-loader-4-line animate-spin text-2xl text-primary"},null,-1)]))):(b(),u("div",Ve,[t("table",Fe,[e[41]||(e[41]=t("thead",null,[t("tr",{class:"border-b dark:border-gray-700"},[t("th",{class:"px-4 py-2 text-left"},"ID"),t("th",{class:"px-4 py-2 text-left"},"名称"),t("th",{class:"px-4 py-2 text-left"},"价格"),t("th",{class:"px-4 py-2 text-left"},"状态"),t("th",{class:"px-4 py-2 text-left"},"创建时间"),t("th",{class:"px-4 py-2 text-left"},"操作")])],-1)),t("tbody",null,[(b(!0),u(z,null,N($.value,a=>(b(),u("tr",{key:a.id,class:"border-b dark:border-gray-700"},[t("td",Re,c(a.id),1),t("td",Le,c(a.name),1),t("td",Oe,"¥"+c(a.price),1),t("td",Ae,[t("span",{class:m(["px-2 py-1 rounded-full text-xs",a.status==="enabled"?"bg-green-100 text-green-800":"bg-red-100 text-red-800"])},c(a.status==="enabled"?"启用":"禁用"),3)]),t("td",Me,c(new Date(a.created_at).toLocaleString()),1),t("td",je,[t("div",Ie,[t("button",{class:"icon-btn",onClick:T=>Z(a),disabled:a.status!=="enabled"},e[40]||(e[40]=[t("div",{class:"i-ri-edit-line"},null,-1)]),8,qe),t("button",{class:m(["icon-btn",a.status==="enabled"?"text-red-500 hover:text-red-600":"text-green-500 hover:text-green-600"]),onClick:T=>a.status==="enabled"?X(a.id):Y(a.id)},[t("div",{class:m(a.status==="enabled"?"i-ri-close-circle-line":"i-ri-checkbox-circle-line")},null,2)],10,Be)])])]))),128))])])])),t("div",ze,[t("div",Ne,[e[42]||(e[42]=t("span",{class:"text-sm text-gray-600 dark:text-gray-400"},"每页显示",-1)),d(t("select",{"onUpdate:modelValue":e[12]||(e[12]=a=>g.value=a),class:"px-2 py-1 text-sm rounded border dark:border-gray-700 bg-white dark:bg-gray-800",onChange:e[13]||(e[13]=a=>H(Number(a.target.value)))},[(b(),u(z,null,N(W,a=>t("option",{key:a,value:a},c(a===-1?"全部":a),9,Qe)),64))],544),[[B,g.value]]),e[43]||(e[43]=t("span",{class:"text-sm text-gray-600 dark:text-gray-400"},"条",-1))]),F.value?(b(),u("div",Ke," 共 "+c(D.value)+" 条数据 ",1)):(b(),u("div",We,[t("button",{class:"px-3 py-1 rounded border dark:border-gray-700",disabled:i.value===1,onClick:e[14]||(e[14]=a=>j(i.value-1))}," 上一页 ",8,Ge),t("span",He," 第 "+c(i.value)+" 页 / 共 "+c(Math.ceil(D.value/g.value))+" 页 ",1),t("button",{class:"px-3 py-1 rounded border dark:border-gray-700",disabled:i.value>=Math.ceil(D.value/g.value),onClick:e[15]||(e[15]=a=>j(i.value+1))}," 下一页 ",8,Je)]))])]),C.value?(b(),u("div",Xe,[t("div",Ye,[e[47]||(e[47]=t("h2",{class:"text-xl font-bold mb-4"},"创建产品",-1)),t("form",{onSubmit:Q(J,["prevent"]),class:"space-y-4"},[t("div",null,[e[44]||(e[44]=t("label",{class:"block text-sm font-medium mb-1"},"产品名称",-1)),d(t("input",{"onUpdate:modelValue":e[16]||(e[16]=a=>l.value.name=a),type:"text",class:"w-full px-4 py-2 rounded-lg border dark:border-gray-700 bg-white dark:bg-gray-900",required:""},null,512),[[p,l.value.name]])]),t("div",null,[e[45]||(e[45]=t("label",{class:"block text-sm font-medium mb-1"},"价格",-1)),d(t("input",{"onUpdate:modelValue":e[17]||(e[17]=a=>l.value.price=a),type:"number",step:"0.01",class:"w-full px-4 py-2 rounded-lg border dark:border-gray-700 bg-white dark:bg-gray-900",required:""},null,512),[[p,l.value.price]])]),t("div",Ze,[t("button",{type:"button",class:"btn",onClick:e[18]||(e[18]=a=>C.value=!1)}," 取消 "),e[46]||(e[46]=t("button",{type:"submit",class:"btn btn-primary"}," 创建 ",-1))])],32)])])):A("",!0),S.value?(b(),u("div",et,[t("div",tt,[e[51]||(e[51]=t("h2",{class:"text-xl font-bold mb-4"},"编辑产品",-1)),t("form",{onSubmit:Q(K,["prevent"]),class:"space-y-4"},[t("div",null,[e[48]||(e[48]=t("label",{class:"block text-sm font-medium mx-2"},"产品名称",-1)),d(t("input",{"onUpdate:modelValue":e[19]||(e[19]=a=>l.value.name=a),type:"text",class:"w-full px-4 py-2 rounded-lg border dark:border-gray-700 bg-white dark:bg-gray-900",required:""},null,512),[[p,l.value.name]])]),t("div",null,[e[49]||(e[49]=t("label",{class:"block text-sm font-medium mx-2"},"价格",-1)),d(t("input",{"onUpdate:modelValue":e[20]||(e[20]=a=>l.value.price=a),type:"number",step:"0.01",class:"w-full px-4 py-2 rounded-lg border dark:border-gray-700 bg-white dark:bg-gray-900",required:""},null,512),[[p,l.value.price]])]),t("div",at,[t("button",{type:"button",class:"btn",onClick:e[21]||(e[21]=a=>S.value=!1)}," 取消 "),e[50]||(e[50]=t("button",{type:"submit",class:"btn btn-primary"}," 保存 ",-1))])],32)])])):A("",!0)]))}});export{dt as default};
