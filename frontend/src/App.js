import { useState, useEffect } from "react";

function App() {
    const [events, setEvents] = useState([]);
    const [userId, setUserId] = useState(42);
    const [action, setAction] = useState("");

    const fetchEvents = async () => {
        const res = await fetch(`http://localhost:8080/events?user_id=${userId}&from=2026-01-01&to=2026-12-31`);
        const data = await res.json();
        setEvents(data);
    };

    const createEvent = async () => {
        await fetch("http://localhost:8080/events", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ user_id: userId, action, metadata: { page: "/home" } }),
        });
        fetchEvents();
    };

    useEffect(() => { fetchEvents(); }, []);

    return (
        <div>
            <h1>Events</h1>
            <input placeholder="Action" value={action} onChange={e => setAction(e.target.value)} />
            <button onClick={createEvent}>Create Event</button>
            <ul>
                {events.map(e => (
                    <li key={e.id}>{`User ${e.user_id}: ${e.action} at ${e.created_at}`}</li>
                ))}
            </ul>
        </div>
    );
}

export default App;