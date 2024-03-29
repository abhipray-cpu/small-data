import { useNavigate } from "react-router-dom";
import { isAuthenticated } from "../util/Authentication";
import CitizenForm from "../components/CitizenForm";
import { useEffect } from "react";
export default function AddPage() {
  const navigate = useNavigate();
  useEffect(() => {
    if (!isAuthenticated()) {
      navigate("/login");
    }
  }, [navigate]);
  return (
    <div className="w-screen h-min-screen flex flex-col items-center pt-4">
      <h1 className="font-serif font-medium text-2xl text-gray-700 mb-10 ">
        Add Citizen
      </h1>
      <CitizenForm method="post"></CitizenForm>
    </div>
  );
}
