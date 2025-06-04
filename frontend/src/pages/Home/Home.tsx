import React from "react";
import { Greet } from "@/components/ui/Greet";

const Home: React.FC = () => {
  return (
    <div>
      <div>
        <header>
          <h1>Connect gRPC sample app</h1>
          <p>Welcome to our Connect gRPC sample app</p>
        </header>

        <main>
          <Greet />
        </main>

        <footer>
          <p>Powered by Rintarooo</p>
        </footer>
      </div>
    </div>
  );
};

export default Home;
