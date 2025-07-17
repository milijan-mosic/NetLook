import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import Homepage from "./pages/Homepage.tsx";
import { BrowserRouter, Routes, Route } from "react-router";
import "./index.css";

const root: HTMLElement | null = document.getElementById("root");

createRoot(root!).render(
  <StrictMode>
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Homepage />} />
      </Routes>
    </BrowserRouter>
  </StrictMode>,
);
