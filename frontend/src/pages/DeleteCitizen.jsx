import { getToken } from "../util/Authentication";
import { redirect,useLoaderData,useNavigate } from "react-router-dom";
import { Player } from "@lottiefiles/react-lottie-player";
import LoaderJSON from "../assets/Lottie/loading.json";
import ErrorJSON from "../assets/Lottie/error.json";
import SuccessJSON from "../assets/Lottie/success.json";
import {useState,useEffect} from 'react'
export default function DeleteCitizen() {
    const [loading,setLoading] = useState(1)
    const [error,setError] = useState('')
    const [userData,setUserData] = useState({})
    const navigate = useNavigate()
    const loaderData = useLoaderData()
    useEffect(()=>{
        if(loaderData && loaderData.fail === true){
            setError(loaderData.error)
            setLoading(2)
        }
        if(loaderData && loaderData.fail === false){
            setUserData(loaderData.data)
            setLoading(0)
        }
    },[loaderData])

    const deleteUser = async ()=>{
        setLoading(1)
        let token = getToken()
    if(!token || token === 'EXPIRED'){
        navigate('/login')
    }
        const response = await fetch(`http://localhost:8080/deleteCitizen/${userData.ID}`,{
            method:"DELETE",
            headers:{
                'AUTHORIZATION':token,
                "Content-Type": "application/json",
            }

        })
        if(!response.ok){
          const resData = await response.json()
          setError(resData.error)
          setLoading(2)
          setTimeout(()=>{
            setLoading(0)
          },2000)   
        }
        const resData = await response.json()
        setError(`Deleted ${resData.data} record`)
        setLoading(3)
        setTimeout(()=>{
            navigate('/')
        },2000)
    }
    return(
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
      {loading === 0 && <section className="w-screen">
        <h2 className="text-2xl font-mono font-medium text-red-600 text-center mb-6">Citizen data</h2>
        <h3 className="ml-5 text-xl text-center font-serif font-medium text-gray-500 tracking-wide mb-2">First Name:    {userData.FirstName}</h3>
        <h3 className="ml-5 text-xl text-center font-serif font-medium text-gray-500 tracking-wide mb-2">Last Name:     {userData.LastName}</h3>
        <h3 className="ml-5 text-xl text-center font-serif font-medium text-gray-500 tracking-wide mb-2">DOB:           {userData.DOB}</h3>
        <h3 className="ml-5 text-xl text-center font-serif font-medium text-gray-500 tracking-wide mb-2">Gender:        {userData.Gender}</h3>
        <h3 className="ml-5 text-xl text-center font-serif font-medium text-gray-500 tracking-wide mb-2">City:          {userData.City}</h3>
        <h3 className="ml-5 text-xl text-center font-serif font-medium text-gray-500 tracking-wide mb-2">PinCode:       {userData.PinCode}</h3>
        <h3 className="ml-5 text-xl text-center font-serif font-medium text-gray-500 tracking-wide mb-2">State:         {userData.State}</h3>
        <h3 className="ml-5 text-xl text-center font-serif font-medium text-gray-500 tracking-wide mb-2">Address:       {userData.Address}</h3>
        <div className="w-screen flex flex-col items-center justify-center mt-5">
            <div className="bg-red-500 w-60 h-14 flex flex-col items-center justify-center text-xl text-white font-serif rounded-xl"
            onClick={deleteUser}>
                Delete Citizen
            </div>
        </div>
      </section>}
        </div>
    )
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

