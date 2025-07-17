import axios from "axios";
import { useEffect } from "react";

function Homepage() {
  const url: string = "/api/1.0/metrics";

  const fetchMetrics = () => {
    axios
      .get(url)
      .then(function (response) {
        console.log(response?.data?.data?.agents);
      })
      .catch(function (error) {
        console.log(error);
      });
  };

  useEffect(() => {
    fetchMetrics();
  }, []);

  return (
    <div className="flex flex-col">
      <div className="flex flex-col justify-center items-center">
        <h1 className="text-3xl font-bold">Hello world!</h1>
      </div>
    </div>
  );
}

export default Homepage;
