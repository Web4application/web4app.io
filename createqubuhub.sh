#!/bin/bash

# QubuHub Web4 Jekyll Starter Generator
echo "🚀 Creating QubuHub Web4 starter structure..."

# Create folder structure
mkdir -p qubuhub-web4/_layouts
mkdir -p qubuhub-web4/_data
mkdir -p qubuhub-web4/_posts

cd qubuhub-web4

# 1️⃣ _layouts/index.html
cat > _layouts/index.html << 'EOF'
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>{{ site.title }}</title>
  <style>
    h1, h2 { text-shadow: 0 0 5px #00d1ff, 0 0 10px #00d1ff, 0 0 20px #00d1ff; }
    a { transition: all 0.3s ease; text-decoration: none; }
    a:hover { color: #00ffe0; text-shadow: 0 0 5px #00ffe0, 0 0 10px #00ffe0, 0 0 20px #00ffe0; }
    .project-card { background-color: #1a1c29; border:1px solid #00d1ff; border-radius:12px; padding:20px; margin:15px 0; transition: transform 0.3s ease, box-shadow 0.3s ease; }
    .project-card:hover { transform: translateY(-5px); box-shadow: 0 0 20px #00d1ff; }
    .project-title { color: #00d1ff; font-size: 1.3em; margin-bottom:10px; }
    .project-desc { color: #e0e0e0; }
    .container { max-width:900px; margin:0 auto; padding:20px; }
    #bg-canvas { position:fixed; top:0; left:0; width:100%; height:100%; z-index:-1; background-color:#0d0e18; }
    body { margin:0; font-family:sans-serif; color:#e0e0e0; background-color:transparent; }
  </style>
</head>
<body>
<canvas id="bg-canvas"></canvas>

<header class="container" style="text-align:center; padding:40px 20px;">
  <h1>{{ site.title }}</h1>
  <p style="color:#ccc; font-size:1.2em;">AI & Blockchain-Powered Web4 Projects by {{ site.author }}</p>
</header>

<section class="container">
  <h2>🌐 Featured Projects</h2>
  {% for project in site.data.projects %}
    <div class="project-card">
      <div class="project-title">{{ project.name }}</div>
      <div class="project-desc">{{ project.description }}</div>
    </div>
  {% endfor %}
</section>

<section class="container">
  <h2>📰 Latest Blog Posts</h2>
  <ul style="list-style:none; padding-left:0;">
    {% for post in site.posts limit:5 %}
      <li style="margin-bottom:15px;">
        <a href="{{ post.url | relative_url }}">{{ post.title }}</a>
        <span style="color:#888; font-size:0.9em;"> — {{ post.date | date: "%B %-d, %Y" }}</span>
      </li>
    {% endfor %}
  </ul>
</section>

<section class="container">
  <h2>✨ About the Author</h2>
  <p style="color:#e0e0e0; font-size:1.1em;">
    {{ site.author }} is the founder and visionary behind QubuHub. A developer and innovator specializing in AI, blockchain, and Web4, creating tools that empower users and shape the decentralized web.
  </p>
</section>

<footer style="text-align:center; padding:20px; color:#888;">
  &copy; {{ site.time | date: "%Y" }} {{ site.title }}. All Rights Reserved.
</footer>

<script>
const canvas=document.getElementById('bg-canvas'),ctx=canvas.getContext('2d');
const baseColors=[{r:0,g:209,b:255},{r:0,g:255,b:224},{r:0,g:170,b:255}];
let hueShift=0,particles=[],trails=[],maxTrails=30;
const particleCount=window.innerWidth<768?60:120;
let mouse={x:null,y:null},scrollOffset=0;

window.addEventListener('mousemove',e=>{mouse.x=e.clientX;mouse.y=e.clientY;trails.push({x:mouse.x,y:mouse.y});if(trails.length>maxTrails)trails.shift();});
window.addEventListener('mouseleave',()=>{mouse.x=null;mouse.y=null;});
window.addEventListener('scroll',()=>{scrollOffset=window.scrollY*0.2;});
window.addEventListener('resize',()=>{canvas.width=window.innerWidth; canvas.height=window.innerHeight;});
canvas.width=window.innerWidth; canvas.height=window.innerHeight;

class Particle{constructor(){this.reset();}reset(){this.x=Math.random()*canvas.width;this.y=Math.random()*canvas.height;this.radius=Math.random()*2+1;this.speedX=(Math.random()-0.5)*0.3;this.speedY=(Math.random()-0.5)*0.3;this.baseColor=baseColors[Math.floor(Math.random()*baseColors.length)];}update(){this.x+=this.speedX;this.y+=this.speedY;if(this.x<0||this.x>canvas.width||this.y<0||this.y>canvas.height)this.reset();if(mouse.x&&mouse.y){let dx=this.x-mouse.x,dy=this.y-mouse.y,dist=Math.sqrt(dx*dx+dy*dy);if(dist<100){let angle=Math.atan2(dy,dx),force=(100-dist)*0.02;this.x+=Math.cos(angle)*force;this.y+=Math.sin(angle)*force;}}}draw(){const r=Math.min(255,Math.max(0,this.baseColor.r+Math.sin(hueShift)*50));const g=Math.min(255,Math.max(0,this.baseColor.g+Math.cos(hueShift)*50));const b=Math.min(255,Math.max(0,this.baseColor.b+Math.sin(hueShift/2)*50));const color=`rgb(${r},${g},${b})`;ctx.beginPath();ctx.arc(this.x,this.y,this.radius,0,Math.PI*2);ctx.fillStyle=color;ctx.shadowBlur=8;ctx.shadowColor=color;ctx.fill();}}
for(let i=0;i<particleCount;i++)particles.push(new Particle());
function drawGrid(){const spacing=100;ctx.strokeStyle='rgba(0,209,255,0.05)';ctx.lineWidth=1;for(let x=0;x<canvas.width;x+=spacing){ctx.beginPath();ctx.moveTo(x,0);ctx.lineTo(x,canvas.height);ctx.stroke();}for(let y=0;y<canvas.height;y+=spacing){ctx.beginPath();ctx.moveTo(0,y);ctx.lineTo(canvas.width,y);ctx.stroke();}for(let x=0;x<canvas.width;x+=spacing){for(let y=0;y<canvas.height;y+=spacing){ctx.beginPath();ctx.arc(x,y,1.2,0,Math.PI*2);ctx.fillStyle='rgba(0,209,255,0.1)';ctx.shadowBlur=6;ctx.shadowColor='rgba(0,209,255,0.2)';ctx.fill();}}}
function drawLines(){const connectionRadius=100;for(let i=0;i<particles.length;i++){let p1=particles[i];for(let j=i+1;j<i+10&&j<particles.length;j++){let p2=particles[j];let dx=p1.x-p2.x,dy=p1.y-p2.y,dist=Math.sqrt(dx*dx+dy*dy);if(dist<connectionRadius){let alpha=1-dist/connectionRadius;if(mouse.x&&mouse.y){let midX=(p1.x+p2.x)/2,midY=(p1.y+p2.y)/2,mdx=midX-mouse.x,mdy=midY-mouse.y,mdist=Math.sqrt(mdx*mdx+mdy*mdy);if(mdist<150)alpha+=0.3;if(alpha>1)alpha=1;}const r=Math.min(255,Math.max(0,0+Math.sin(hueShift)*50));const g=Math.min(255,Math.max(0,209+Math.cos(hueShift)*50));const b=Math.min(255,Math.max(0,255+Math.sin(hueShift/2)*50));ctx.beginPath();ctx.moveTo(p1.x,p1.y);ctx.lineTo(p2.x,p2.y);ctx.strokeStyle=`rgba(${r},${g},${b},${alpha})`;ctx.lineWidth=0.5;ctx.shadowBlur=4;ctx.shadowColor=`rgb(${r},${g},${b})`;ctx.stroke();}}}}
function drawTrails(){for(let i=0;i<trails.length;i++){const point=trails[i];const alpha=i/trails.length*0.5;ctx.beginPath();ctx.arc(point.x,point.y,2,0,Math.PI*2);ctx.fillStyle=`rgba(0,255,224,${alpha})`;ctx.shadowBlur=6;ctx.shadowColor=`rgba(0,255,224,${alpha})`;ctx.fill();}}
function animate(){ctx.clearRect(0,0,canvas.width,canvas.height);ctx.save();ctx.translate(0,scrollOffset);drawGrid();drawTrails();particles.forEach(p=>{p.update();p.draw();});drawLines();ctx.restore();hueShift+=0.002;requestAnimationFrame(animate);}
animate();
document.addEventListener('visibilitychange',()=>{if(!document.hidden)animate();});
</script>
</body>
</html>
EOF

# 2️⃣ _data/projects.yml
cat > _data/projects.yml << 'EOF'
- name: Lola
  description: Personal AI companion for voice, journaling, and productivity.
- name: RODA / RODAAI
  description: AI platform for analytics, automation, and intelligent insights.
- name: AgbakoAI
  description: Modular AI framework providing ML, NLP, and automation capabilities.
- name: ai-webapp
  description: Web applications integrating AI for personalized automation.
- name: ChatGPT5 Mini
  description: Lightweight conversational AI for web and mobile platforms.
- name: ProjectPilot
  description: AI system for project planning, analysis, and optimization.
- name: Congen
  description: Generative AI
