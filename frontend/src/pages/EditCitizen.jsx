import { useNavigate } from "react-router-dom";
import { isAuthenticated } from "../util/Authentication";
export default function AddPage() {
  const navigate = useNavigate();
  if (!isAuthenticated()) {
    navigate("/login");
  }
  return (
    <div className="w-screen h-min-screen flex flex-col items-center pt-4">
      <h1 className="font-serif font-medium text-2xl text-gray-700 ">
        Edit Citizen
      </h1>
    </div>
  );
}
