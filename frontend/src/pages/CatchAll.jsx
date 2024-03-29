import { Player } from "@lottiefiles/react-lottie-player";
import ErrorJSON from '../assets/Lottie/404.json'
export default function NotFound() {
    return(
        <div className="w-screen h-screen flex flex-col items-center justify-center">
             <Player
            src={ErrorJSON}
            loop
            autoplay
            style={{ height: "300px", width: "300px" }}
          />
        </div>
    )
}
