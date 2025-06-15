import {useState, useEffect} from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import {StartWorkSession, StartBreakSession, PauseSession, ResumeSession, GetCurrentState, GetRemainingTime} from "../wailsjs/go/main/App";
import {EventsOn} from "../wailsjs/runtime/runtime";

function App() {
    const [currentState, setCurrentState] = useState("Idle");
    const [remainingTime, setRemainingTime] = useState(0);
    const [isPaused, setIsPaused] = useState(false);

    useEffect(() => {
        // Listen for session events
        EventsOn("session:start", (data: any) => {
            setCurrentState(data.state);
        });

        EventsOn("session:end", (data: any) => {
            setCurrentState("Idle");
        });

        EventsOn("session:pause", (data: any) => {
            setIsPaused(true);
        });

        EventsOn("session:resume", (data: any) => {
            setIsPaused(false);
        });

        EventsOn("timer:tick", (data: any) => {
            setRemainingTime(data.remainingTime);
        });

        EventsOn("warning", (data: any) => {
            console.log("Warning: Idle for", data.idleDuration, "minutes");
        });

        // Update state periodically
        const interval = setInterval(async () => {
            const state = await GetCurrentState();
            const time = await GetRemainingTime();
            setCurrentState(state);
            setRemainingTime(time);
        }, 1000);

        return () => clearInterval(interval);
    }, []);

    const formatTime = (seconds: number) => {
        const minutes = Math.floor(seconds / 60);
        const secs = seconds % 60;
        return `${minutes.toString().padStart(2, '0')}:${secs.toString().padStart(2, '0')}`;
    };

    const handleStartWork = async () => {
        await StartWorkSession();
    };

    const handleStartBreak = async () => {
        await StartBreakSession();
    };

    const handlePause = async () => {
        await PauseSession();
    };

    const handleResume = async () => {
        await ResumeSession();
    };

    return (
        <div id="App">
            <img src={logo} id="logo" alt="logo"/>
            <div id="result" className="result">
                <h1>Karedoro ポモドーロタイマー</h1>
                <p>現在の状態: {currentState}</p>
                <p>残り時間: {formatTime(remainingTime)}</p>
            </div>
            <div id="input" className="input-box">
                {currentState === "Idle" && (
                    <>
                        <button className="btn" onClick={handleStartWork}>作業セッション開始</button>
                        <button className="btn" onClick={handleStartBreak}>休憩セッション開始</button>
                    </>
                )}
                {currentState !== "Idle" && (
                    <>
                        {isPaused ? (
                            <button className="btn" onClick={handleResume}>再開</button>
                        ) : (
                            <button className="btn" onClick={handlePause}>一時停止</button>
                        )}
                    </>
                )}
            </div>
        </div>
    )
}

export default App
