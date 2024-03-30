import { useNavigate, redirect, useLoaderData } from "react-router-dom";
import { isAuthenticated } from "../util/Authentication";
import { getToken } from "../util/Authentication";
import { useState, useEffect } from "react";
import { Player } from "@lottiefiles/react-lottie-player";
import LoaderJSON from "../assets/Lottie/loading.json";
import ErrorJSON from "../assets/Lottie/error.json";
import CitizenForm from "../components/CitizenForm";
export default function AddPage() {
  const navigate = useNavigate();
  const loaderData = useLoaderData();
  const [loading, setLoading] = useState(1);
  const [userData, setUserData] = useState([]);
  useEffect(() => {
    if (loaderData && loaderData.fail === false) {
      setUserData(loaderData.data);
      setLoading(0);
    }
    if (loaderData && loaderData.fail === true) {
      setLoading(2);
    }
  }, [loaderData, setLoading, setUserData]);

  if (!isAuthenticated()) {
    navigate("/login");
  }
  return (
    <div className="w-screen min-h-screen pt-5 flex flex-col items-center">
      {loading === 1 && <Player src={LoaderJSON} loop autoplay />}
      {loading === 2 && (
        <section className="w-screen flex flex-col items-center">
          <Player
            src={ErrorJSON}
            loop
            autoplay
            speed={3}
            style={{ height: "200px", width: "200px" }}
          />
          <h3 className="text-xl font-serif text-gray-500 mt-8">
            Failed to load data
          </h3>
        </section>
      )}
      {loading === 0 && (
        <section className="w-screen flex flex-col items-center">
          <h2 className="text-2xl font-mono font-medium text-orange-600 text-center mb-6">
            Update Citizen
          </h2>
          <CitizenForm method="put" data={userData}></CitizenForm>
        </section>
      )}
    </div>
  );
}

export async function loader({ params }) {
  try {
    const token = getToken();
    if (!token || token === "EXPIRED") {
      redirect("/login");
    }
    const id = params.id;
    const response = await fetch(`http://localhost:8080/citizen/${id}`, {
      method: "GET",
      headers: {
        Authorization: token,
        "Content-Type": "application/json",
      },
    });
    if (!response.ok) {
      return { data: "", error: "Failed to load data", fail: true };
    }
    const resData = await response.json();
    return { data: resData.data, error: "", fail: false };
  } catch (err) {
    return { data: "", err: "failed to load data", fail: true };
  }
}
