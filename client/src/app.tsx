import { createSignal } from "solid-js";
import { Homepage } from "./pages/homepage";

function App() {
  const [count, setCount] = createSignal(0);

  return <Homepage />;
}

export default App;
