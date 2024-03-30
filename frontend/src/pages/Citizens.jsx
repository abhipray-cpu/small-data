import { isAuthenticated,getToken } from "../util/Authentication";
import { useNavigate,defer, redirect,useLoaderData,Await } from "react-router-dom";
import { useEffect,Suspense } from "react";
import { Player } from "@lottiefiles/react-lottie-player";
import LoaderJSON from "../assets/Lottie/loading.json";
import ErrorJSON from "../assets/Lottie/error.json";
import CitizenTable from "../components/CitizenTable";
export default function CitizensPage() {
  const navigate = useNavigate();
  const {citizens} = useLoaderData()
 // const {events} = useLoaderData()
  useEffect(() => {
    if (!isAuthenticated()) {
      navigate("/login");
    }
  }, [navigate]);
  return (
    <div className="w-screen min-h-screen flex flex-col items-center pt-3">
      <h3 className="text-lg font-mono font-medium tracking-wider text-gray-500">Small Data</h3>
      <section className="w-screen mt-5 flex flex-col items-center">
      <Suspense fallback={<Player src={LoaderJSON} loop autoplay />}>
        <Await resolve={citizens} errorElement={ <section className="w-screen flex flex-col items-center">
          <Player
            src={ErrorJSON}
            loop
            autoplay
            speed={3}
            style={{ height: "200px", width: "200px" }}
          />
          <h3 className="text-xl font-serif text-gray-500 mt-8">Failed to load data</h3>
        </section>}>
          {(resolvedCitizens)=> <CitizenTable citizens={resolvedCitizens}></CitizenTable> }
        </Await>
      </Suspense>
      </section>
    </div>
  );
}


async function getCitizens(){ // this function will fetch first 10 events
    try{
      const token = getToken();
      if(!token || token==='EXPIRED'){
          redirect('/login');
      }
      const response = await fetch('http://localhost:8080/citizens/1',{
        method:"GET",
        headers: {
        "Content-Type": "application/json",
        Authorization: token,
      },
      })
      if(!response.ok){
        const resData = await response.json()
        return {data:'',error:resData.error,fail:true}
      }
      const resData = await response.json()
      return {data:resData.data,error:'',fail:false}
    }
    catch(err){
      return {data:'',error:"Failed to load data",fail:true}
    }
}
export async function loader() {
  return defer({
    citizens:getCitizens(),
  })
}
export async function action() {}
