<template>
    <div class="p-4">
        <!-- 查询条件 -->
        <el-card class="mb-4">
            <el-form :inline="true" :model="query">
                <el-form-item label="消费类型">
                    <el-select v-model="query.consume_type" clearable style="width: 200px">
                        <el-option label="模型消费" value="model" />
                        <el-option label="充值" value="recharge" />
                    </el-select>
                </el-form-item>

                <el-form-item>
                    <el-button type="primary" @click="fetchList">查询</el-button>
                </el-form-item>
            </el-form>
        </el-card>

        <!-- 列表 -->
        <el-card>
            <el-table :data="list" border stripe>
                <el-table-column prop="id" label="ID" width="80" />

                <el-table-column prop="consume_type" label="消费类型" width="120">
                    <template #default="{ row }">
                        <el-tag type="success">{{ row.consume_type }}</el-tag>
                    </template>
                </el-table-column>

                <!-- 服务商 -->
                <el-table-column label="服务商" width="140">
                    <template #default="{ row }">
                        <el-tag effect="plain">
                            {{ row.actual_provider || "-" }}
                        </el-tag>
                    </template>
                </el-table-column>

                <!-- 服务协议 -->
                <el-table-column label="服务协议" min-width="180">
                    <template #default="{ row }">
                        <el-tooltip v-if="row.actual_provider_id" :content="row.actual_provider_id" placement="top">
                            <el-tag type="info" class="truncate max-w-[160px]">
                                {{ row.actual_provider_id }}
                            </el-tag>
                        </el-tooltip>
                        <span v-else>-</span>
                    </template>
                </el-table-column>

                <!-- 金额 -->
                <el-table-column label="消费金额">
                    <template #default="{ row }">
                        {{ (row.total_consumed / 100).toFixed(2) }}
                    </template>
                </el-table-column>

                <!-- 时间 -->
                <el-table-column label="时间" width="180">
                    <template #default="{ row }">
                        {{ formatTime(row.created) }}
                    </template>
                </el-table-column>

                <el-table-column label="操作" width="120">
                    <template #default="{ row }">
                        <el-button link type="primary" @click="openDetail(row)">
                            查看明细
                        </el-button>
                    </template>
                </el-table-column>
            </el-table>

            <el-pagination class="mt-4" background layout="total, prev, pager, next" :total="total"
                :page-size="query.page.limit" @current-change="handlePageChange" />
        </el-card>

        <!-- 明细弹窗（保持不动） -->
        <el-dialog v-model="detailVisible" title="消费明细" width="900px">
            <!-- Text / Model -->
            <el-table v-if="detailType === 'text'" :data="detailList" border>
                <el-table-column prop="input_tokens" label="输入 Tokens" />
                <el-table-column prop="output_tokens" label="输出 Tokens" />
                <el-table-column prop="cache_tokens" label="缓存 Tokens" />

                <el-table-column label="输入费用">
                    <template #default="{ row }">
                        {{ (row.input_tokens * row.input_price / 100).toFixed(4) }}
                    </template>
                </el-table-column>

                <el-table-column label="输出费用">
                    <template #default="{ row }">
                        {{ (row.output_tokens * row.output_price / 100).toFixed(4) }}
                    </template>
                </el-table-column>

                <el-table-column label="时间">
                    <template #default="{ row }">
                        {{ formatTime(row.created) }}
                    </template>
                </el-table-column>
            </el-table>

            <!-- Image -->
            <el-table v-else-if="detailType === 'image'" :data="detailList" border>
                <el-table-column prop="quality" label="质量" />
                <el-table-column prop="size" label="尺寸" />
                <el-table-column label="时间">
                    <template #default="{ row }">
                        {{ formatTime(row.created) }}
                    </template>
                </el-table-column>
            </el-table>

            <!-- Video -->
            <el-table v-else-if="detailType === 'video'" :data="detailList" border>
                <el-table-column label="时长（秒）">
                    <template #default="{ row }">
                        {{ row.Seconds.toFixed(2) }}
                    </template>
                </el-table-column>

                <el-table-column prop="size" label="尺寸" />

                <el-table-column label="时间">
                    <template #default="{ row }">
                        {{ formatTime(row.created) }}
                    </template>
                </el-table-column>
            </el-table>

            <el-pagination class="mt-4" background layout="total, prev, pager, next" :total="detailTotal"
                :page-size="detailQuery.page.limit" @current-change="handleDetailPageChange" />
        </el-dialog>

    </div>
</template>


<script setup lang="ts">
import { getUserConsumeDetailApi, getUserConsumeRecordApi } from "@/api/managerApi";
import { ref, reactive, onMounted } from "vue";

// 弹窗类型
const detailType = ref<"text" | "image" | "video">("text");

const query = reactive({
    consume_type: "",
    page: {
        page: 1,
        limit: 20
    }
});

const list = ref<any[]>([]);
const total = ref(0);

/*******************************************/
// 测试数
const mockUserConsumeRecords = Array.from({ length: 50 }).map((_, i) => {
    const id = i + 1;
    return {
        id,
        user_id: 10001,
        node_id: (i % 3) + 1,

        discount_amount: (i % 5) * 100,
        total_consumed: 5000 + i * 120,

        caller: i % 2 === 0 ? "api" : "web",
        model: i % 3 === 0 ? "gpt-4o" : i % 3 === 1 ? "gpt-4.1" : "claude-3.5",
        model_id: (i % 3) + 10,

        actual_provider: i % 2 === 0 ? "openai" : "anthropic",
        actual_provider_id: i % 2 === 0 ? "openai-us" : "anthropic-us",

        consume_type: "model",

        total_cost: 3000 + i * 80,

        created: 1700000000 + i * 60,
        updated: 1700000000 + i * 60
    };
});
// 测试函数
function mockGetUserConsumeRecord(params: {
    page: { page: number; limit: number };
    consume_type?: string;
}) {
    const { page, limit } = params.page;
    const start = (page - 1) * limit;
    const end = start + limit;

    return Promise.resolve({
        data: {
            records: mockUserConsumeRecords.slice(start, end),
            total: mockUserConsumeRecords.length
        }
    });
}

const mockTextDetails = Array.from({ length: 12 }).map((_, i) => ({
    id: i + 1,
    consume_id: 1,
    input_tokens: 1000 + i * 50,
    output_tokens: 800 + i * 40,
    cache_tokens: 200,
    input_price: 2,
    output_price: 4,
    cache_price: 1,
    created: 1700000000 + i * 60
}));

const mockImageDetails = Array.from({ length: 8 }).map((_, i) => ({
    id: i + 1,
    consume_id: 1,
    quality: "high",
    size: "1024x1024",
    created: 1700000000 + i * 120
}));

const mockVideoDetails = Array.from({ length: 6 }).map((_, i) => ({
    id: i + 1,
    consume_id: 1,
    Seconds: 30.5 + i * 5,
    size: "1080p",
    created: 1700000000 + i * 180
}));
function mockGetUserConsumeDetail(_: any) {
    return {
        data: {
            consume_record: mockUserConsumeRecords[0],
            details: mockVideoDetails, // 👈 可切 image / video
            total: mockTextDetails.length
        }
    };
}

/****************************************** */

const fetchList = async () => {
    const res = await getUserConsumeRecordApi(query);
    // const res = await mockGetUserConsumeRecord(query);
    console.log("res record =============", res)
    list.value = res.data.records;
    total.value = res.data.total;
};

const handlePageChange = (page: number) => {
    query.page.page = page;
    fetchList();
};

/* ---------- 明细 ---------- */

const detailVisible = ref(false);
const detailList = ref<any[]>([]);
const detailTotal = ref(0);

const detailQuery = reactive({
    consume_id: 0,
    consume_type: "",
    page: {
        page: 1,
        limit: 10
    }
});

const openDetail = async (row: any) => {
    detailQuery.consume_id = row.id;
    detailQuery.consume_type = row.consume_type;
    detailQuery.page.page = 1;
    detailVisible.value = true;
    await fetchDetail();
};

const fetchDetail = async () => {
    const res = await getUserConsumeDetailApi(detailQuery);

    // const res = mockGetUserConsumeDetail(detailQuery); // 测试用


    detailList.value = res.data.details;
    detailTotal.value = res.data.total;
    // 自动识别明细类型
    const first = res.data.details?.[0];
    if (!first) return;

    if ("input_tokens" in first) detailType.value = "text";
    else if ("quality" in first) detailType.value = "image";
    else if ("Seconds" in first) detailType.value = "video";
};

const handleDetailPageChange = (page: number) => {
    detailQuery.page.page = page;
    fetchDetail();
};

const formatTime = (ts: number) => {
    if (!ts) return "-";
    return new Date(ts * 1000).toLocaleString();
};

onMounted(() => {
    fetchList();
});
</script>
