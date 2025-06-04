import React, { useState } from "react";
import { getGreeting } from "@/api/greetApi";

const Greet: React.FC = () => {
  const [name, setName] = useState("");
  const [greeting, setGreeting] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  const handleGreet = async () => {
    if (!name.trim()) {
      setError("Please enter a name");
      return;
    }

    setLoading(true);
    setError("");
    setGreeting("");

    try {
      const response = await getGreeting(name);
      if (response.greeting) {
        setGreeting(response.greeting.message);
      }
    } catch (err) {
      setError(err instanceof Error ? err.message : "Failed to get greeting");
    } finally {
      setLoading(false);
    }
  };

  const handleKeyPress = (e: React.KeyboardEvent) => {
    if (e.key === "Enter") {
      handleGreet();
    }
  };

  return (
    <div>
      <h2>Greeting</h2>
      <div style={{ marginBottom: "10px" }}>
        <input
          type="text"
          id="name"
          value={name}
          onChange={(e) => setName(e.target.value)}
          onKeyDown={handleKeyPress}
          placeholder="Enter your name"
          disabled={loading}
          style={{ marginRight: "10px", padding: "5px" }}
        />
        <button
          id="greet"
          onClick={handleGreet}
          disabled={loading || !name.trim()}
          style={{ padding: "5px 10px" }}
        >
          {loading ? "Loading..." : "Greet"}
        </button>
      </div>

      {error && (
        <div style={{ color: "red", marginBottom: "10px" }}>Error: {error}</div>
      )}

      {greeting && (
        <div style={{ color: "green", fontWeight: "bold" }}>{greeting}</div>
      )}
    </div>
  );
};

export default Greet;
