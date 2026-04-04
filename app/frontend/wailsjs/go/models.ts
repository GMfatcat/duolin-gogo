export namespace dashboard {
	
	export class WeakTopic {
	    tag: string;
	    wrongCount: number;
	    seenCount: number;
	    accuracy: number;
	
	    static createFrom(source: any = {}) {
	        return new WeakTopic(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.tag = source["tag"];
	        this.wrongCount = source["wrongCount"];
	        this.seenCount = source["seenCount"];
	        this.accuracy = source["accuracy"];
	    }
	}
	export class Summary {
	    studiedToday: number;
	    correctRate: number;
	    nextReviewAt: string;
	    weakTopics: WeakTopic[];
	
	    static createFrom(source: any = {}) {
	        return new Summary(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.studiedToday = source["studiedToday"];
	        this.correctRate = source["correctRate"];
	        this.nextReviewAt = source["nextReviewAt"];
	        this.weakTopics = this.convertValues(source["weakTopics"], WeakTopic);
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

export namespace diagnostics {
	
	export class Error {
	    source_path: string;
	    severity?: string;
	    code: string;
	    field?: string;
	    message: string;
	
	    static createFrom(source: any = {}) {
	        return new Error(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.source_path = source["source_path"];
	        this.severity = source["severity"];
	        this.code = source["code"];
	        this.field = source["field"];
	        this.message = source["message"];
	    }
	}

}

export namespace main {
	
	export class ActionStatus {
	    message: string;
	
	    static createFrom(source: any = {}) {
	        return new ActionStatus(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.message = source["message"];
	    }
	}
	export class AnswerChoice {
	    value: string;
	    labelZh: string;
	    labelEn: string;
	
	    static createFrom(source: any = {}) {
	        return new AnswerChoice(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.value = source["value"];
	        this.labelZh = source["labelZh"];
	        this.labelEn = source["labelEn"];
	    }
	}
	export class AppInfo {
	    name: string;
	    focusTopic: string;
	    defaultLanguage: string;
	
	    static createFrom(source: any = {}) {
	        return new AppInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.focusTopic = source["focusTopic"];
	        this.defaultLanguage = source["defaultLanguage"];
	    }
	}
	export class StudyCard {
	    id: string;
	    title: string;
	    titleZh: string;
	    titleEn: string;
	    questionType: string;
	    questionText: string;
	    questionTextZh: string;
	    questionTextEn: string;
	    choices: AnswerChoice[];
	    clickbait: string;
	    clickbaitZh: string;
	    clickbaitEn: string;
	    reviewHint: string;
	    reviewHintZh: string;
	    reviewHintEn: string;
	    explanationZh: string;
	    explanationEn: string;
	    shownAt: string;
	
	    static createFrom(source: any = {}) {
	        return new StudyCard(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.title = source["title"];
	        this.titleZh = source["titleZh"];
	        this.titleEn = source["titleEn"];
	        this.questionType = source["questionType"];
	        this.questionText = source["questionText"];
	        this.questionTextZh = source["questionTextZh"];
	        this.questionTextEn = source["questionTextEn"];
	        this.choices = this.convertValues(source["choices"], AnswerChoice);
	        this.clickbait = source["clickbait"];
	        this.clickbaitZh = source["clickbaitZh"];
	        this.clickbaitEn = source["clickbaitEn"];
	        this.reviewHint = source["reviewHint"];
	        this.reviewHintZh = source["reviewHintZh"];
	        this.reviewHintEn = source["reviewHintEn"];
	        this.explanationZh = source["explanationZh"];
	        this.explanationEn = source["explanationEn"];
	        this.shownAt = source["shownAt"];
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
	export class AuthoringPreviewFile {
	    path: string;
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new AuthoringPreviewFile(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.path = source["path"];
	        this.name = source["name"];
	    }
	}
	export class AuthoringPreviewData {
	    files: AuthoringPreviewFile[];
	    selectedPath: string;
	    currentCard?: StudyCard;
	    importErrors: diagnostics.Error[];
	
	    static createFrom(source: any = {}) {
	        return new AuthoringPreviewData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.files = this.convertValues(source["files"], AuthoringPreviewFile);
	        this.selectedPath = source["selectedPath"];
	        this.currentCard = this.convertValues(source["currentCard"], StudyCard);
	        this.importErrors = this.convertValues(source["importErrors"], diagnostics.Error);
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
	
	export class ScheduleSettings {
	    notificationIntervalMinutes: number;
	    reviewTime: string;
	    activeHoursEnabled: boolean;
	    activeHoursStart: string;
	    activeHoursEnd: string;
	
	    static createFrom(source: any = {}) {
	        return new ScheduleSettings(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.notificationIntervalMinutes = source["notificationIntervalMinutes"];
	        this.reviewTime = source["reviewTime"];
	        this.activeHoursEnabled = source["activeHoursEnabled"];
	        this.activeHoursStart = source["activeHoursStart"];
	        this.activeHoursEnd = source["activeHoursEnd"];
	    }
	}
	export class NotificationSettings {
	    style: string;
	    titleMode: string;
	
	    static createFrom(source: any = {}) {
	        return new NotificationSettings(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.style = source["style"];
	        this.titleMode = source["titleMode"];
	    }
	}
	export class DashboardStats {
	    studiedToday: number;
	    correctRate: number;
	
	    static createFrom(source: any = {}) {
	        return new DashboardStats(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.studiedToday = source["studiedToday"];
	        this.correctRate = source["correctRate"];
	    }
	}
	export class DashboardData {
	    info: AppInfo;
	    preferredLanguage: string;
	    stats: DashboardStats;
	    summary: dashboard.Summary;
	    importErrors: diagnostics.Error[];
	    notificationSettings: NotificationSettings;
	    scheduleSettings: ScheduleSettings;
	    currentCard?: StudyCard;
	    reviewQueue: StudyCard[];
	    reviewMode: boolean;
	
	    static createFrom(source: any = {}) {
	        return new DashboardData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.info = this.convertValues(source["info"], AppInfo);
	        this.preferredLanguage = source["preferredLanguage"];
	        this.stats = this.convertValues(source["stats"], DashboardStats);
	        this.summary = this.convertValues(source["summary"], dashboard.Summary);
	        this.importErrors = this.convertValues(source["importErrors"], diagnostics.Error);
	        this.notificationSettings = this.convertValues(source["notificationSettings"], NotificationSettings);
	        this.scheduleSettings = this.convertValues(source["scheduleSettings"], ScheduleSettings);
	        this.currentCard = this.convertValues(source["currentCard"], StudyCard);
	        this.reviewQueue = this.convertValues(source["reviewQueue"], StudyCard);
	        this.reviewMode = source["reviewMode"];
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
	
	export class DraftReviewData {
	    currentCard?: StudyCard;
	    importErrors: diagnostics.Error[];
	
	    static createFrom(source: any = {}) {
	        return new DraftReviewData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.currentCard = this.convertValues(source["currentCard"], StudyCard);
	        this.importErrors = this.convertValues(source["importErrors"], diagnostics.Error);
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
	
	export class SaveDraftStatus {
	    message: string;
	    savedPath: string;
	    topic: string;
	    successful: boolean;
	
	    static createFrom(source: any = {}) {
	        return new SaveDraftStatus(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.message = source["message"];
	        this.savedPath = source["savedPath"];
	        this.topic = source["topic"];
	        this.successful = source["successful"];
	    }
	}
	
	
	export class SubmitAnswerResult {
	    cardId: string;
	    isCorrect: boolean;
	    correctAnswer: string;
	    feedback: string;
	    reviewHint: string;
	    preferredLanguage: string;
	    stats: DashboardStats;
	
	    static createFrom(source: any = {}) {
	        return new SubmitAnswerResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.cardId = source["cardId"];
	        this.isCorrect = source["isCorrect"];
	        this.correctAnswer = source["correctAnswer"];
	        this.feedback = source["feedback"];
	        this.reviewHint = source["reviewHint"];
	        this.preferredLanguage = source["preferredLanguage"];
	        this.stats = this.convertValues(source["stats"], DashboardStats);
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
	export class ValidationStatus {
	    message: string;
	    importErrors: diagnostics.Error[];
	
	    static createFrom(source: any = {}) {
	        return new ValidationStatus(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.message = source["message"];
	        this.importErrors = this.convertValues(source["importErrors"], diagnostics.Error);
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

