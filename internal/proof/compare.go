package proof

// comparePage walks each mockup frame and its proof screen in parallel at
// several frame widths. Same tree, so item i of one is item i of the other;
// text nodes are items too, so a text run is compared by its own box.
const comparePage = `<!doctype html><html lang="ru"><head><meta charset="utf-8"><title>Proof · frames vs mockup</title>
<style>
body{margin:0;padding:16px 20px;font:13px/1.45 system-ui,sans-serif;background:#f5f1e8;color:#17211c}
h1{font-size:16px;margin:0 0 4px}
#sum{margin:10px 0;border-collapse:collapse}
#sum td,#sum th{border:1px solid #d8cfbb;padding:4px 10px;text-align:center}
#sum td.ok{background:#e3eadf}#sum td.bad{background:#f1dcd6;cursor:pointer}#sum td{cursor:pointer}
.pair{display:flex;gap:16px;margin:12px 0;align-items:flex-start}
.pair figure{margin:0}.pair figcaption{font-size:12px;color:#56503f;margin-bottom:4px}
.clip{overflow:hidden;border:1px solid #d8cfbb;background:#fff;position:relative}
.clip iframe{border:0;transform:scale(.5);transform-origin:0 0;position:absolute;left:0;top:0}
table.d{border-collapse:collapse;font-size:12px}
table.d td,table.d th{border-bottom:1px solid #e2daca;padding:3px 10px 3px 0;text-align:left;vertical-align:top}
code{font:11.5px ui-monospace,monospace}
</style></head><body>
<h1>Кадры макета и пруф: построчное сравнение</h1>
<div>Каждый кадр и его экран пруфа обходятся одним деревом на пяти ширинах кадра. Считаются координаты, размеры и вычисленные стили каждого узла и каждого куска текста. Отдельно — расхождения только в токене <code>--dusty</code>.</div>
<table id="sum"></table>
<div id="where"></div>
<div class="pair">
<figure><figcaption>Макет</figcaption><div class="clip" id="cm"><iframe id="mock" src="/mockup" width="1500" height="3000"></iframe></div></figure>
<figure><figcaption>Пруф</figcaption><div class="clip" id="cp"><iframe id="proof" width="1500" height="3000"></iframe></div></figure>
</div>
<table class="d" id="diffs"></table>
<script>
const CFG=/*CONFIG*/;
const WIDTHS=[1351,1100,960,820,700];
const IGNORE=new Set(CFG.ignore);
const TOKENS=[["rgb(106, 99, 81)","rgb(111, 104, 86)"]];
const PROPS=["display","fontSize","fontWeight","fontFamily","fontStyle","lineHeight","letterSpacing","textTransform","color",
 "backgroundColor","backgroundImage","borderTopWidth","borderTopColor","borderRightWidth","borderBottomWidth","borderLeftWidth",
 "borderTopLeftRadius","opacity","textDecorationLine","paddingTop","paddingLeft","paddingRight","paddingBottom","whiteSpace","overflow"];
// Layout (container queries included) is computed synchronously when measured,
// and fonts are awaited separately. No timers: a background tab throttles them.
const tick=()=>Promise.resolve();
const loaded=(f,src)=>new Promise(ok=>{f.addEventListener("load",()=>ok(),{once:true}); if(src) f.src=src;});
const first=f=>f.split(",")[0].trim().replace(/"/g,"");
const mapClass=c=>c==="mock"?"tcx":c==="is-dark"?"":c.startsWith("m-")?"tcx-"+c.slice(2):c; // is-dark is the theme attribute
const isCtl=n=>n.nodeType===1&&/^(INPUT|TEXTAREA)$/.test(n.tagName);
function items(el){const o=[];for(const n of el.childNodes){if(n.nodeType===1)o.push(n);else if(n.nodeType===3&&n.textContent.trim())o.push(n);}return o;}
function box(n){if(n.nodeType===3){const r=n.ownerDocument.createRange();r.selectNodeContents(n);return r.getBoundingClientRect();}return n.getBoundingClientRect();}
const norm=s=>s.replace(/\s+/g," ").trim();
function isToken(a,b){return TOKENS.some(([x,y])=>String(a).split(x).join(y)===String(b));}
function add(out,path,k,a,b){out.push({path,k,a,b,token:isToken(a,b)});}
function cmpBox(out,path,a,b,ra,rb,tol){
  const x=box(a),y=box(b);
  if(!x.width&&!x.height&&!y.width&&!y.height)return; // hidden on both sides; display is compared
  for(const [k,va,vb] of [["x",x.left-ra.left,y.left-rb.left],["y",x.top-ra.top,y.top-rb.top],["w",x.width,y.width],["h",x.height,y.height]])
    if(Math.abs(va-vb)>tol) add(out,path,k,va.toFixed(1),vb.toFixed(1));
}
function label(n,i){if(n.nodeType===3)return "#text["+i+"]";const c=typeof n.className==="string"&&n.className?"."+n.className.trim().split(/\s+/).join("."):n.tagName.toLowerCase();return c+"["+i+"]";}
function walk(a,b,ra,rb,path,out,st,skip){
  st.nodes++;
  // The frame puts an open menu inside .m-btn; a real menu is not content of a
  // button, so the widget prints .tcx-drop: button + menu. Compared part by part.
  if(a.nodeType===1&&b.nodeType===1&&b.classList&&b.classList.contains("tcx-drop")){
    const am=a.querySelector(":scope > .m-menu");
    if(am&&b.children.length===2){
      out.push({path,k:"structure",a:"menu inside .m-btn",b:".tcx-drop: button + menu",token:false,dev:true});
      walk(a,b.children[0],ra,rb,path,out,st,am);
      walk(am,b.children[1],ra,rb,path+" > .m-menu",out,st);
      return;
    }
  }
  if(a.nodeType===3&&b.nodeType===3){cmpBox(out,path,a,b,ra,rb,1);if(norm(a.textContent)!==norm(b.textContent))add(out,path,"text",norm(a.textContent),norm(b.textContent));return;}
  if(a.nodeType===3&&isCtl(b)){
    const t=box(a),r=b.getBoundingClientRect(),cs=getComputedStyle(b);
    const bx=r.left+parseFloat(cs.borderLeftWidth)+parseFloat(cs.paddingLeft)-rb.left, ax=t.left-ra.left;
    const bc=r.top+r.height/2-rb.top, ac=t.top+t.height/2-ra.top;
    if(Math.abs(ax-bx)>1)add(out,path,"text x",ax.toFixed(1),bx.toFixed(1));
    if(Math.abs(ac-bc)>1)add(out,path,"text centre",ac.toFixed(1),bc.toFixed(1));
    if(norm(a.textContent)!==norm(b.value||b.placeholder))add(out,path,"text",norm(a.textContent),norm(b.value||b.placeholder));
    return;}
  if(a.nodeType!==b.nodeType){add(out,path,"node",a.nodeName,b.nodeName);return;}
  const ca=[...a.classList].map(mapClass).filter(Boolean).sort().join(" "),cb=[...b.classList].filter(c=>!IGNORE.has(c)).sort().join(" ");
  if(ca!==cb)add(out,path,"class",ca,cb);
  cmpBox(out,path,a,b,ra,rb,.5);
  const sa=getComputedStyle(a),sb=getComputedStyle(b);
  for(const p of PROPS){if(p==="overflow"&&isCtl(b))continue; // the UA fixes overflow on controls
    let va=sa[p],vb=sb[p];if(p==="fontFamily"){va=first(va);vb=first(vb);}if(va!==vb)add(out,path,p,va,vb);}
  if(isCtl(b)){if(norm(a.textContent)!==norm(b.value||b.placeholder))add(out,path,"text",norm(a.textContent),norm(b.value||b.placeholder));return;}
  if(a.tagName.toLowerCase()==="svg")return;
  const ka=items(a).filter(n=>n!==skip),kb=items(b);
  if(ka.length!==kb.length){add(out,path,"items",ka.length,kb.length);return;}
  ka.forEach((k,i)=>walk(k,kb[i],ra,rb,path+" > "+label(k,i),out,st));
}
function isolate(doc,labelText){
  const frames=[...doc.querySelectorAll(".frame")];
  const hit=frames.find(f=>{const l=f.querySelector(":scope > .frame-label > span");return l&&l.textContent===labelText;});
  doc.querySelectorAll(".wrap > *").forEach(e=>{e.style.display=e.contains(hit)?"":"none";});
  frames.forEach(f=>{f.style.display=(f===hit||f.contains(hit)||hit.contains(f))?"":"none";});
  return hit.querySelector(".mock");
}
const RES={};
let mockRoot=null,proofRoot=null;
async function setWidth(W){mockRoot.style.width=W+"px";mockRoot.style.flex="none";proofRoot.style.width=W+"px";await tick(fm.contentWindow);await tick(fp.contentWindow);}
function place(W){
  const mw=fm.contentWindow;mw.scrollTo(0,0);
  const r=mockRoot.getBoundingClientRect();mw.scrollTo(0,r.top+mw.scrollY);
  const r2=mockRoot.getBoundingClientRect();
  fm.style.left=(-r2.left*.5)+"px";fm.style.top=(-r2.top*.5)+"px";
  const h=Math.max(r2.height,proofRoot.getBoundingClientRect().height);
  for(const c of [cm,cp]){c.style.width=(W*.5)+"px";c.style.height=(h*.5)+"px";}
}
const fm=document.getElementById("mock"),fp=document.getElementById("proof"),cm=document.getElementById("cm"),cp=document.getElementById("cp");
async function run(){
  if(!fm.contentDocument||fm.contentDocument.readyState!=="complete"||!fm.contentDocument.querySelector(".frame"))await loaded(fm);
  await fm.contentDocument.fonts.ready;
  for(const c of CFG.cases){
    await loaded(fp,c.path);await fp.contentDocument.fonts.ready;
    mockRoot=isolate(fm.contentDocument,c.label);proofRoot=fp.contentDocument.querySelector(".tcx");
    for(const W of WIDTHS){
      await setWidth(W);
      const out=[],st={nodes:0};
      walk(mockRoot,proofRoot,mockRoot.getBoundingClientRect(),proofRoot.getBoundingClientRect(),".frame",out,st);
      RES[c.name+"@"+W]={nodes:st.nodes,diffs:out};
    }
  }
  summary();await show(CFG.cases[0].name,WIDTHS[0]);
  window.__proof=RES;
}
function summary(){
  let h="<tr><th>Кадр</th>"+WIDTHS.map(w=>"<th>"+w+" px</th>").join("")+"</tr>";
  for(const c of CFG.cases){h+="<tr><th>"+c.name+"</th>";for(const W of WIDTHS){const r=RES[c.name+"@"+W];
    const real=r.diffs.filter(d=>!d.token&&!d.dev).length,tok=r.diffs.filter(d=>d.token).length,dev=r.diffs.filter(d=>d.dev).length;
    h+="<td class='"+(real?"bad":"ok")+"' onclick='show(\""+c.name+"\","+W+")'>"+r.nodes+" узлов<br>расхождений: <b>"+real+"</b><br><small>--dusty: "+tok+(dev?" · структура: "+dev:"")+"</small></td>";}h+="</tr>";}
  document.getElementById("sum").innerHTML=h;
}
async function show(name,W){
  const c=CFG.cases.find(x=>x.name===name);
  if(!fp.contentDocument||!fp.contentDocument.location.pathname.endsWith(c.path)){await loaded(fp,c.path);await fp.contentDocument.fonts.ready;}
  mockRoot=isolate(fm.contentDocument,c.label);proofRoot=fp.contentDocument.querySelector(".tcx");
  await setWidth(W);place(W);
  const r=RES[name+"@"+W];const real=r.diffs.filter(d=>!d.token);
  document.getElementById("where").textContent="Показан: "+name+" при ширине кадра "+W+" px. Расхождений без токена: "+real.length+".";
  document.getElementById("diffs").innerHTML="<tr><th>Узел</th><th>Свойство</th><th>Макет</th><th>Пруф</th></tr>"+
    real.slice(0,400).map(d=>"<tr><td><code>"+d.path+"</code></td><td>"+d.k+"</td><td><code>"+d.a+"</code></td><td><code>"+d.b+"</code></td></tr>").join("");
}
run();
</script></body></html>`
