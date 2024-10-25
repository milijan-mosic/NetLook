/* @refresh reload */
import { render } from "solid-js/web";
import "./base.css";
import App from "./app.tsx";

const root = document.getElementById("root");

render(() => <App />, root!);
