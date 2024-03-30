import { redirect } from "react-router-dom";
import { useState, useEffect } from "react";
import { Form, useNavigation, useActionData } from "react-router-dom";
import { Player } from "@lottiefiles/react-lottie-player";
import LoaderJSON from "../assets/Lottie/loading.json";
import ErrorJSON from "../assets/Lottie/error.json";
import SuccessJSON from "../assets/Lottie/success.json";
import { getToken } from "../util/Authentication";
export default function CitizenForm({ method, data }) {
  const indianStates = [
    "Andhra Pradesh",
    "Arunachal Pradesh",
    "Assam",
    "Bihar",
    "Chhattisgarh",
    "Goa",
    "Gujarat",
    "Haryana",
    "Himachal Pradesh",
    "Jharkhand",
    "Karnataka",
    "Kerala",
    "Madhya Pradesh",
    "Maharashtra",
    "Manipur",
    "Meghalaya",
    "Mizoram",
    "Nagaland",
    "Odisha",
    "Punjab",
    "Rajasthan",
    "Sikkim",
    "Tamil Nadu",
    "Telangana",
    "Tripura",
    "Uttar Pradesh",
    "Uttarakhand",
    "West Bengal",
    "Andaman and Nicobar Islands",
    "Chandigarh",
    "Dadra and Nagar Haveli",
    "Daman and Diu",
    "Lakshadweep",
    "Delhi",
    "Puducherry",
  ];
  const actionData = useActionData();
  const navigation = useNavigation();
  const isSubmitting = navigation.state === "submitting";
  useEffect(() => {
    if (isSubmitting) {
      setLoading(1);
    }
    if (actionData && actionData.fail == false) {
      setError("Operation Completed");
      setLoading(3);
      setTimeout(() => {
        setLoading(0);
        setError("");
      }, 1800);
    }
    if (actionData && actionData.fail === true) {
      setError(actionData.error);
      setLoading(2);
      setTimeout(() => {
        setLoading(0);
        setError("");
      }, 1800);
    }
  }, [isSubmitting]);
  const [loading, setLoading] = useState(0);
  const [error, setError] = useState("");
  return (
    <section className="w-screen flex flex-col items-center">
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
      {loading === 3 && (
        <section className="w-screen flex flex-col items-center">
          <Player
            src={SuccessJSON}
            loop
            autoplay
            speed={3}
            style={{ height: "200px", width: "200px" }}
          />
          <h3 className="text-xl font-serif text-gray-500 mt-8">{error}</h3>
        </section>
      )}
      {loading === 0 && (
        <Form method={method} className="w-screen flex flex-col items-center">
          <input
            type="text"
            name="first"
            id="first"
            placeholder="First Name"
            defaultValue={data ? data.FirstName : ""}
            className="w-[90vw] h-14 border border-gray-700 rounded-lg px-3 py-2 text-lg font-mono tracking-normal text-gray-700  focus:outline-none focus:border-gray900 mb-5"
          />
          <input
            type="text"
            name="last"
            id="last"
            placeholder="Last Name"
            defaultValue={data ? data.LastName : ""}
            className="w-[90vw] h-14 border border-gray-700 rounded-lg px-3 py-2 text-lg font-mono tracking-normal text-gray-700  focus:outline-none focus:border-gray900 mb-5"
          />
          <input
            type="date"
            name="dob"
            id="dob"
            placeholder="DOB"
            defaultValue={data ? data.DOB : ""}
            className="w-[90vw] h-14 border border-gray-700 rounded-lg px-3 py-2 text-lg font-mono tracking-normal text-gray-700  focus:outline-none focus:border-gray900 mb-5"
          />
          <select
            name="gender"
            id="gender"
            className="w-[90vw] h-14 border border-gray-700 rounded-lg px-3 py-2 focus:outline-none focus:border-gray900 mb-5"
          >
            <option
              value=""
              disabled
              selected={!data ? true : false}
              className="text-lg font-mono tracking-normal text-gray-700"
            >
              Select an option
            </option>
            <option
              value="Male"
              selected={data && data.Gender === "Male"}
              className="text-lg font-mono tracking-normal text-gray-700"
            >
              Male
            </option>
            <option
              value="Female"
              selected={data && data.Gender === "Female"}
              className="text-lg font-mono tracking-normal text-gray-700"
            >
              Female
            </option>
          </select>
          <input
            type="text"
            name="address"
            id="address"
            placeholder="Address"
            defaultValue={data ? data.Address : ""}
            className="w-[90vw] h-14 border border-gray-700 rounded-lg px-3 py-2 text-lg font-mono tracking-normal text-gray-700  focus:outline-none focus:border-gray900 mb-5"
          />
          <input
            type="text"
            name="city"
            id="city"
            placeholder="City"
            defaultValue={data ? data.City : ""}
            className="w-[90vw] h-14 border border-gray-700 rounded-lg px-3 py-2 text-lg font-mono tracking-normal text-gray-700  focus:outline-none focus:border-gray900 mb-5"
          />
          <input
            type="text"
            name="pincode"
            id="pincode"
            placeholder="PinCode"
            defaultValue={data ? data.PinCode : ""}
            className="w-[90vw] h-14 border border-gray-700 rounded-lg px-3 py-2 text-lg font-mono tracking-normal text-gray-700  focus:outline-none focus:border-gray900 mb-5"
          />
          <select
            name="state"
            id="state"
            className="w-[90vw] h-14 border border-gray-700 rounded-lg px-3 py-2 focus:outline-none focus:border-gray900 mb-5"
          >
            <option
              value=""
              disabled
              selected={!data ? true : false}
              className="text-lg font-mono tracking-normal text-gray-700"
            >
              Select an option
            </option>
            {indianStates.map((state) => {
              return (
                <option
                  key={state}
                  value={state}
                  selected={data && data.State === state}
                  className="text-lg font-mono tracking-normal text-gray-700"
                >
                  {state}
                </option>
              );
            })}
          </select>
          <button className="w-[60vw] h-14 bg-gray-700 rounded-lg mt-5 text-xl text-white font-serif font-medium">
            {method === "post" ? "Add Citizen" : "Edit Citizen"}
          </button>
        </Form>
      )}
    </section>
  );
}

export async function action({ request, params }) {
  try {
    const method = request.method;
    const userID = params.id;
    const data = await request.formData();
    console.log(data);
    if (
      !data.get("first") ||
      !data.get("last") ||
      !data.get("dob") ||
      !data.get("gender") ||
      !data.get("gender") ||
      !data.get("address") ||
      !data.get("city") ||
      !data.get("pincode") ||
      !data.get("state")
    ) {
      return { data: "", error: "Please fill all fields", fail: true };
    }
    const jsonData = {
      FirstName: data.get("first"),
      LastName: data.get("last"),
      DOB: data.get("dob"),
      Gender: data.get("gender"),
      Address: data.get("address"),
      City: data.get("city"),
      Pincode: data.get("pincode"),
      State: data.get("state"),
    };

    const token = getToken();
    if (!token || token === "EXPIRED") {
      return redirect("/login");
    }
    let url =
      method === "POST"
        ? "http://localhost:8080/createCitizen"
        : `http://localhost:8080/updateCitizen/${userID}`;
    console.log(url);
    const response = await fetch(url, {
      method: method,
      headers: {
        "Content-Type": "application/json",
        Authorization: token,
      },
      body: JSON.stringify(jsonData),
    });
    if (!response.ok) {
      const resData = await response.json();
      return { data: "", err: resData.error, status: "failed" };
    }
    const resData = await response.json();
    return { data: resData, error: "Failed to send request", fail: false };
  } catch (err) {
    return { data: "", error: "Failed to send request", fail: true };
  }
}
