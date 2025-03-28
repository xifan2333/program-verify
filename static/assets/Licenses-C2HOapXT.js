import{d as ce,e as l,f as ge,h as i,o as n,k as t,B as j,n as v,t as d,C as p,j as u,v as ye,F as M,x as N,D as z,q as g,p as H}from"./vendor-CFUBHfbh.js";import{u as xe}from"./index-DmvhG2uU.js";import{a as h,A as C}from"./config-BhZ7YNfN.js";import{u as G,w as me}from"./xlsx-ejz-vvm2.js";const fe={class:"space-y-6"},ke={class:"flex-between mb-4"},_e={class:"flex gap-2"},we={class:"relative"},he={key:0,class:"absolute top-full right-0 mt-1 w-32 bg-white dark:bg-gray-800 rounded-lg shadow-lg border dark:border-gray-700 z-10"},Ce={class:"py-1"},Se={class:"relative"},Ee={key:0,class:"absolute top-full right-0 mt-1 w-32 bg-white dark:bg-gray-800 rounded-lg shadow-lg border dark:border-gray-700 z-10"},Ue={class:"py-1"},De=["disabled"],Le={class:"bg-white dark:bg-gray-800 rounded-lg shadow-sm mb-4 overflow-hidden"},Fe={class:"flex items-center justify-between p-4 border-b dark:border-gray-700"},$e={class:"p-4 space-y-4"},Ve={class:"grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4"},Ie={class:"space-y-2"},Ae=["value"],Te={class:"space-y-2"},je={class:"space-y-2"},Me={class:"flex items-center space-x-2"},Ne={class:"space-y-2"},Oe={class:"flex items-center space-x-2"},Pe={class:"space-y-2"},Re={class:"card"},Be={key:0,class:"flex-center py-8"},qe={key:1,class:"overflow-x-auto"},ze={class:"w-full"},Ge={class:"px-4 py-2"},Qe={class:"px-4 py-2"},We={class:"px-4 py-2 font-mono"},He={class:"px-4 py-2"},Je={class:"px-4 py-2"},Ke={class:"px-4 py-2"},Xe={class:"px-4 py-2"},Ye={class:"px-4 py-2"},Ze={class:"px-4 py-2"},et={class:"flex gap-2"},tt=["onClick","disabled"],at=["onClick"],rt={class:"flex justify-between items-center mt-4 px-4"},st={class:"flex items-center gap-2"},lt=["value"],ot={key:0,class:"flex gap-2"},dt=["disabled"],it={class:"px-3 py-1"},nt=["disabled"],ut={key:1,class:"text-sm text-gray-600 dark:text-gray-400"},vt={key:0,class:"fixed inset-0 bg-black/50 flex-center"},bt={class:"bg-white dark:bg-gray-800 rounded-lg p-6 w-96"},pt=["value"],ct={class:"flex justify-end gap-2 my-4"},gt={key:1,class:"fixed inset-0 bg-black/50 flex-center"},yt={class:"bg-white dark:bg-gray-800 rounded-lg p-6 w-96"},xt={class:"flex justify-end gap-2 my-4"},ht=ce({__name:"Licenses",setup(mt){const o=xe(),O=l([]),P=l([]),R=l(!1),U=l(!1),D=l(!1),S=l(!1),L=l(!1),E=l(null),F=l(!1),$=l(!1),s=l({activation_status:"",enable_status:"",product_id:"",license_key:"",activated_start_date:"",activated_end_date:"",expires_start_date:"",expires_end_date:"",remark:""}),y=l({product_id:0,duration_days:365,count:1,remark:""}),c=l({remark:"",expires_at:""}),x=l(1),_=l(10),V=l(0),B=l(!1),J=[10,20,50,100,-1],k=l("enabled"),m=l("all"),K=()=>{switch(m.value){case"activated":return"已激活";case"inactive":return"未激活";case"expired":return"已过期";case"all":return"激活状态";default:return"激活状态"}},X=()=>{switch(m.value){case"activated":return"i-ri-checkbox-circle-line";case"inactive":return"i-ri-time-line";case"expired":return"i-ri-close-circle-line";case"all":return"i-ri-filter-3-line";default:return"i-ri-filter-3-line"}},Y=()=>{switch(m.value){case"activated":return"text-green-600";case"inactive":return"text-yellow-600";case"expired":return"text-red-600";case"all":return"text-blue-600";default:return"text-gray-600"}},I=r=>{m.value=r,s.value.activation_status=r==="all"?"":r,$.value=!1,A()},Z=()=>{switch(k.value){case"enabled":return"已启用";case"disabled":return"已禁用";case"all":return"启用状态";default:return"启用状态"}},ee=()=>{switch(k.value){case"enabled":return"i-ri-checkbox-circle-line";case"disabled":return"i-ri-close-circle-line";case"all":return"i-ri-list-check";default:return"i-ri-filter-3-line"}},te=()=>{switch(k.value){case"enabled":return"text-green-600";case"disabled":return"text-red-600";case"all":return"text-blue-600";default:return"text-gray-600"}},q=r=>{k.value=r,s.value.enable_status=r==="all"?"":r,F.value=!1,A()},ae=r=>{_.value=r,B.value=r===-1,x.value=1,f()},Q=r=>{x.value=r,f()},f=async()=>{R.value=!0;try{const r={...B.value?{}:{page:x.value.toString(),page_size:_.value.toString()},...s.value.activation_status&&{activation_status:s.value.activation_status},...s.value.enable_status&&{enable_status:s.value.enable_status},...s.value.product_id&&{product_id:s.value.product_id},...s.value.license_key&&{license_key:s.value.license_key},...s.value.activated_start_date&&{activated_start_date:s.value.activated_start_date},...s.value.activated_end_date&&{activated_end_date:s.value.activated_end_date},...s.value.expires_start_date&&{expires_start_date:s.value.expires_start_date},...s.value.expires_end_date&&{expires_end_date:s.value.expires_end_date},...s.value.remark&&{remark:s.value.remark}},e=await h.get(C.LICENSES.LIST,r);e.status===200?(O.value=e.data.items,V.value=e.data.total):o.error(e.message||"获取许可证列表失败")}catch{o.error("获取许可证列表失败，请稍后重试")}finally{R.value=!1}},re=async()=>{try{const r=await h.get(C.PRODUCTS.LIST,{status:"enabled"});r.status===200?P.value=r.data.items:o.error(r.message||"获取产品列表失败")}catch(r){console.error("获取产品列表失败:",r),o.error("获取产品列表失败，请稍后重试")}},se=async()=>{if(y.value.product_id)try{const r=await h.post(C.LICENSES.CREATE,y.value);r.status===200?(D.value=!1,y.value={product_id:0,duration_days:365,count:1,remark:""},o.success(r.message),f()):o.error(r.message||"生成失败")}catch{o.error("生成失败，请稍后重试")}},le=async r=>{try{const e=await h.put(C.LICENSES.UPDATE(r),{enable_status:"disabled"});e.status===200?(o.success(e.message),f()):o.error(e.message||"操作失败")}catch{o.error("操作失败，请稍后重试")}},oe=async r=>{try{const e=await h.put(C.LICENSES.UPDATE(r),{enable_status:"enabled"});e.status===200?(o.success(e.message),f()):o.error(e.message||"操作失败")}catch{o.error("操作失败，请稍后重试")}},de=r=>{E.value=r,c.value={remark:r.remark||"",expires_at:r.expires_at?new Date(r.expires_at).toISOString().slice(0,16):""},S.value=!0},ie=async()=>{if(E.value)try{const r={};if(c.value.remark!==E.value.remark&&(r.remark=c.value.remark),c.value.expires_at&&(r.expires_at=new Date(c.value.expires_at).toISOString()),Object.keys(r).length===0){S.value=!1;return}const e=await h.put(C.LICENSES.UPDATE(E.value.id),r);e.status===200?(S.value=!1,E.value=null,c.value={remark:"",expires_at:""},o.success(e.message),f()):o.error(e.message||"更新失败")}catch{o.error("更新失败，请稍后重试")}},A=()=>{x.value=1,f()},ne=()=>{s.value={activation_status:"",enable_status:"",product_id:"",license_key:"",activated_start_date:"",activated_end_date:"",expires_start_date:"",expires_end_date:"",remark:""},A()},ue=async()=>{U.value=!0;try{const r=["ID","产品","许可证密钥","有效期(天)","状态","激活时间","过期时间","备注"],e=O.value.map(b=>[b.id,b.product.name,b.license_key,b.duration_days,b.enable_status==="disabled"?"已禁用":b.activation_status==="expired"?"已过期":b.activation_status==="activated"?"已激活":"未激活",b.activated_at?new Date(b.activated_at).toLocaleString():"-",b.expires_at?new Date(b.expires_at).toLocaleString():"-",b.remark||"-"]),a=G.book_new(),T=G.aoa_to_sheet([r,...e]),ve=[{wch:8},{wch:20},{wch:32},{wch:10},{wch:8},{wch:20},{wch:20},{wch:20}];T["!cols"]=ve,G.book_append_sheet(a,T,"许可证列表");const be=me(a,{bookType:"xlsx",type:"array"}),pe=new Blob([be],{type:"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"}),w=document.createElement("a"),W=URL.createObjectURL(pe);w.setAttribute("href",W),w.setAttribute("download",`licenses_${new Date().toISOString().slice(0,10)}.xlsx`),w.style.visibility="hidden",document.body.appendChild(w),w.click(),document.body.removeChild(w),URL.revokeObjectURL(W),o.success("导出成功")}catch(r){o.error(r instanceof Error?r.message:"导出失败，请稍后重试")}finally{U.value=!1}};return ge(()=>{f(),re()}),(r,e)=>(n(),i("div",fe,[t("div",ke,[e[40]||(e[40]=t("h1",{class:"text-2xl font-bold"},"许可证管理",-1)),t("div",_e,[t("div",we,[t("button",{class:"flex items-center gap-2 px-3 py-1.5 rounded-lg border dark:border-gray-700 hover:bg-gray-100 dark:hover:bg-gray-800 transition-colors",onClick:e[0]||(e[0]=a=>$.value=!$.value)},[t("div",{class:v([X(),Y()])},null,2),t("span",null,d(K()),1),e[29]||(e[29]=t("div",{class:"i-ri-arrow-down-s-line text-sm"},null,-1))]),$.value?(n(),i("div",he,[t("div",Ce,[t("button",{class:v(["w-full px-4 py-2 text-left hover:bg-gray-100 dark:hover:bg-gray-700 flex items-center gap-2",m.value==="activated"?"text-green-600":"text-gray-600 dark:text-gray-300"]),onClick:e[1]||(e[1]=a=>I("activated"))},e[30]||(e[30]=[t("div",{class:"i-ri-checkbox-circle-line"},null,-1),p(" 已激活 ")]),2),t("button",{class:v(["w-full px-4 py-2 text-left hover:bg-gray-100 dark:hover:bg-gray-700 flex items-center gap-2",m.value==="inactive"?"text-yellow-600":"text-gray-600 dark:text-gray-300"]),onClick:e[2]||(e[2]=a=>I("inactive"))},e[31]||(e[31]=[t("div",{class:"i-ri-time-line"},null,-1),p(" 未激活 ")]),2),t("button",{class:v(["w-full px-4 py-2 text-left hover:bg-gray-100 dark:hover:bg-gray-700 flex items-center gap-2",m.value==="expired"?"text-red-600":"text-gray-600 dark:text-gray-300"]),onClick:e[3]||(e[3]=a=>I("expired"))},e[32]||(e[32]=[t("div",{class:"i-ri-close-circle-line"},null,-1),p(" 已过期 ")]),2),t("button",{class:v(["w-full px-4 py-2 text-left hover:bg-gray-100 dark:hover:bg-gray-700 flex items-center gap-2",m.value==="all"?"text-blue-600":"text-gray-600 dark:text-gray-300"]),onClick:e[4]||(e[4]=a=>I("all"))},e[33]||(e[33]=[t("div",{class:"i-ri-list-check"},null,-1),p(" 全部 ")]),2)])])):j("",!0)]),t("div",Se,[t("button",{class:"flex items-center gap-2 px-3 py-1.5 rounded-lg border dark:border-gray-700 hover:bg-gray-100 dark:hover:bg-gray-800 transition-colors",onClick:e[5]||(e[5]=a=>F.value=!F.value)},[t("div",{class:v([ee(),te()])},null,2),t("span",null,d(Z()),1),e[34]||(e[34]=t("div",{class:"i-ri-arrow-down-s-line text-sm"},null,-1))]),F.value?(n(),i("div",Ee,[t("div",Ue,[t("button",{class:v(["w-full px-4 py-2 text-left hover:bg-gray-100 dark:hover:bg-gray-700 flex items-center gap-2",k.value==="enabled"?"text-green-600":"text-gray-600 dark:text-gray-300"]),onClick:e[6]||(e[6]=a=>q("enabled"))},e[35]||(e[35]=[t("div",{class:"i-ri-checkbox-circle-line"},null,-1),p(" 已启用 ")]),2),t("button",{class:v(["w-full px-4 py-2 text-left hover:bg-gray-100 dark:hover:bg-gray-700 flex items-center gap-2",k.value==="disabled"?"text-red-600":"text-gray-600 dark:text-gray-300"]),onClick:e[7]||(e[7]=a=>q("disabled"))},e[36]||(e[36]=[t("div",{class:"i-ri-close-circle-line"},null,-1),p(" 已禁用 ")]),2),t("button",{class:v(["w-full px-4 py-2 text-left hover:bg-gray-100 dark:hover:bg-gray-700 flex items-center gap-2",k.value==="all"?"text-blue-600":"text-gray-600 dark:text-gray-300"]),onClick:e[8]||(e[8]=a=>q("all"))},e[37]||(e[37]=[t("div",{class:"i-ri-list-check"},null,-1),p(" 全部 ")]),2)])])):j("",!0)]),t("button",{class:"btn bg-blue-400 text-white hover:bg-blue-500",onClick:ue,disabled:U.value},[e[38]||(e[38]=t("div",{class:"i-ri-download-line mr-1"},null,-1)),p(" "+d(U.value?"导出中...":"导出"),1)],8,De),t("button",{class:"btn btn-primary",onClick:e[9]||(e[9]=a=>D.value=!0)},e[39]||(e[39]=[t("i",{class:"i-ri-add-line"},null,-1),p(" 生成许可证 ")]))])]),t("div",Le,[t("div",Fe,[e[41]||(e[41]=t("h3",{class:"text-lg font-medium text-gray-900 dark:text-gray-100"},"筛选条件",-1)),t("button",{class:"text-gray-500 hover:text-gray-700 dark:text-gray-400 dark:hover:text-gray-200 transition-colors",onClick:e[10]||(e[10]=a=>L.value=!L.value)},[t("div",{class:v(["i-ri-arrow-up-s-line text-xl",{"rotate-180":!L.value}])},null,2)])]),u(t("div",$e,[t("div",Ve,[t("div",Ie,[e[43]||(e[43]=t("label",{class:"block text-sm font-medium text-gray-700 dark:text-gray-300"},"产品",-1)),u(t("select",{"onUpdate:modelValue":e[11]||(e[11]=a=>s.value.product_id=a),class:"block w-full px-3 py-2 bg-white dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-primary-500"},[e[42]||(e[42]=t("option",{value:""},"全部",-1)),(n(!0),i(M,null,N(P.value,a=>(n(),i("option",{key:a.id,value:a.id},d(a.name),9,Ae))),128))],512),[[z,s.value.product_id]])]),t("div",Te,[e[44]||(e[44]=t("label",{class:"block text-sm font-medium text-gray-700 dark:text-gray-300"},"许可证密钥",-1)),u(t("input",{"onUpdate:modelValue":e[12]||(e[12]=a=>s.value.license_key=a),type:"text",placeholder:"输入许可证密钥",class:"block w-full px-3 py-2 bg-white dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-primary-500"},null,512),[[g,s.value.license_key]])]),t("div",je,[e[46]||(e[46]=t("label",{class:"block text-sm font-medium text-gray-700 dark:text-gray-300"},"激活时间",-1)),t("div",Me,[u(t("input",{"onUpdate:modelValue":e[13]||(e[13]=a=>s.value.activated_start_date=a),type:"date",class:"block w-full px-3 py-2 bg-white dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-primary-500"},null,512),[[g,s.value.activated_start_date]]),e[45]||(e[45]=t("span",{class:"text-gray-500 dark:text-gray-400"},"-",-1)),u(t("input",{"onUpdate:modelValue":e[14]||(e[14]=a=>s.value.activated_end_date=a),type:"date",class:"block w-full px-3 py-2 bg-white dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-primary-500"},null,512),[[g,s.value.activated_end_date]])])]),t("div",Ne,[e[48]||(e[48]=t("label",{class:"block text-sm font-medium text-gray-700 dark:text-gray-300"},"过期时间",-1)),t("div",Oe,[u(t("input",{"onUpdate:modelValue":e[15]||(e[15]=a=>s.value.expires_start_date=a),type:"date",class:"block w-full px-3 py-2 bg-white dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-primary-500"},null,512),[[g,s.value.expires_start_date]]),e[47]||(e[47]=t("span",{class:"text-gray-500 dark:text-gray-400"},"-",-1)),u(t("input",{"onUpdate:modelValue":e[16]||(e[16]=a=>s.value.expires_end_date=a),type:"date",class:"block w-full px-3 py-2 bg-white dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-primary-500"},null,512),[[g,s.value.expires_end_date]])])]),t("div",Pe,[e[49]||(e[49]=t("label",{class:"block text-sm font-medium text-gray-700 dark:text-gray-300"},"备注",-1)),u(t("input",{"onUpdate:modelValue":e[17]||(e[17]=a=>s.value.remark=a),type:"text",placeholder:"输入备注信息",class:"block w-full px-3 py-2 bg-white dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-primary-500"},null,512),[[g,s.value.remark]])])]),t("div",{class:"flex justify-end space-x-3 mt-4 pt-4 border-t dark:border-gray-700"},[t("button",{onClick:ne,class:"btn mr-4"},e[50]||(e[50]=[t("div",{class:"i-ri-refresh-line mr-1"},null,-1),p(" 重置 ")])),t("button",{onClick:A,class:"btn btn-primary"},e[51]||(e[51]=[t("div",{class:"i-ri-filter-3-line mr-1"},null,-1),p(" 筛选 ")]))])],512),[[ye,L.value]])]),t("div",Re,[R.value?(n(),i("div",Be,e[52]||(e[52]=[t("div",{class:"i-ri-loader-4-line animate-spin text-2xl text-primary"},null,-1)]))):(n(),i("div",qe,[t("table",ze,[e[54]||(e[54]=t("thead",null,[t("tr",{class:"border-b dark:border-gray-700"},[t("th",{class:"px-4 py-2 text-left"},"ID"),t("th",{class:"px-4 py-2 text-left"},"产品"),t("th",{class:"px-4 py-2 text-left"},"许可证密钥"),t("th",{class:"px-4 py-2 text-left"},"有效期(天)"),t("th",{class:"px-4 py-2 text-left"},"状态"),t("th",{class:"px-4 py-2 text-left"},"激活时间"),t("th",{class:"px-4 py-2 text-left"},"过期时间"),t("th",{class:"px-4 py-2 text-left"},"备注"),t("th",{class:"px-4 py-2 text-left"},"操作")])],-1)),t("tbody",null,[(n(!0),i(M,null,N(O.value,a=>(n(),i("tr",{key:a.id,class:"border-b dark:border-gray-700"},[t("td",Ge,d(a.id),1),t("td",Qe,d(a.product.name),1),t("td",We,d(a.license_key),1),t("td",He,d(a.duration_days),1),t("td",Je,[t("span",{class:v(["px-2 py-1 rounded-full text-xs",{"bg-green-100 text-green-800":a.activation_status==="activated"&&a.enable_status==="enabled","bg-yellow-100 text-yellow-800":a.activation_status==="inactive"&&a.enable_status==="enabled","bg-red-100 text-red-800":a.enable_status==="disabled","bg-gray-100 text-gray-800":a.activation_status==="expired"}])},d(a.enable_status==="disabled"?"已禁用":a.activation_status==="expired"?"已过期":a.activation_status==="activated"?"已激活":"未激活"),3)]),t("td",Ke,d(a.activated_at?new Date(a.activated_at).toLocaleString():"-"),1),t("td",Xe,d(a.expires_at?new Date(a.expires_at).toLocaleString():"-"),1),t("td",Ye,d(a.remark||"-"),1),t("td",Ze,[t("div",et,[t("button",{class:"icon-btn",onClick:T=>de(a),disabled:a.enable_status==="disabled"},e[53]||(e[53]=[t("div",{class:"i-ri-edit-line"},null,-1)]),8,tt),t("button",{class:v(["icon-btn",a.enable_status==="enabled"?"text-red-500 hover:text-red-600":"text-green-500 hover:text-green-600"]),onClick:T=>a.enable_status==="enabled"?le(a.id):oe(a.id)},[t("div",{class:v(a.enable_status==="enabled"?"i-ri-close-circle-line":"i-ri-checkbox-circle-line")},null,2)],10,at)])])]))),128))])])])),t("div",rt,[t("div",st,[e[55]||(e[55]=t("span",{class:"text-sm text-gray-600 dark:text-gray-400"},"每页显示",-1)),u(t("select",{"onUpdate:modelValue":e[18]||(e[18]=a=>_.value=a),class:"px-2 py-1 text-sm rounded border dark:border-gray-700 bg-white dark:bg-gray-800",onChange:e[19]||(e[19]=a=>ae(Number(a.target.value)))},[(n(),i(M,null,N(J,a=>t("option",{key:a,value:a},d(a===-1?"全部":a),9,lt)),64))],544),[[z,_.value]]),e[56]||(e[56]=t("span",{class:"text-sm text-gray-600 dark:text-gray-400"},"条",-1))]),B.value?(n(),i("div",ut," 共 "+d(V.value)+" 条数据 ",1)):(n(),i("div",ot,[t("button",{class:"px-3 py-1 rounded border dark:border-gray-700",disabled:x.value===1,onClick:e[20]||(e[20]=a=>Q(x.value-1))}," 上一页 ",8,dt),t("span",it," 第 "+d(x.value)+" 页 / 共 "+d(Math.ceil(V.value/_.value))+" 页 ",1),t("button",{class:"px-3 py-1 rounded border dark:border-gray-700",disabled:x.value>=Math.ceil(V.value/_.value),onClick:e[21]||(e[21]=a=>Q(x.value+1))}," 下一页 ",8,nt)]))])]),D.value?(n(),i("div",vt,[t("div",bt,[e[62]||(e[62]=t("h2",{class:"text-xl font-bold mb-4"},"生成许可证",-1)),t("form",{onSubmit:H(se,["prevent"]),class:"space-y-4"},[t("div",null,[e[58]||(e[58]=t("label",{class:"block text-sm font-medium mb-1"},"选择产品",-1)),u(t("select",{"onUpdate:modelValue":e[22]||(e[22]=a=>y.value.product_id=a),class:"w-full px-4 py-2 rounded-lg border dark:border-gray-700 bg-white dark:bg-gray-900",required:""},[e[57]||(e[57]=t("option",{value:"0"},"请选择产品",-1)),(n(!0),i(M,null,N(P.value,a=>(n(),i("option",{key:a.id,value:a.id},d(a.name),9,pt))),128))],512),[[z,y.value.product_id]])]),t("div",null,[e[59]||(e[59]=t("label",{class:"block text-sm font-medium mb-1"},"有效期(天)",-1)),u(t("input",{"onUpdate:modelValue":e[23]||(e[23]=a=>y.value.duration_days=a),type:"number",class:"w-full px-4 py-2 rounded-lg border dark:border-gray-700 bg-white dark:bg-gray-900",required:""},null,512),[[g,y.value.duration_days]])]),t("div",null,[e[60]||(e[60]=t("label",{class:"block text-sm font-medium mb-1"},"生成数量",-1)),u(t("input",{"onUpdate:modelValue":e[24]||(e[24]=a=>y.value.count=a),type:"number",min:"1",max:"100",class:"w-full px-4 py-2 rounded-lg border dark:border-gray-700 bg-white dark:bg-gray-900",required:""},null,512),[[g,y.value.count]])]),t("div",ct,[t("button",{type:"button",class:"btn",onClick:e[25]||(e[25]=a=>D.value=!1)}," 取消 "),e[61]||(e[61]=t("button",{type:"submit",class:"btn btn-primary"}," 生成 ",-1))])],32)])])):j("",!0),S.value?(n(),i("div",gt,[t("div",yt,[e[66]||(e[66]=t("h2",{class:"text-xl font-bold mb-4"},"编辑许可证",-1)),t("form",{onSubmit:H(ie,["prevent"]),class:"space-y-4"},[t("div",null,[e[63]||(e[63]=t("label",{class:"block text-sm font-medium mb-1"},"过期时间",-1)),u(t("input",{"onUpdate:modelValue":e[26]||(e[26]=a=>c.value.expires_at=a),type:"datetime-local",class:"w-full px-4 py-2 rounded-lg border dark:border-gray-700 bg-white dark:bg-gray-900"},null,512),[[g,c.value.expires_at]])]),t("div",null,[e[64]||(e[64]=t("label",{class:"block text-sm font-medium mb-1"},"备注",-1)),u(t("input",{"onUpdate:modelValue":e[27]||(e[27]=a=>c.value.remark=a),type:"text",placeholder:"输入备注信息",class:"w-full px-4 py-2 rounded-lg border dark:border-gray-700 bg-white dark:bg-gray-900"},null,512),[[g,c.value.remark]])]),t("div",xt,[t("button",{type:"button",class:"btn",onClick:e[28]||(e[28]=a=>S.value=!1)}," 取消 "),e[65]||(e[65]=t("button",{type:"submit",class:"btn btn-primary"}," 保存 ",-1))])],32)])])):j("",!0)]))}});export{ht as default};
