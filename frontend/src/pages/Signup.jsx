import { Player } from "@lottiefiles/react-lottie-player";
import { useState, useEffect } from "react";
import { redirect, useFetcher, useNavigate } from "react-router-dom";
import LoaderJSON from "../assets/Lottie/loading.json";
import ErrorJSON from "../assets/Lottie/error.json";
import { isAuthenticated } from "../util/Authentication";
function SignupPage() {
  const fetcher = useFetcher();
  const navigate = useNavigate();
  const { data, state } = fetcher;
  const render = () => {
    navigate("/login");
  };

  useEffect(() => {
    if (isAuthenticated()) {
      navigate("/");
    }
    if (state === "submitting") {
      setLoading(1);
    }
    if (data && data.status === "failed") {
      setError(data.err);
      setLoading(2);
      setTimeout(() => {
        setLoading(0);
      }, 1800);
    }
  }, [data, state, navigate]);
  const [loading, setLoading] = useState(0);
  const [error, setError] = useState("");
  return (
    <div className="w-screen min-h-screen flex flex-col items-center pt-6">
      <h2 className="text-4xl font-normal text-gray-600 tracking-wide mb-8">
        Small Data
      </h2>
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
          <h3 className="text-xl font-serif text-gray-500 mt-8">{error}</h3>
        </section>
      )}
      {loading === 0 && (
        <>
          <fetcher.Form
            method="post"
            className="mt-24 w-screen flex flex-col items-center justify-center gap-5"
          >
            <input
              type="text"
              name="name"
              id="name"
              placeholder="username"
              className="w-[90vw] h-14 border border-gray-700 rounded-lg px-3 py-2 text-lg font-mono tracking-normal text-gray-700  focus:outline-none focus:border-gray900 "
            />
            <input
              type="text"
              name="email"
              id="email"
              placeholder="email"
              className="w-[90vw] h-14 border border-gray-700 rounded-lg px-3 py-2 text-lg font-mono tracking-normal text-gray-700  focus:outline-none focus:border-gray900 "
            />
            <input
              type="password"
              name="password"
              id="password"
              placeholder="password"
              className="w-[90vw] h-14 border border-gray-700 rounded-lg px-3 py-2 text-lg font-mono tracking-normal text-gray-700  focus:outline-none focus:border-gray900 "
            />
            <input
              type="password"
              name="confirm"
              id="confirm"
              placeholder="confirm password"
              className="w-[90vw] h-14 border border-gray-700 rounded-lg px-3 py-2 text-lg font-mono tracking-normal text-gray-700  focus:outline-none focus:border-gray900 "
            />
            <button className="w-[60vw] h-14 bg-gray-700 rounded-lg mt-10 text-xl text-white font-serif font-medium">
              Signup
            </button>
          </fetcher.Form>
          <span className="fixed bottom-10 text-lg font-serif text-gray-400">
            Already have an account?{" "}
            <strong
              className="font-serif text-xl text-gray-600"
              onClick={render}
            >
              Login
            </strong>
          </span>
        </>
      )}
    </div>
  );
}

export default SignupPage;

export async function action({ request }) {
  const data = await request.formData();

  if (
    !data.get("email") ||
    !data.get("name") ||
    !data.get("password") ||
    !data.get("confirm")
  ) {
    return { data: "", err: "Please fill all fields", status: "failed" };
  }
  if (data.get("password") !== data.get("confirm")) {
    return { data: "", err: "Password do not match", status: "failed" };
  }
  if (data.get("password").length < 8) {
    return {
      data: "",
      err: "Password must be min 8 characters",
      status: "failed",
    };
  }
  const userData = {
    Name: data.get("name"),
    EmaiL: data.get("email"),
    Type: "Admin",
    Password: data.get("password"),
  };
  try {
    const response = await fetch("http://localhost:8080/signup", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(userData),
    });
    if (!response.ok) {
      const resData = await response.json();
      return { data: "", err: resData.error, status: "failed" };
    }
    return redirect("/login");
  } catch (err) {
    return { data: "", err: "failed to send request", status: "failed" };
  }
}
