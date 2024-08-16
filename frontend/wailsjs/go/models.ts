export namespace fundraising {
	
	export class FundraisingWithHistory {
	    id: number;
	    name: string;
	    description: string;
	    goal: number;
	    url: string;
	    history: fundraising_history.FundraisingHistory[];
	
	    static createFrom(source: any = {}) {
	        return new FundraisingWithHistory(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.description = source["description"];
	        this.goal = source["goal"];
	        this.url = source["url"];
	        this.history = this.convertValues(source["history"], fundraising_history.FundraisingHistory);
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

export namespace fundraising_history {
	
	export class FundraisingHistory {
	    id: number;
	    fundraising_id: number;
	    raised: number;
	    sync_time: string;
	
	    static createFrom(source: any = {}) {
	        return new FundraisingHistory(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.fundraising_id = source["fundraising_id"];
	        this.raised = source["raised"];
	        this.sync_time = source["sync_time"];
	    }
	}

}

