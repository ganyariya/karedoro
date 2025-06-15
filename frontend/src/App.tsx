import {useState, useEffect} from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import {StartWorkSession, StartBreakSession, PauseSession, ResumeSession, GetCurrentState, GetRemainingTime} from "../wailsjs/go/main/App";
import {EventsOn} from "../wailsjs/runtime/runtime";
import { APP_CONSTANTS, BUTTON_LABELS } from './constants/app';
import { EVENTS } from './constants/events';
import { SESSION_STATE, SessionStateType, isIdleState, isWorkSession, isBreakSession } from './types/session';

function App() {
    const [currentState, setCurrentState] = useState<SessionStateType>(SESSION_STATE.IDLE);
    const [remainingTime, setRemainingTime] = useState(0);
    const [isPaused, setIsPaused] = useState(false);

    useEffect(() => {
        // Listen for session events
        EventsOn(EVENTS.SESSION_START, (data: any) => {
            setCurrentState(data.state as SessionStateType);
        });

        EventsOn(EVENTS.SESSION_END, (data: any) => {
            setCurrentState(SESSION_STATE.IDLE);
        });

        EventsOn(EVENTS.SESSION_PAUSE, (data: any) => {
            setIsPaused(true);
        });

        EventsOn(EVENTS.SESSION_RESUME, (data: any) => {
            setIsPaused(false);
        });

        EventsOn(EVENTS.TIMER_TICK, (data: any) => {
            setRemainingTime(data.remainingTime);
        });

        EventsOn(EVENTS.WARNING, (data: any) => {
            console.log("Warning: Idle for", data.idleDuration, "minutes");
        });

        // Update state periodically
        const interval = setInterval(async () => {
            const state = await GetCurrentState();
            const time = await GetRemainingTime();
            setCurrentState(state as SessionStateType);
            setRemainingTime(time);
        }, APP_CONSTANTS.TIMER_UPDATE_INTERVAL);

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
                <h1>{APP_CONSTANTS.TITLE}</h1>
                <p>現在の状態: {currentState}</p>
                <p>残り時間: {formatTime(remainingTime)}</p>
            </div>
            <div id="input" className="input-box">
                {isIdleState(currentState) && (
                    <>
                        <button className="btn" onClick={handleStartWork}>{BUTTON_LABELS.START_WORK}</button>
                        <button className="btn" onClick={handleStartBreak}>{BUTTON_LABELS.START_BREAK}</button>
                    </>
                )}
                {!isIdleState(currentState) && (
                    <>
                        {isPaused ? (
                            <button className="btn" onClick={handleResume}>{BUTTON_LABELS.RESUME}</button>
                        ) : (
                            <button className="btn" onClick={handlePause}>{BUTTON_LABELS.PAUSE}</button>
                        )}
                    </>
                )}
            </div>
        </div>
    )
}

export default App
