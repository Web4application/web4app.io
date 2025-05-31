import AOS from 'aos';
import 'aos/dist/aos.css';
import Hero from "./components/Hero";
import Features from "./components/Features";
import Solutions from "./components/Solutions";
import About from "./components/About";
import Footer from "./components/Footer";
import { useEffect } from 'react';

useEffect(() => {
  AOS.init({ duration: 1000 });
}, []);
<div data-aos="fade-up">
  {/* Content */}
</div>

export default function App() {
  return (
    <>
      <Hero />
      <Features />
      <Solutions />
      <About />
      <Footer />
    </>
  );
}
