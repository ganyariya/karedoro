import {useState, useEffect} from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import {StartWorkSession, StartBreakSession, PauseSession, ResumeSession, GetCurrentState, GetRemainingTime} from "../wailsjs/go/main/App";
import {EventsOn, EventsOff} from "../wailsjs/runtime/runtime";
import { APP_CONSTANTS, BUTTON_LABELS } from './constants/app';
import { EVENTS } from './constants/events';
import { SESSION_STATE, SessionStateType, isIdleState, isWorkSession, isBreakSession } from './types/session';

function App() {
    const [currentState, setCurrentState] = useState<SessionStateType>(SESSION_STATE.IDLE);
    const [remainingTime, setRemainingTime] = useState(0);
    const [isPaused, setIsPaused] = useState(false);

    useEffect(() => {
        // Event handler functions
        const handleSessionStart = (data: any) => {
            console.log('Session start event:', data);
            setCurrentState(data.state as SessionStateType);
        };

        const handleSessionEnd = (data: any) => {
            console.log('Session end event:', data);
            setCurrentState(SESSION_STATE.IDLE);
        };

        const handleSessionPause = (data: any) => {
            console.log('Session pause event:', data);
            setIsPaused(true);
        };

        const handleSessionResume = (data: any) => {
            console.log('Session resume event:', data);
            setIsPaused(false);
        };

        const handleTimerTick = (data: any) => {
            setRemainingTime(data.remainingTime);
        };

        const handleWarning = (data: any) => {
            console.log("Warning: Idle for", data.idleDuration, "minutes");
        };

        // Register event listeners - EventsOn may return cleanup functions
        const cleanupFunctions: Array<() => void> = [];
        
        try {
            const sessionStartCleanup = EventsOn(EVENTS.SESSION_START, handleSessionStart);
            const sessionEndCleanup = EventsOn(EVENTS.SESSION_END, handleSessionEnd);
            const sessionPauseCleanup = EventsOn(EVENTS.SESSION_PAUSE, handleSessionPause);
            const sessionResumeCleanup = EventsOn(EVENTS.SESSION_RESUME, handleSessionResume);
            const timerTickCleanup = EventsOn(EVENTS.TIMER_TICK, handleTimerTick);
            const warningCleanup = EventsOn(EVENTS.WARNING, handleWarning);
            
            // If EventsOn returns cleanup functions, store them
            if (typeof sessionStartCleanup === 'function') cleanupFunctions.push(sessionStartCleanup);
            if (typeof sessionEndCleanup === 'function') cleanupFunctions.push(sessionEndCleanup);
            if (typeof sessionPauseCleanup === 'function') cleanupFunctions.push(sessionPauseCleanup);
            if (typeof sessionResumeCleanup === 'function') cleanupFunctions.push(sessionResumeCleanup);
            if (typeof timerTickCleanup === 'function') cleanupFunctions.push(timerTickCleanup);
            if (typeof warningCleanup === 'function') cleanupFunctions.push(warningCleanup);
        } catch (error) {
            console.error('Error registering event listeners:', error);
        }

        // Update state periodically
        const interval = setInterval(async () => {
            try {
                const state = await GetCurrentState();
                const time = await GetRemainingTime();
                setCurrentState(state as SessionStateType);
                setRemainingTime(time);
            } catch (error) {
                console.error('Error updating state:', error);
            }
        }, APP_CONSTANTS.TIMER_UPDATE_INTERVAL);

        // Cleanup function
        return () => {
            clearInterval(interval);
            
            // Call cleanup functions if available
            cleanupFunctions.forEach(cleanup => {
                try {
                    cleanup();
                } catch (error) {
                    console.error('Error cleaning up event listener:', error);
                }
            });
            
            // Fallback: try EventsOff if available
            try {
                if (typeof EventsOff === 'function') {
                    EventsOff(EVENTS.SESSION_START);
                    EventsOff(EVENTS.SESSION_END);
                    EventsOff(EVENTS.SESSION_PAUSE);
                    EventsOff(EVENTS.SESSION_RESUME);
                    EventsOff(EVENTS.TIMER_TICK);
                    EventsOff(EVENTS.WARNING);
                }
            } catch (error) {
                console.error('EventsOff not available or failed:', error);
            }
        };
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
