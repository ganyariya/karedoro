export namespace domain {
	
	export class Session {
	    State: number;
	    // Go type: time
	    StartTime: any;
	    Duration: number;
	    IsPaused: boolean;
	    // Go type: time
	    PausedAt: any;
	    PauseDuration: number;
	
	    static createFrom(source: any = {}) {
	        return new Session(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.State = source["State"];
	        this.StartTime = this.convertValues(source["StartTime"], null);
	        this.Duration = source["Duration"];
	        this.IsPaused = source["IsPaused"];
	        this.PausedAt = this.convertValues(source["PausedAt"], null);
	        this.PauseDuration = source["PauseDuration"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

