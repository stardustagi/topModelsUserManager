<template>
  <el-form
    ref="ruleFormRef"
    :model="ruleForm"
    :rules="updateRules"
    size="large"
  >
    <Motion :delay="100">
      <el-form-item prop="email">
        <el-input
          v-model="ruleForm.mail"
          clearable
          placeholder="邮箱"
          :prefix-icon="useRenderIcon(Mail)"
        />
      </el-form-item>
    </Motion>

    <Motion :delay="200">
      <el-form-item prop="password">
        <el-input
          v-model="ruleForm.password"
          clearable
          show-password
          :placeholder="t('login.purePassword')"
          :prefix-icon="useRenderIcon(Lock)"
        />
      </el-form-item>
    </Motion>

    <!-- <Motion :delay="150">
      <el-form-item prop="verifyCode">
        <div class="w-full flex justify-between">
          <el-input
            v-model="ruleForm.verifyCode"
            clearable
            placeholder="邮箱验证码"
            :prefix-icon="useRenderIcon(Keyhole)"
          />
          <el-button
            :disabled="isDisabled"
            class="ml-2!"
            @click="useVerifyCode().start(ruleFormRef, 'email')"
          >
            {{
              text.length > 0
                ? text + t("login.pureInfo")
                : t("login.pureGetVerifyCode")
            }}
          </el-button>
        </div>
      </el-form-item>
    </Motion> -->

    <!--验证码-->
    <!-- <Motion :delay="200">
      <el-form-item prop="imageCode">
        <el-input
          v-model="ruleForm.imageCode"
          clearable
          :placeholder="t('login.pureVerifyCode')"
          :prefix-icon="useRenderIcon(Keyhole)"
        >
          <template v-slot:append>
            <ReImageVerify v-model:code="imgCode" />
          </template>
        </el-input>
      </el-form-item>
    </Motion> -->

    <!-- <Motion :delay="300">
      <el-form-item>
        <el-checkbox v-model="checked">
          {{ t("login.pureReadAccept") }}
        </el-checkbox>
        <el-button link type="primary">
          {{ t("login.purePrivacyPolicy") }}
        </el-button>
      </el-form-item>
    </Motion> -->

    <Motion :delay="350">
      <el-form-item>
        <el-button
          class="w-full"
          size="default"
          type="primary"
          :loading="loading"
          @click="onUpdate(ruleFormRef)"
        >
          {{ t("login.pureDefinite") }}
        </el-button>
      </el-form-item>
    </Motion>

    <Motion :delay="400">
      <el-form-item>
        <el-button class="w-full" size="default" @click="onBack">
          {{ t("login.pureBack") }}
        </el-button>
      </el-form-item>
    </Motion>
  </el-form>
</template>

<script setup lang="ts">
import Lock from "~icons/ri/lock-fill";
import { useI18n } from "vue-i18n";
import { ref, reactive, onMounted } from "vue";
import Motion from "../utils/motion";
import { message } from "@/utils/message";
import { updateRules } from "../utils/rule";
import type { FormInstance } from "element-plus";
import { useVerifyCode } from "../utils/verifyCode";
import { $t, transformI18n } from "@/plugins/i18n";
import { useUserStoreHook } from "@/store/modules/user";
import { useRenderIcon } from "@/components/ReIcon/src/hooks";
import Iphone from "~icons/ep/iphone";
import Keyhole from "~icons/ri/shield-keyhole-line";
import { ReImageVerify } from "@/components/ReImageVerify";
import { Mail } from "@icon-park/vue-next";
import { nodeUserEmailRegisterApi } from "@/api/managerApi";

defineOptions({
  name: "MailRegister"
});

const { t } = useI18n();
const checked = ref(false);
const loading = ref(false);
const ruleForm = reactive({
  mail: "",
  password: ""
});
const ruleFormRef = ref<FormInstance>();
const { isDisabled, text } = useVerifyCode();

const onUpdate = async (formEl: FormInstance | undefined) => {
  loading.value = true;
  console.log("qqqqqqqqqqqqqqqqqqq");
  if (!formEl) return;
  await formEl.validate(async valid => {
    if (valid) {
      const data = {
        mail: ruleForm.mail,
        password: ruleForm.password
      };
      const res = await nodeUserEmailRegisterApi(data);
      console.log("res === ", res);
      if (res.errcode === 0) {
        message("注册成功，请查收邮件");
      }
      // if (checked.value) {
      //   // 模拟请求，需根据实际开发进行修改
      //   setTimeout(() => {
      //     message(transformI18n($t("login.pureRegisterSuccess")), {
      //       type: "success"
      //     });
      //     loading.value = false;
      //   }, 2000);
      // } else {
      //   loading.value = false;
      //   message(transformI18n($t("login.pureTickPrivacy")), {
      //     type: "warning"
      //   });
      // }
      loading.value = false;
    } else {
      loading.value = false;
    }
  });
};

function onBack() {
  useVerifyCode().end();
  useUserStoreHook().SET_CURRENTPAGE(0);
}
</script>
