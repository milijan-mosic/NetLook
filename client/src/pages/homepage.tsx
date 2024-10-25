import { createSignal, onMount } from "solid-js";
import axios from "axios";
import SimpleChart from "../components/chart";

export const Homepage = () => {
  const [data, setData] = createSignal<Array<Object>>([]);
  const [loading, setLoading] = createSignal<boolean>(true);
  const [error, setError] = createSignal<string>("");

  onMount(async () => {
    try {
      const response = await axios.get("http://localhost:11000/");
      setData(response.data);
      console.log(response.data.data.agents);
    } catch (err) {
      setError("Error fetching data");
    } finally {
      setLoading(false);
    }
  });

  return (
    <div class="flex flex-col">
      <div class="flex flex-col justify-center items-center">
        <h1 class="text-3xl font-bold">Hello world!</h1>
      </div>

      <SimpleChart />
    </div>
  );
};
