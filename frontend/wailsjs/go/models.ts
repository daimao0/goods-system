export namespace api {
	
	export class Page[goods-system/internal/application/vo.GoodsVO] {
	    page: number;
	    size: number;
	    total: number;
	    data: vo.GoodsVO[];
	
	    static createFrom(source: any = {}) {
	        return new Page[goods-system/internal/application/vo.GoodsVO](source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.page = source["page"];
	        this.size = source["size"];
	        this.total = source["total"];
	        this.data = this.convertValues(source["data"], vo.GoodsVO);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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
	export class RespData[[]goods-system/internal/application/vo.GoodsTypeVO] {
	    code: number;
	    message: string;
	    data: vo.GoodsTypeVO[];
	
	    static createFrom(source: any = {}) {
	        return new RespData[[]goods-system/internal/application/vo.GoodsTypeVO](source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.message = source["message"];
	        this.data = this.convertValues(source["data"], vo.GoodsTypeVO);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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
	export class RespData[[]goods-system/internal/application/vo.GoodsVO] {
	    code: number;
	    message: string;
	    data: vo.GoodsVO[];
	
	    static createFrom(source: any = {}) {
	        return new RespData[[]goods-system/internal/application/vo.GoodsVO](source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.message = source["message"];
	        this.data = this.convertValues(source["data"], vo.GoodsVO);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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
	export class RespData[[]goods-system/internal/application/vo.MemberLevelVO] {
	    code: number;
	    message: string;
	    data: vo.MemberLevelVO[];
	
	    static createFrom(source: any = {}) {
	        return new RespData[[]goods-system/internal/application/vo.MemberLevelVO](source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.message = source["message"];
	        this.data = this.convertValues(source["data"], vo.MemberLevelVO);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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
	export class RespData[[]goods-system/internal/application/vo.MemberVO] {
	    code: number;
	    message: string;
	    data: vo.MemberVO[];
	
	    static createFrom(source: any = {}) {
	        return new RespData[[]goods-system/internal/application/vo.MemberVO](source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.message = source["message"];
	        this.data = this.convertValues(source["data"], vo.MemberVO);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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
	export class RespData[[]goods-system/internal/application/vo.OrderVO] {
	    code: number;
	    message: string;
	    data: vo.OrderVO[];
	
	    static createFrom(source: any = {}) {
	        return new RespData[[]goods-system/internal/application/vo.OrderVO](source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.message = source["message"];
	        this.data = this.convertValues(source["data"], vo.OrderVO);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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
	export class RespData[go/types.Nil] {
	    code: number;
	    message: string;
	    // Go type: types
	    data: any;
	
	    static createFrom(source: any = {}) {
	        return new RespData[go/types.Nil](source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.message = source["message"];
	        this.data = this.convertValues(source["data"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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
	export class RespData[goods-system/internal/application/vo.GoodsVO] {
	    code: number;
	    message: string;
	    data: vo.GoodsVO;
	
	    static createFrom(source: any = {}) {
	        return new RespData[goods-system/internal/application/vo.GoodsVO](source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.message = source["message"];
	        this.data = this.convertValues(source["data"], vo.GoodsVO);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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
	export class RespData[goods-system/internal/application/vo.MemberVO] {
	    code: number;
	    message: string;
	    data: vo.MemberVO;
	
	    static createFrom(source: any = {}) {
	        return new RespData[goods-system/internal/application/vo.MemberVO](source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.message = source["message"];
	        this.data = this.convertValues(source["data"], vo.MemberVO);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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
	export class RespData[goods-system/internal/application/vo.OrderStatisticsVO] {
	    code: number;
	    message: string;
	    data: vo.OrderStatisticsVO;
	
	    static createFrom(source: any = {}) {
	        return new RespData[goods-system/internal/application/vo.OrderStatisticsVO](source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.message = source["message"];
	        this.data = this.convertValues(source["data"], vo.OrderStatisticsVO);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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
	export class RespData[goods-system/internal/application/vo.SystemDesktopPathVO] {
	    code: number;
	    message: string;
	    data: vo.SystemDesktopPathVO;
	
	    static createFrom(source: any = {}) {
	        return new RespData[goods-system/internal/application/vo.SystemDesktopPathVO](source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.message = source["message"];
	        this.data = this.convertValues(source["data"], vo.SystemDesktopPathVO);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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
	export class RespData[goods-system/internal/infrastructure/common/api.Page[goods-system/internal/application/vo.GoodsVO]] {
	    code: number;
	    message: string;
	    // Go type: Page[goods-system/internal/application/vo
	    data: any;
	
	    static createFrom(source: any = {}) {
	        return new RespData[goods-system/internal/infrastructure/common/api.Page[goods-system/internal/application/vo.GoodsVO]](source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.message = source["message"];
	        this.data = this.convertValues(source["data"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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
	export class Page[goods-system/internal/application/vo.MemberVO] {
	    page: number;
	    size: number;
	    total: number;
	    data: vo.MemberVO[];
	
	    static createFrom(source: any = {}) {
	        return new Page[goods-system/internal/application/vo.MemberVO](source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.page = source["page"];
	        this.size = source["size"];
	        this.total = source["total"];
	        this.data = this.convertValues(source["data"], vo.MemberVO);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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
	export class RespData[goods-system/internal/infrastructure/common/api.Page[goods-system/internal/application/vo.MemberVO]] {
	    code: number;
	    message: string;
	    // Go type: Page[goods-system/internal/application/vo
	    data: any;
	
	    static createFrom(source: any = {}) {
	        return new RespData[goods-system/internal/infrastructure/common/api.Page[goods-system/internal/application/vo.MemberVO]](source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.message = source["message"];
	        this.data = this.convertValues(source["data"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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
	export class Page[goods-system/internal/application/vo.OrderVO] {
	    page: number;
	    size: number;
	    total: number;
	    data: vo.OrderVO[];
	
	    static createFrom(source: any = {}) {
	        return new Page[goods-system/internal/application/vo.OrderVO](source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.page = source["page"];
	        this.size = source["size"];
	        this.total = source["total"];
	        this.data = this.convertValues(source["data"], vo.OrderVO);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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
	export class RespData[goods-system/internal/infrastructure/common/api.Page[goods-system/internal/application/vo.OrderVO]] {
	    code: number;
	    message: string;
	    // Go type: Page[goods-system/internal/application/vo
	    data: any;
	
	    static createFrom(source: any = {}) {
	        return new RespData[goods-system/internal/infrastructure/common/api.Page[goods-system/internal/application/vo.OrderVO]](source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.message = source["message"];
	        this.data = this.convertValues(source["data"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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
	export class RespData[interface {}] {
	    code: number;
	    message: string;
	    data: any;
	
	    static createFrom(source: any = {}) {
	        return new RespData[interface {}](source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.message = source["message"];
	        this.data = source["data"];
	    }
	}

}

export namespace request {
	
	export class GoodsEditRequest {
	    id: string;
	    name: string;
	    goodsType: string;
	    goodsNumber: string;
	    price: string;
	    count: string;
	    desc: string;
	
	    static createFrom(source: any = {}) {
	        return new GoodsEditRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.goodsType = source["goodsType"];
	        this.goodsNumber = source["goodsNumber"];
	        this.price = source["price"];
	        this.count = source["count"];
	        this.desc = source["desc"];
	    }
	}
	export class GoodsPageRequest {
	    name: string;
	    goodsType: string;
	    page: number;
	    size: number;
	
	    static createFrom(source: any = {}) {
	        return new GoodsPageRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.goodsType = source["goodsType"];
	        this.page = source["page"];
	        this.size = source["size"];
	    }
	}
	export class GoodsSaveRequest {
	    name: string;
	    goodsType: string;
	    goodsNumber: string;
	    price: string;
	    count: string;
	    desc: string;
	
	    static createFrom(source: any = {}) {
	        return new GoodsSaveRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.goodsType = source["goodsType"];
	        this.goodsNumber = source["goodsNumber"];
	        this.price = source["price"];
	        this.count = source["count"];
	        this.desc = source["desc"];
	    }
	}
	export class GoodsTypeDeleteRequest {
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new GoodsTypeDeleteRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	    }
	}
	export class GoodsTypeEditRequest {
	    id: string;
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new GoodsTypeEditRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	    }
	}
	export class GoodsTypeSaveRequest {
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new GoodsTypeSaveRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	    }
	}
	export class MemberLevelSaveRequest {
	    name: string;
	    discount: string;
	
	    static createFrom(source: any = {}) {
	        return new MemberLevelSaveRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.discount = source["discount"];
	    }
	}
	export class MemberLevelUpdateRequest {
	    id: string;
	    name: string;
	    discount: string;
	
	    static createFrom(source: any = {}) {
	        return new MemberLevelUpdateRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.discount = source["discount"];
	    }
	}
	export class MemberPageRequest {
	    name: string;
	    phone: string;
	    levelId: string;
	    account: string;
	    page: number;
	    size: number;
	
	    static createFrom(source: any = {}) {
	        return new MemberPageRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.phone = source["phone"];
	        this.levelId = source["levelId"];
	        this.account = source["account"];
	        this.page = source["page"];
	        this.size = source["size"];
	    }
	}
	export class MemberSaveRequest {
	    name: string;
	    phone: string;
	    levelId: string;
	    account: string;
	
	    static createFrom(source: any = {}) {
	        return new MemberSaveRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.phone = source["phone"];
	        this.levelId = source["levelId"];
	        this.account = source["account"];
	    }
	}
	export class MemberUpdateRequest {
	    id: string;
	    name: string;
	    phone: string;
	    levelId: string;
	    account: string;
	    discount: string;
	
	    static createFrom(source: any = {}) {
	        return new MemberUpdateRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.phone = source["phone"];
	        this.levelId = source["levelId"];
	        this.account = source["account"];
	        this.discount = source["discount"];
	    }
	}
	export class OrderItemRequest {
	    goodsId: string;
	    num: number;
	    realPrice: string;
	
	    static createFrom(source: any = {}) {
	        return new OrderItemRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.goodsId = source["goodsId"];
	        this.num = source["num"];
	        this.realPrice = source["realPrice"];
	    }
	}
	export class OrderPageRequest {
	    memberId: string;
	    startTime: string;
	    endTime: string;
	    page: number;
	    size: number;
	
	    static createFrom(source: any = {}) {
	        return new OrderPageRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.memberId = source["memberId"];
	        this.startTime = source["startTime"];
	        this.endTime = source["endTime"];
	        this.page = source["page"];
	        this.size = source["size"];
	    }
	}
	export class OrderRecordRequest {
	    startTime: string;
	    endTime: string;
	
	    static createFrom(source: any = {}) {
	        return new OrderRecordRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.startTime = source["startTime"];
	        this.endTime = source["endTime"];
	    }
	}
	export class OrderRequest {
	    pay: string;
	    memberId: string;
	    orderItem: OrderItemRequest[];
	
	    static createFrom(source: any = {}) {
	        return new OrderRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.pay = source["pay"];
	        this.memberId = source["memberId"];
	        this.orderItem = this.convertValues(source["orderItem"], OrderItemRequest);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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

export namespace vo {
	
	export class GoodsTypeVO {
	    id: string;
	    name: string;
	    createTime: string;
	    updateTime: string;
	
	    static createFrom(source: any = {}) {
	        return new GoodsTypeVO(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.createTime = source["createTime"];
	        this.updateTime = source["updateTime"];
	    }
	}
	export class GoodsVO {
	    id: string;
	    name: string;
	    goodsNumber: string;
	    goodsType: string;
	    price: string;
	    count: string;
	    desc: string;
	    createTime: string;
	    updateTime: string;
	
	    static createFrom(source: any = {}) {
	        return new GoodsVO(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.goodsNumber = source["goodsNumber"];
	        this.goodsType = source["goodsType"];
	        this.price = source["price"];
	        this.count = source["count"];
	        this.desc = source["desc"];
	        this.createTime = source["createTime"];
	        this.updateTime = source["updateTime"];
	    }
	}
	export class MemberLevelVO {
	    id: string;
	    name: string;
	    discount: string;
	    createTime: string;
	    updateTime: string;
	
	    static createFrom(source: any = {}) {
	        return new MemberLevelVO(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.discount = source["discount"];
	        this.createTime = source["createTime"];
	        this.updateTime = source["updateTime"];
	    }
	}
	export class MemberVO {
	    id: string;
	    name: string;
	    phone: string;
	    level: MemberLevelVO;
	    discount: string;
	    account: string;
	    createTime: string;
	    updateTime: string;
	
	    static createFrom(source: any = {}) {
	        return new MemberVO(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.phone = source["phone"];
	        this.level = this.convertValues(source["level"], MemberLevelVO);
	        this.discount = source["discount"];
	        this.account = source["account"];
	        this.createTime = source["createTime"];
	        this.updateTime = source["updateTime"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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
	export class OrderItemVO {
	    id: string;
	    goods: GoodsVO;
	    realPrice: string;
	    num: number;
	    createTime: string;
	
	    static createFrom(source: any = {}) {
	        return new OrderItemVO(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.goods = this.convertValues(source["goods"], GoodsVO);
	        this.realPrice = source["realPrice"];
	        this.num = source["num"];
	        this.createTime = source["createTime"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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
	export class OrderStatisticsVO {
	    orderCount: string;
	    goodsPrice: string;
	    memberPay: string;
	    preferentialPrice: string;
	
	    static createFrom(source: any = {}) {
	        return new OrderStatisticsVO(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.orderCount = source["orderCount"];
	        this.goodsPrice = source["goodsPrice"];
	        this.memberPay = source["memberPay"];
	        this.preferentialPrice = source["preferentialPrice"];
	    }
	}
	export class OrderVO {
	    id: string;
	    member: MemberVO;
	    orderItems: OrderItemVO[];
	    pay: string;
	    createTime: string;
	
	    static createFrom(source: any = {}) {
	        return new OrderVO(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.member = this.convertValues(source["member"], MemberVO);
	        this.orderItems = this.convertValues(source["orderItems"], OrderItemVO);
	        this.pay = source["pay"];
	        this.createTime = source["createTime"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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
	export class SystemDesktopPathVO {
	    path: string;
	
	    static createFrom(source: any = {}) {
	        return new SystemDesktopPathVO(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.path = source["path"];
	    }
	}

}

