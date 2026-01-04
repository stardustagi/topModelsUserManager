<template>
    <div class="p-4">
        <el-card shadow="never">
            <el-table :data="list" border stripe v-loading="loading">
                <el-table-column prop="day" label="日期" width="120" />

                <el-table-column prop="request_count" label="请求数" align="right" />

                <el-table-column prop="success_count" label="成功数" align="right" />

                <el-table-column prop="fail_count" label="失败数" align="right" />

                <el-table-column label="费用 (万)" align="right">
                    <template #default="{ row }">
                        {{ (row.cost / 10000).toFixed(2) }}
                    </template>
                </el-table-column>
            </el-table>

            <el-pagination background layout="total, prev, pager, next" :total="total" :page-size="20"
                :current-page="currentPage" @current-change="onPageChange" />
        </el-card>
    </div>
</template>

<script setup lang="ts">
import { DayReportItem } from "@/api/apiParamsResp";
import { getDailyReportApi } from "@/api/managerApi";
import { ref, onMounted } from "vue";

// 测试数据***************/
function mockFetch(skip: number, limit: number) {
    const total = 50;

    const all = Array.from({ length: total }).map((_, i) => ({
        day: `2026-01-${(i + 1).toString().padStart(2, "0")}`,
        request_count: 1000 + i * 10,
        success_count: 980 + i * 8,
        fail_count: 20,
        cost: 300000 + i * 5000 // 分
    }));

    return {
        items: all.slice(skip, skip + limit),
        total
    };
}


/************************* */

const loading = ref(false);
const list = ref<DayReportItem[]>([]);
const total = ref(0);

const currentPage = ref(1);
const pageSize = ref(20);

const fetchData = async () => {
    loading.value = true;
    try {
        const skip = (currentPage.value - 1) * pageSize.value;
        const res = await getDailyReportApi({
            skip,
            limit: pageSize.value,
            sort: "id desc"
        });

        list.value = res.data.items || [];
        total.value = res.data.total || 0;

        // const res = mockFetch(skip, pageSize.value);
        // list.value = res.items;
        // total.value = res.total;
    } finally {
        loading.value = false;
    }
};

const onPageChange = (page: number) => {
    currentPage.value = page;
    fetchData();
};

const onSizeChange = (size: number) => {
    pageSize.value = size;
    currentPage.value = 1;
    fetchData();
};

onMounted(() => {
    fetchData();
});
</script>
