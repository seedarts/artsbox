export namespace config {
	
	export class Settings {
	    api_key: string;
	    base_url: string;
	    model: string;
	    save_dir: string;
	
	    static createFrom(source: any = {}) {
	        return new Settings(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.api_key = source["api_key"];
	        this.base_url = source["base_url"];
	        this.model = source["model"];
	        this.save_dir = source["save_dir"];
	    }
	}

}

export namespace main {
	
	export class GenerateResult {
	    success: boolean;
	    message: string;
	    images: string[];
	
	    static createFrom(source: any = {}) {
	        return new GenerateResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.success = source["success"];
	        this.message = source["message"];
	        this.images = source["images"];
	    }
	}

}

