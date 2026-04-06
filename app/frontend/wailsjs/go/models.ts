export namespace dashboard {
	
	export class TopicProgress {
	    topic: string;
	    seenCount: number;
	    correctCount: number;
	    wrongCount: number;
	    accuracy: number;
	
	    static createFrom(source: any = {}) {
	        return new TopicProgress(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.topic = source["topic"];
	        this.seenCount = source["seenCount"];
	        this.correctCount = source["correctCount"];
	        this.wrongCount = source["wrongCount"];
	        this.accuracy = source["accuracy"];
	    }
	}
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
	    studyStreak: number;
	    nextReviewAt: string;
	    weakTopics: WeakTopic[];
	    topicProgress: TopicProgress[];
	    weakestDeck?: TopicProgress;
	
	    static createFrom(source: any = {}) {
	        return new Summary(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.studiedToday = source["studiedToday"];
	        this.correctRate = source["correctRate"];
	        this.studyStreak = source["studyStreak"];
	        this.nextReviewAt = source["nextReviewAt"];
	        this.weakTopics = this.convertValues(source["weakTopics"], WeakTopic);
	        this.topicProgress = this.convertValues(source["topicProgress"], TopicProgress);
	        this.weakestDeck = this.convertValues(source["weakestDeck"], TopicProgress);
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
	    suggestion_zh?: string;
	    suggestion_en?: string;
	
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
	        this.suggestion_zh = source["suggestion_zh"];
	        this.suggestion_en = source["suggestion_en"];
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
	    modifiedAt: string;
	    cardId: string;
	    titleZh: string;
	    titleEn: string;
	    topic: string;
	    searchText: string;
	
	    static createFrom(source: any = {}) {
	        return new AuthoringPreviewFile(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.path = source["path"];
	        this.name = source["name"];
	        this.modifiedAt = source["modifiedAt"];
	        this.cardId = source["cardId"];
	        this.titleZh = source["titleZh"];
	        this.titleEn = source["titleEn"];
	        this.topic = source["topic"];
	        this.searchText = source["searchText"];
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
	
	export class AuthoringPromptData {
	    content: string;
	
	    static createFrom(source: any = {}) {
	        return new AuthoringPromptData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.content = source["content"];
	    }
	}
	export class BatchImportItem {
	    index: number;
	    status: string;
	    cardId: string;
	    savedPath: string;
	    warningCount: number;
	    errorCount: number;
	
	    static createFrom(source: any = {}) {
	        return new BatchImportItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.index = source["index"];
	        this.status = source["status"];
	        this.cardId = source["cardId"];
	        this.savedPath = source["savedPath"];
	        this.warningCount = source["warningCount"];
	        this.errorCount = source["errorCount"];
	    }
	}
	export class BatchImportReport {
	    savedCount: number;
	    skippedCount: number;
	    warningCount: number;
	    errorCount: number;
	    items: BatchImportItem[];
	
	    static createFrom(source: any = {}) {
	        return new BatchImportReport(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.savedCount = source["savedCount"];
	        this.skippedCount = source["skippedCount"];
	        this.warningCount = source["warningCount"];
	        this.errorCount = source["errorCount"];
	        this.items = this.convertValues(source["items"], BatchImportItem);
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
	export class DGInteractionStatus {
	    key: string;
	    title: string;
	    body: string;
	    variant: string;
	    pose: string;
	    stage: number;
	
	    static createFrom(source: any = {}) {
	        return new DGInteractionStatus(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.key = source["key"];
	        this.title = source["title"];
	        this.body = source["body"];
	        this.variant = source["variant"];
	        this.pose = source["pose"];
	        this.stage = source["stage"];
	    }
	}
	export class ScheduleSettings {
	    notificationIntervalMinutes: number;
	    reviewTime: string;
	    activeHoursEnabled: boolean;
	    activeHoursStart: string;
	    activeHoursEnd: string;
	    revealSpeed: string;
	
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
	        this.revealSpeed = source["revealSpeed"];
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
	    selectedTopic: string;
	    availableTopics: string[];
	    petStage: number;
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
	        this.selectedTopic = source["selectedTopic"];
	        this.availableTopics = source["availableTopics"];
	        this.petStage = source["petStage"];
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
	
	export class DraftReviewItem {
	    index: number;
	    currentCard?: StudyCard;
	    importErrors: diagnostics.Error[];
	    valid: boolean;
	
	    static createFrom(source: any = {}) {
	        return new DraftReviewItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.index = source["index"];
	        this.currentCard = this.convertValues(source["currentCard"], StudyCard);
	        this.importErrors = this.convertValues(source["importErrors"], diagnostics.Error);
	        this.valid = source["valid"];
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
	    items: DraftReviewItem[];
	    currentCard?: StudyCard;
	    importErrors: diagnostics.Error[];
	
	    static createFrom(source: any = {}) {
	        return new DraftReviewData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.items = this.convertValues(source["items"], DraftReviewItem);
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
	
	export class DraftScaffoldData {
	    raw: string;
	
	    static createFrom(source: any = {}) {
	        return new DraftScaffoldData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.raw = source["raw"];
	    }
	}
	export class LearnBreakStatus {
	    message: string;
	    unlockAt: string;
	    durationMinutes: number;
	
	    static createFrom(source: any = {}) {
	        return new LearnBreakStatus(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.message = source["message"];
	        this.unlockAt = source["unlockAt"];
	        this.durationMinutes = source["durationMinutes"];
	    }
	}
	
	export class SaveDraftStatus {
	    message: string;
	    savedPath: string;
	    topic: string;
	    successful: boolean;
	    report?: BatchImportReport;
	
	    static createFrom(source: any = {}) {
	        return new SaveDraftStatus(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.message = source["message"];
	        this.savedPath = source["savedPath"];
	        this.topic = source["topic"];
	        this.successful = source["successful"];
	        this.report = this.convertValues(source["report"], BatchImportReport);
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

