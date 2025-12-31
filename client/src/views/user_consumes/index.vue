<template>
    <div class="p-4">
        <!-- 查询条件 -->
        <el-card class="mb-4">
            <el-form :inline="true" :model="query">
                <el-form-item label="消费类型">
                    <el-select v-model="query.consume_type" clearable style="width: 200px">
                        <el-option label="文字" value="text" />
                        <el-option label="图片" value="image" />
                        <el-option label="视频" value="video" />
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
import { ref, reactive, onMounted, watch } from "vue";

// 弹窗类型
const detailType = ref<"text" | "image" | "video">("text");

const query = reactive({
    consume_type: "text",
    page: {
        page: 1,
        limit: 20
    }
});

const list = ref<any[]>([]);
const total = ref(0);


const fetchList = async () => {
    const res = await getUserConsumeRecordApi(query);
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
    consume_type: "text",
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
