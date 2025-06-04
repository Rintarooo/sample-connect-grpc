import React from "react";

const Greet: React.FC = () => {
  return (
    <div>
      <h2>Greeting</h2>
      <input type="text" id="name" />
      <button id="greet">Greet</button>
    </div>
  );
};

export default Greet;
